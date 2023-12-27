package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	charts "github.com/vicanso/go-charts/v2"
)

type PieChartHandler struct {
	chart *PieChart
}

func NewPieChartHandler() *PieChartHandler {
	return &PieChartHandler{
		chart: NewPieChart(),
	}
}

type PieChartRequest struct {
	ChartData     string `json:"data" query:"data" form:"data" validate:"required" message:"data is required"`
	ChartTitle    string `json:"title" query:"title" form:"title"`
	ChartSubtitle string `json:"subtitle" query:"subtitle" form:"subtitle"`
	Theme         string `json:"theme" query:"theme" form:"theme"`
	Height        int    `json:"height" query:"height" form:"height"`
	Width         int    `json:"width" query:"width" form:"width"`
}

type PieChartData struct {
	Names  []string  `json:"names"`
	Values []float64 `json:"values"`
}

func (h *PieChartHandler) Get(c echo.Context) ([]byte, error) {
	req := new(PieChartRequest)
	if err := BindRequest(c, req); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	var data PieChartData
	if err := json.Unmarshal([]byte(req.ChartData), &data); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if len(data.Values) == 0 || len(data.Values) != len(data.Names) {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, "data provided is invalid")
	}

	p, err := charts.PieRender(
		data.Values,
		charts.TitleOptionFunc(charts.TitleOption{
			Text:    req.ChartTitle,
			Subtext: req.ChartSubtitle,
			Left:    charts.PositionCenter,
		}),
		charts.HeightOptionFunc(req.Height),
		charts.WidthOptionFunc(req.Width),
		charts.PaddingOptionFunc(h.chart.GetPadding()),
		charts.LegendOptionFunc(charts.LegendOption{
			Orient: charts.OrientVertical,
			Data:   data.Names,
			Left:   charts.PositionLeft,
		}),
		charts.PieSeriesShowLabel(),
		func(opt *charts.ChartOption) {
			opt.Theme = req.Theme
		},
	)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	buf, err := p.Bytes()
	return buf, err
}
