package pkg

import (
	"github.com/wcharczuk/go-chart/v2"
)

type Chart struct {
}

func NewChart() *Chart {
	return &Chart{}
}

func (c *Chart) GetBackground() chart.Style {
	return chart.Style{
		Padding: chart.Box{
			Top:    30,
			Left:   30,
			Bottom: 30,
			Right:  30,
		},
	}
}

func (c *Chart) GetChartStroke(variant int, color string, fill bool) chart.Style {
	var alpha uint8 = 80
	strokeColor := chart.ColorBlue
	fillColor := chart.ColorBlue.WithAlpha(alpha)
	if color != "" {
		switch color {
		case "red":
			strokeColor = chart.ColorRed
			fillColor = chart.ColorRed.WithAlpha(alpha)
		case "green":
			strokeColor = chart.ColorGreen
			fillColor = chart.ColorGreen.WithAlpha(alpha)
		case "yellow":
			strokeColor = chart.ColorYellow
			fillColor = chart.ColorYellow.WithAlpha(alpha)
		case "black":
			strokeColor = chart.ColorBlack
			fillColor = chart.ColorBlack.WithAlpha(alpha)
		case "cyan":
			strokeColor = chart.ColorCyan
			fillColor = chart.ColorCyan.WithAlpha(alpha)
		case "orange":
			strokeColor = chart.ColorOrange
			fillColor = chart.ColorOrange.WithAlpha(alpha)
		case "blue":
			strokeColor = chart.ColorBlue
			fillColor = chart.ColorBlue.WithAlpha(alpha)
		}
	}
	switch variant {
	case 1:
		strokeColor = chart.ColorRed
		fillColor = chart.ColorRed.WithAlpha(alpha + 10)
	case 2:
		strokeColor = chart.ColorGreen
		fillColor = chart.ColorGreen.WithAlpha(alpha + 20)
	case 3:
		strokeColor = chart.ColorYellow
		fillColor = chart.ColorYellow.WithAlpha(alpha + 30)
	case 4:
		strokeColor = chart.ColorBlack
		fillColor = chart.ColorBlack.WithAlpha(alpha + 40)
	case 5:
		strokeColor = chart.ColorCyan
		fillColor = chart.ColorCyan.WithAlpha(alpha + 50)
	case 6:
		strokeColor = chart.ColorOrange
		fillColor = chart.ColorOrange.WithAlpha(alpha + 60)
	case 7:
		strokeColor = chart.ColorAlternateBlue
		fillColor = chart.ColorAlternateBlue.WithAlpha(alpha + 70)
	case 8:
		strokeColor = chart.ColorAlternateLightGray
		fillColor = chart.ColorAlternateLightGray.WithAlpha(alpha + 80)
	case 9:
		strokeColor = chart.ColorAlternateGreen
		fillColor = chart.ColorAlternateGreen.WithAlpha(alpha + 90)
	case 10:
		strokeColor = chart.ColorAlternateYellow
		fillColor = chart.ColorAlternateYellow.WithAlpha(alpha + 100)
	}

	style := chart.Style{}
	style.StrokeColor = strokeColor
	style.StrokeWidth = 2
	if fill {
		style.FillColor = fillColor
	}

	return style
}
