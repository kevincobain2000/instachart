package pkg

import (
	"strconv"
	"time"

	"github.com/wcharczuk/go-chart/v2"
)

type LineChart struct {
	chart *Chart
}

func NewLineChart() *LineChart {
	return &LineChart{
		chart: NewChart(),
	}
}

func (c *LineChart) GetBackground() chart.Style {
	return c.chart.GetBackground()
}
func (c *LineChart) GetChartStroke(variant int) chart.Style {
	return c.chart.GetChartStroke(variant)
}

func (c *LineChart) GetXAxis(label string) chart.XAxis {
	return chart.XAxis{
		Name: label,
	}
}
func (c *LineChart) GetYAxis(label string) chart.YAxis {
	return chart.YAxis{
		Name:     label,
		AxisType: chart.YAxisSecondary,
	}
}

func (c *LineChart) GetYValues(data []float64) []float64 {
	var yValues []float64
	yValues = append(yValues, data...)
	return yValues
}

func (c *LineChart) GetSeries(xData [][]string, yData [][]float64, names []string) []chart.Series {
	var series []chart.Series
	isTimeSeries := c.IsTimeseries(xData[0][0])
	for i := 0; i < len(xData); i++ {
		name := "Series " + strconv.Itoa(i+1) + " "
		if len(names) > i {
			name = names[i]
		}

		if isTimeSeries {
			series = append(series, chart.TimeSeries{
				Name:    name,
				Style:   c.GetChartStroke(i),
				XValues: c.GetXValuesAsTime(xData[i]),
				YValues: c.GetYValues(yData[i]),
			})
		} else {
			series = append(series, chart.ContinuousSeries{
				Name:    name,
				Style:   c.GetChartStroke(i),
				XValues: c.GetXValuesAsFloat(xData[i]),
				YValues: c.GetYValues(yData[i]),
			})
		}

	}
	return series
}

func (c *LineChart) GetXValuesAsTime(data []string) []time.Time {
	var xValues []time.Time
	for _, x := range data {
		t, _ := time.Parse("2006-01-02", x)
		xValues = append(xValues, t)
	}
	return xValues
}
func (c *LineChart) GetXValuesAsFloat(data []string) []float64 {
	var xValues []float64
	for _, x := range data {
		if xValue, err := strconv.ParseFloat(x, 32); err == nil {
			xValues = append(xValues, xValue)
		} else {
			xValues = append(xValues, 0)
		}
	}
	return xValues
}

func (c *LineChart) IsTimeseries(str string) bool {
	isTimeSeries := false
	if _, err := time.Parse("2006-01-02", str); err == nil {
		isTimeSeries = true
	}
	return isTimeSeries
}
