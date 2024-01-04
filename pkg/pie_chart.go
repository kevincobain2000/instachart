package pkg

import (
	charts "github.com/vicanso/go-charts/v2"
)

type PieChart struct {
	chart *Chart
}

func NewPieChart() *PieChart {
	return &PieChart{
		chart: NewChart(),
	}
}

func (c *PieChart) Get(values []float64, names []string, req *ChartRequest) ([]byte, error) {
	p, err := charts.PieRender(
		values,
		charts.TitleOptionFunc(charts.TitleOption{
			Text:            req.ChartTitle,
			Subtext:         req.ChartSubtitle,
			SubtextFontSize: DEFAULT_SUBTITLE_FONT_SIZE,
			Left:            charts.PositionCenter,
		}),
		charts.HeightOptionFunc(req.Height),
		charts.WidthOptionFunc(req.Width),
		charts.PaddingOptionFunc(charts.Box{
			Top:    DEFAULT_PADDING_TOP,
			Right:  DEFAULT_PADDING_RIGHT,
			Bottom: DEFAULT_PADDING_BOTTOM,
			Left:   DEFAULT_PADDING_LEFT,
		}),
		charts.LegendOptionFunc(charts.LegendOption{
			Orient: charts.OrientVertical,
			Data:   names,
			Left:   charts.PositionLeft,
		}),
		charts.PieSeriesShowLabel(),
		func(opt *charts.ChartOption) {
			opt.Theme = req.Theme
			opt.Type = req.Output
		},
	)
	if err != nil {
		return nil, err
	}

	buf, err := p.Bytes()
	return buf, err
}
