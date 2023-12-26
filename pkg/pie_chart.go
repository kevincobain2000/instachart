package pkg

import "github.com/wcharczuk/go-chart/v2"

type PieChart struct {
	chart *Chart
}

func NewPieChart() *PieChart {
	return &PieChart{
		chart: NewChart(),
	}
}
func (c *PieChart) GetValues(names []string, values []float64) []chart.Value {
	var chartValues []chart.Value
	for i := 0; i < len(names); i++ {
		chartValues = append(chartValues, chart.Value{
			Value: values[i],
			Label: names[i],
		})
	}
	return chartValues
}
