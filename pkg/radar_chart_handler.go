package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	charts "github.com/vicanso/go-charts/v2"
)

type RadarChartHandler struct {
	chart *RadarChart
}

func NewRadarChartHandler() *RadarChartHandler {
	return &RadarChartHandler{
		chart: NewRadarChart(),
	}
}

type RadarChartRequest struct {
	ChartData     string `json:"data" query:"data" form:"data" validate:"required" message:"data is required"`
	ChartTitle    string `json:"title" query:"title" form:"title"`
	ChartSubtitle string `json:"subtitle" query:"subtitle" form:"subtitle"`
	Height        int    `json:"height" query:"height" form:"height"`
	Width         int    `json:"width" query:"width" form:"width"`
}

type RadarChartData struct {
	Labels     []string    `json:"labels"`
	Names      []string    `json:"names"`
	Indicators []float64   `json:"indicators"`
	Values     [][]float64 `json:"values"`
}

func (h *RadarChartHandler) Get(c echo.Context) ([]byte, error) {
	req := new(RadarChartRequest)
	if err := BindRequest(c, req); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	var data RadarChartData
	if err := json.Unmarshal([]byte(req.ChartData), &data); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if len(data.Names) == 0 || len(data.Names) != len(data.Values[0]) {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, "data count is invalid")
	}

	graph, err := charts.RadarRender(
		data.Values,
		charts.TitleOptionFunc(charts.TitleOption{
			Text:    req.ChartTitle,
			Subtext: req.ChartSubtitle,
			Left:    charts.PositionCenter,
		}),
		charts.HeightOptionFunc(req.Height),
		charts.WidthOptionFunc(req.Width),
		charts.LegendLabelsOptionFunc(data.Labels),
		charts.RadarIndicatorOptionFunc(data.Names, h.chart.GetIndicators(data.Values)),
	)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	buf, err := graph.Bytes()
	if err != nil {
		return nil, err
	}
	return buf, nil
}
