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
func (c *PieChart) GetPadding() charts.Box {
	return charts.Box{
		Top:    20,
		Right:  20,
		Bottom: 20,
		Left:   20,
	}
}
