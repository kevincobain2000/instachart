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

type PieChartData struct {
	Names  []string  `json:"names"`
	Values []float64 `json:"values"`
}

func (h *PieChartHandler) Get(c echo.Context) ([]byte, error) {
	req := new(ChartRequest)
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
			Text:            req.ChartTitle,
			Subtext:         req.ChartSubtitle,
			SubtextFontSize: 9,
			Left:            charts.PositionCenter,
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
	SetHeaders(c.Response().Header())
	return buf, err
}
