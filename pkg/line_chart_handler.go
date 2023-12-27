package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	charts "github.com/vicanso/go-charts/v2"
)

type LineChartHandler struct {
	chart *LineChart
}

func NewLineChartHandler() *LineChartHandler {
	return &LineChartHandler{
		chart: NewLineChart(),
	}
}

type LineChartRequest struct {
	ChartData     string `json:"data" query:"data" form:"data" validate:"required" message:"data is required"`
	ChartTitle    string `json:"title" query:"title" form:"title"`
	ChartSubtitle string `json:"subtitle" query:"subtitle" form:"subtitle"`
	Metric        string `json:"metric" query:"metric" form:"metric"`
	Fill          bool   `json:"fill" query:"fill" form:"fill"`
	Theme         string `json:"theme" query:"theme" form:"theme"`
	Color         string `json:"color" query:"color" form:"color"`
	Height        int    `json:"height" query:"height" form:"height"`
	Width         int    `json:"width" query:"width" form:"width"`
}

type LineChartData struct {
	XData [][]string  `json:"x"`
	YData [][]float64 `json:"y"`
	Names []string    `json:"names"`
}

func (h *LineChartHandler) Get(c echo.Context) ([]byte, error) {
	req := new(LineChartRequest)
	if err := BindRequest(c, req); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	var data LineChartData
	if err := json.Unmarshal([]byte(req.ChartData), &data); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	p, err := charts.LineRender(
		data.YData,
		charts.HeightOptionFunc(req.Height),
		charts.WidthOptionFunc(req.Width),
		charts.TitleOptionFunc(charts.TitleOption{
			Text:            req.ChartTitle,
			Subtext:         req.ChartSubtitle,
			SubtextFontSize: 9,
			Left:            charts.PositionCenter,
		}),
		charts.XAxisDataOptionFunc(data.XData[0]),
		charts.LegendOptionFunc(charts.LegendOption{
			Orient: charts.OrientVertical,
			Data:   data.Names,
			Left:   charts.PositionLeft,
		}),
		func(opt *charts.ChartOption) {
			opt.Theme = req.Theme
			opt.Legend.Padding = charts.Box{
				Top:    25,
				Bottom: 25,
			}
			opt.ValueFormatter = func(f float64) string {
				return fmt.Sprintf("%.0f%s", f, req.Metric)
			}
			opt.FillArea = req.Fill

			idx := len(opt.SeriesList) - 1
			if len(opt.SeriesList) > 1 {
				opt.SeriesList[idx].MarkPoint = charts.NewMarkPoint(
					charts.SeriesMarkDataTypeMax,
					charts.SeriesMarkDataTypeMin,
				)
			} else {
				opt.SeriesList[idx].MarkPoint = charts.NewMarkPoint(
					charts.SeriesMarkDataTypeMax,
					charts.SeriesMarkDataTypeMax,
				)
			}

			opt.SeriesList[idx].MarkLine = charts.NewMarkLine(
				charts.SeriesMarkDataTypeAverage,
			)
		},
	)

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	buf, err := p.Bytes()
	if err != nil {
		return nil, err
	}
	return buf, err
}
