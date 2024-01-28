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
	fill := true
	if req.Line == "fill" {
		fill = true
	}
	isMini := false
	showLegend := true
	paddings := charts.Box{
		Top:    10,
		Bottom: 10,
		Left:   10,
		Right:  10,
	}
	titleSizes := charts.TitleOption{
		Text:             req.ChartTitle,
		Subtext:          req.ChartSubtitle,
		FontSize:         DEFAULT_TITLE_FONT_SIZE,
		SubtextFontSize:  DEFAULT_SUBTITLE_FONT_SIZE,
		Left:             charts.PositionCenter,
		SubtextFontColor: DEFAULT_SUBTITLE_COLOR,
	}
	if req.Width <= 300 && req.Height <= 300 {
		showLegend = false
		isMini = true
		paddings = charts.Box{
			Top:    10,
			Bottom: -20,
			Left:   -10,
			Right:  10,
		}
		titleSizes = charts.TitleOption{
			Text:             Truncate(req.ChartTitle, 17),
			Subtext:          Truncate(req.ChartSubtitle, 17),
			FontSize:         DEFAULT_TITLE_FONT_SIZE,
			SubtextFontSize:  DEFAULT_SUBTITLE_FONT_SIZE,
			Left:             charts.PositionCenter,
			SubtextFontColor: DEFAULT_SUBTITLE_COLOR,
		}
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
			opt.Legend.Padding = charts.Box{
				Top:    0,
				Bottom: 0,
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
