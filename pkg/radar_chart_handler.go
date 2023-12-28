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

type RadarChartData struct {
	Labels     []string    `json:"labels"`
	Names      []string    `json:"names"`
	Indicators []float64   `json:"indicators"`
	Values     [][]float64 `json:"values"`
}

func (h *RadarChartHandler) Get(c echo.Context) ([]byte, error) {
	req := new(ChartRequest)
	if err := BindRequest(c, req); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}

	var data RadarChartData
	if err := json.Unmarshal([]byte(req.ChartData), &data); err != nil {
		msgs := map[string]string{
			"data": "Invalid JSON",
		}
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
	}

	if len(data.Names) == 0 || len(data.Names) != len(data.Values[0]) {
		msgs := map[string]string{
			"data": "Counts are invalid",
		}
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
	}

	graph, err := charts.RadarRender(
		data.Values,
		charts.TitleOptionFunc(charts.TitleOption{
			Text:            req.ChartTitle,
			Subtext:         req.ChartSubtitle,
			SubtextFontSize: 9,
			Left:            charts.PositionCenter,
		}),
		charts.HeightOptionFunc(req.Height),
		charts.WidthOptionFunc(req.Width),
		charts.LegendOptionFunc(charts.LegendOption{
			Orient: charts.OrientVertical,
			Data:   data.Labels,
			Left:   charts.PositionLeft,
		}),
		charts.RadarIndicatorOptionFunc(data.Names, h.chart.GetIndicators(data.Values)),
		func(opt *charts.ChartOption) {
			opt.Theme = req.Theme
		},
	)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	buf, err := graph.Bytes()
	if err != nil {
		return nil, err
	}
	SetHeaders(c.Response().Header())
	return buf, nil
}
