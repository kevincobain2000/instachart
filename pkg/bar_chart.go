package pkg

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	charts "github.com/vicanso/go-charts/v2"
)

type BarChart struct {
	chart *Chart
}

func NewBarChart() *BarChart {
	return &BarChart{
		chart: NewChart(),
	}
}

func (c *BarChart) GetVertical(xData []string, yData [][]float64, names []string, req *ChartRequest) ([]byte, error) {
	p, err := charts.BarRender(
		yData,
		charts.TitleOptionFunc(charts.TitleOption{
			Text:            req.ChartTitle,
			Subtext:         req.ChartSubtitle,
			SubtextFontSize: DEFAULT_SUBTITLE_FONT_SIZE,
			Left:            charts.PositionCenter,
		}),
		charts.HeightOptionFunc(req.Height),
		charts.WidthOptionFunc(req.Width),
		charts.XAxisDataOptionFunc(xData),
		charts.LegendOptionFunc(charts.LegendOption{
			Orient: charts.OrientVertical,
			Data:   names,
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
		return nil, err
	}

	buf, err := p.Bytes()
	if err != nil {
		return nil, err
	}
	return buf, err
}

func (c *BarChart) GetHorizontal(xData []string, yData [][]float64, names []string, req *ChartRequest) ([]byte, error) {
	p, err := charts.HorizontalBarRender(
		yData,
		charts.TitleOptionFunc(charts.TitleOption{
			Text:            req.ChartTitle,
			Subtext:         req.ChartSubtitle,
			SubtextFontSize: DEFAULT_SUBTITLE_FONT_SIZE,
			Left:            charts.PositionCenter,
		}),
		charts.HeightOptionFunc(req.Height),
		charts.WidthOptionFunc(req.Width),
		charts.YAxisDataOptionFunc(xData),
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
	return buf, err
}
