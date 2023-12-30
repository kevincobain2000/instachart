package pkg

import (
	"fmt"

	charts "github.com/vicanso/go-charts/v2"
)

type LineChart struct {
	chart *Chart
}

func NewLineChart() *LineChart {
	return &LineChart{
		chart: NewChart(),
	}
}

func (c *LineChart) Get(xData [][]string, yData [][]float64, names []string, req *ChartRequest) ([]byte, error) {
	p, err := charts.LineRender(
		yData,
		charts.HeightOptionFunc(req.Height),
		charts.WidthOptionFunc(req.Width),
		charts.TitleOptionFunc(charts.TitleOption{
			Text:            req.ChartTitle,
			Subtext:         req.ChartSubtitle,
			SubtextFontSize: DEFAULT_SUBTITLE_FONT_SIZE,
			Left:            charts.PositionCenter,
		}),
		charts.XAxisDataOptionFunc(xData[0]),
		charts.LegendOptionFunc(charts.LegendOption{
			Orient: charts.OrientHorizontal,
			Data:   names,
			Left:   charts.PositionLeft,
		}),
		func(opt *charts.ChartOption) {
			opt.Theme = req.Theme
			opt.Legend.Padding = charts.Box{
				Top:    DEFAULT_PADDING_TOP,
				Bottom: DEFAULT_PADDING_BOTTOM * 2,
			}
			opt.ValueFormatter = func(f float64) string {
				return fmt.Sprintf("%.0f%s", f, req.Metric)
			}
			opt.FillArea = req.Fill

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
