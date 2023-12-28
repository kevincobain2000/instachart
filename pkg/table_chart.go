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
	header := names
	rows := values

	if req.Theme == "dark" {
		charts.SetDefaultTableSetting(charts.TableDarkThemeSetting)
	} else {
		charts.SetDefaultTableSetting(charts.TableLightThemeSetting)
	}
	charts.SetDefaultWidth(req.Width)
	p, err := charts.TableRender(header, rows)
	if err != nil {
		return nil, err
	}
	buf, err := p.Bytes()
	return buf, err
}
