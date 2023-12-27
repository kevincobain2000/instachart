package pkg

type LineChart struct {
	chart *Chart
}

func NewLineChart() *LineChart {
	return &LineChart{
		chart: NewChart(),
	}
}
