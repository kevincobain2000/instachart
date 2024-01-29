package pkg

import (
	"net/http"

	"github.com/labstack/echo/v4"
	charts "github.com/vicanso/go-charts/v2"
)

type RadarChart struct {
}

func NewRadarChart() *RadarChart {
	return &RadarChart{}
}

func (c *RadarChart) GetIndicators(values [][]float64) []float64 {
	if len(values) == 0 {
		return nil
	}

	// Initialize maxValues with the first set of values
	maxValues := make([]float64, len(values[0]))
	copy(maxValues, values[0])

	// Iterate over each set of values
	for _, set := range values {
		for i, value := range set {
			if value > maxValues[i] {
				maxValues[i] = value
			}
		}
	}
	return maxValues
}

func (c *RadarChart) Get(values [][]float64, names []string, labels []string, req *ChartRequest) ([]byte, error) {
	p, err := charts.RadarRender(
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
			Orient: charts.OrientVertical,
			Data:   labels,
			Left:   charts.PositionLeft,
		}),
		charts.RadarIndicatorOptionFunc(names, c.GetIndicators(values)),
		func(opt *charts.ChartOption) {
			opt.Theme = req.Theme
			opt.BackgroundColor = DEFAULT_BACKGROUND_COLOR
			opt.Type = req.Output
			opt.Legend.Padding = charts.Box{
				Top:    DEFAULT_PADDING_TOP,
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
