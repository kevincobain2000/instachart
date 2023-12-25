package pkg

import "github.com/wcharczuk/go-chart/v2"

type DonutChart struct {
	chart *Chart
}

func NewDonutChart() *DonutChart {
	return &DonutChart{
		chart: NewChart(),
	}
}
func (c *DonutChart) GetValues(names []string, values []float64) []chart.Value {
	var chartValues []chart.Value
	for i := 0; i < len(names); i++ {
		chartValues = append(chartValues, chart.Value{
			Value: values[i],
			Label: names[i],
		})
	}
	return chartValues
}
