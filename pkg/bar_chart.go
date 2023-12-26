package pkg

type BarChart struct {
	chart *Chart
}

func NewBarChart() *BarChart {
	return &BarChart{
		chart: NewChart(),
	}
}
