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

func (c *LineChart) Get(xData []string, yData [][]float64, names []string, req *ChartRequest) ([]byte, error) {
	fill := false
	if req.Line == "fill" {
		fill = true
	}
	isMini := IsMiniChart(req)

	showLegend := true
	paddings := GetPaddings(req)
	titleSizes := GetTitleSizes(req)
	if isMini {
		showLegend = false
	}
	p, err := charts.LineRender(
		yData,
		charts.HeightOptionFunc(req.Height),
		charts.WidthOptionFunc(req.Width),
		charts.PaddingOptionFunc(paddings),
		charts.TitleOptionFunc(titleSizes),
		charts.XAxisDataOptionFunc(xData),
		charts.LegendOptionFunc(charts.LegendOption{
			Orient: charts.OrientHorizontal,
			Data:   names,
			Left:   charts.PositionLeft,
			Show:   &showLegend,
		}),
		func(opt *charts.ChartOption) {
			opt.Type = req.Output
			opt.Theme = req.Theme
			opt.BackgroundColor = DEFAULT_BACKGROUND_COLOR
			opt.Legend.Padding = charts.Box{
				Top:    DEFAULT_PADDING_TOP * 2,
				Bottom: DEFAULT_PADDING_BOTTOM,
			}
			opt.ValueFormatter = func(f float64) string {
				if isMini {
					return "-"
				}
				return fmt.Sprintf("%s %s", NumberToK(&f), req.Metric)

			}
			opt.FillArea = fill

			idx := len(opt.SeriesList) - 1
			if len(opt.SeriesList) > 1 {
				idx = 1
			}
			opt.SeriesList[idx].MarkPoint = charts.NewMarkPoint(
				charts.SeriesMarkDataTypeMax,
				charts.SeriesMarkDataTypeMin,
			)
			if !isMini {
				opt.SeriesList[idx].MarkLine = charts.NewMarkLine(
					charts.SeriesMarkDataTypeAverage,
				)
			}

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
