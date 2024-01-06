package pkg

import (
	charts "github.com/vicanso/go-charts/v2"
)

type TableChart struct {
}

func NewTableChart() *TableChart {
	return &TableChart{}
}

func (c *TableChart) Get(names []string, values [][]string, req *ChartRequest) ([]byte, error) {
	opts := charts.TableChartOption{
		Type:   req.Output,
		Header: names,
		Data:   values,
		Width:  req.Width,
	}

	if req.Theme == "dark" {
		charts.SetDefaultTableSetting(charts.TableDarkThemeSetting)
	} else {
		charts.SetDefaultTableSetting(charts.TableLightThemeSetting)
	}
	p, err := charts.TableOptionRender(opts)

	if err != nil {
		return nil, err
	}
	buf, err := p.Bytes()
	return buf, err
}
