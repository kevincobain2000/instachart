package pkg

import (
	"net/http"

	"github.com/labstack/echo/v4"
	charts "github.com/vicanso/go-charts/v2"
)

type FunnelChart struct {
}

func NewFunnelChart() *FunnelChart {
	return &FunnelChart{}
}

func (c *FunnelChart) Get(values []float64, names []string, req *ChartRequest) ([]byte, error) {
	values, names = c.Sort(values, names)
	p, err := charts.FunnelRender(
		values,
		charts.TitleOptionFunc(charts.TitleOption{
			Text:             req.ChartTitle,
			Subtext:          req.ChartSubtitle,
			SubtextFontSize:  DEFAULT_SUBTITLE_FONT_SIZE,
			SubtextFontColor: DEFAULT_SUBTITLE_COLOR,
			Left:             charts.PositionCenter,
		}),
		charts.HeightOptionFunc(req.Height),
		charts.WidthOptionFunc(req.Width),
		charts.LegendOptionFunc(charts.LegendOption{
			Orient: charts.OrientHorizontal,
			Data:   names,
			Left:   charts.PositionLeft,
		}),
		func(opt *charts.ChartOption) {
			opt.Theme = req.Theme
			opt.Type = req.Output
			opt.Legend.Padding = charts.Box{
				Top:    DEFAULT_PADDING_TOP * 2,
				Bottom: DEFAULT_PADDING_BOTTOM,
			}
		},
	)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	buf, err := p.Bytes()
	return buf, err
}

func (c *FunnelChart) Sort(values []float64, names []string) ([]float64, []string) {
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] < values[j] {
				values[i], values[j] = values[j], values[i]
				names[i], names[j] = names[j], names[i]
			}
		}
	}
	return values, names
}
