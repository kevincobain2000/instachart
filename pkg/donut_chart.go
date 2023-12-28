package pkg

import (
	"bytes"

	"github.com/wcharczuk/go-chart/v2"
)

type DonutChart struct {
	chart *Chart
}

func NewDonutChart() *DonutChart {
	return &DonutChart{
		chart: NewChart(),
	}
}

func (c *DonutChart) Get(values []float64, names []string, req *ChartRequest) ([]byte, error) {
	var chartValues []chart.Value
	for i := 0; i < len(names); i++ {
		chartValues = append(chartValues, chart.Value{
			Value: values[i],
			Label: names[i],
		})
	}
	graph := chart.DonutChart{
		Title:  req.ChartTitle,
		Height: req.Height,
		Width:  req.Width,
		Values: chartValues,
	}

	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	return buffer.Bytes(), err
}
