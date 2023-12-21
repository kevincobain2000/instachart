package pkg

import (
	"github.com/wcharczuk/go-chart/v2"
)

type BarChart struct {
	chart *Chart
}

func NewBarChart() *BarChart {
	return &BarChart{
		chart: NewChart(),
	}
}

func (c *BarChart) GetBackground() chart.Style {
	return c.chart.GetBackground()
}
func (c *BarChart) GetValues(xData []string, yData []float64) []chart.Value {
	var values []chart.Value
	min := yData[0]
	max := yData[0]
	sum := 0.0
	for i := 0; i < len(xData); i++ {
		sum += yData[i]
		values = append(values, chart.Value{
			Value: yData[i],
			Label: xData[i],
		})
		if yData[i] < min {
			min = yData[i]
		}
		if yData[i] > max {
			max = yData[i]
		}
	}
	return values
}
