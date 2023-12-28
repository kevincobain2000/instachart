package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	charts "github.com/vicanso/go-charts/v2"
)

type BarChartHandler struct {
	chart *BarChart
}

func NewBarChartHandler() *BarChartHandler {
	return &BarChartHandler{
		chart: NewBarChart(),
	}
}

type BarChartData struct {
	XData []string    `json:"x"`
	YData [][]float64 `json:"y"`
	Names []string    `json:"names"`
}

func (h *BarChartHandler) Get(c echo.Context) ([]byte, error) {
	req := new(ChartRequest)
	if err := BindRequest(c, req); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}

	var data BarChartData
	if err := json.Unmarshal([]byte(req.ChartData), &data); err != nil {
		msgs := map[string]string{
			"data": "Invalid JSON",
		}
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
	}

	if len(data.XData) == 0 || len(data.XData) != len(data.YData[0]) {
		msgs := map[string]string{
			"data": "Counts are invalid",
		}
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
	}

	if !req.Horizontal {
		p, err := charts.BarRender(
			data.YData,
			charts.TitleOptionFunc(charts.TitleOption{
				Text:            req.ChartTitle,
				Subtext:         req.ChartSubtitle,
				SubtextFontSize: 9,
				Left:            charts.PositionCenter,
			}),
			charts.HeightOptionFunc(req.Height),
			charts.WidthOptionFunc(req.Width),
			charts.XAxisDataOptionFunc(data.XData),
			charts.LegendOptionFunc(charts.LegendOption{
				Orient: charts.OrientVertical,
				Data:   data.Names,
				Left:   charts.PositionLeft,
			}),
			charts.MarkLineOptionFunc(0, charts.SeriesMarkDataTypeAverage),
			charts.MarkPointOptionFunc(0, charts.SeriesMarkDataTypeMax,
				charts.SeriesMarkDataTypeMin),
			func(opt *charts.ChartOption) {
				opt.Theme = req.Theme
				opt.ValueFormatter = func(f float64) string {
					return fmt.Sprintf("%.0f%s", f, req.Metric)
				}
				idx := len(opt.SeriesList) - 1
				if len(opt.SeriesList) > 1 {
					idx = 1
				}
				opt.SeriesList[idx].MarkPoint = charts.NewMarkPoint(
					charts.SeriesMarkDataTypeMax,
					charts.SeriesMarkDataTypeMin,
				)
				opt.SeriesList[idx].MarkLine = charts.NewMarkLine(
					charts.SeriesMarkDataTypeAverage,
				)
			},
		)
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}
		buf, err := p.Bytes()
		return buf, err
	}

	p, err := charts.HorizontalBarRender(
		data.YData,
		charts.TitleOptionFunc(charts.TitleOption{
			Text:            req.ChartTitle,
			Subtext:         req.ChartSubtitle,
			SubtextFontSize: 9,
			Left:            charts.PositionCenter,
		}),
		charts.HeightOptionFunc(req.Height),
		charts.WidthOptionFunc(req.Width),
		charts.YAxisDataOptionFunc(data.XData),
		func(opt *charts.ChartOption) {
			opt.Theme = req.Theme
			opt.ValueFormatter = func(f float64) string {
				return fmt.Sprintf("%.0f%s", f, req.Metric)
			}
		},
	)

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	buf, err := p.Bytes()
	SetHeaders(c.Response().Header())
	return buf, err
}
