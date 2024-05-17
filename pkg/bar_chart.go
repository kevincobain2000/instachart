package pkg

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	charts "github.com/vicanso/go-charts/v2"
)

type BarChart struct {
	chart *Chart
}

func NewBarChart() *BarChart {
	return &BarChart{
		chart: NewChart(),
	}
}

func (c *BarChart) GetVertical(xData []string, yData [][]float64, names []string, req *ChartRequest) ([]byte, error) {
	isMini := IsMiniChart(req)
	showLegend := true
	paddings := GetPaddings(req)
	titleSizes := GetTitleSizes(req)
	showGrid := charts.TrueFlag()
	if req.Grid == "hide" {
		showGrid = charts.FalseFlag()
	}
	if isMini {
		showLegend = false
	}
	p, err := charts.BarRender(
		yData,
		charts.TitleOptionFunc(titleSizes),
		charts.HeightOptionFunc(req.Height),
		charts.WidthOptionFunc(req.Width),
		charts.PaddingOptionFunc(paddings),
		charts.XAxisDataOptionFunc(xData),
		charts.LegendOptionFunc(charts.LegendOption{
			Orient: charts.OrientHorizontal,
			Data:   names,
			Left:   charts.PositionLeft,
			Show:   &showLegend,
		}),
		charts.MarkLineOptionFunc(0, charts.SeriesMarkDataTypeAverage),
		charts.MarkPointOptionFunc(0, charts.SeriesMarkDataTypeMax,
			charts.SeriesMarkDataTypeMin),
		func(opt *charts.ChartOption) {
			opt.Theme = req.Theme
			opt.Type = req.Output
			opt.Legend.Padding = charts.Box{
				Top:    DEFAULT_PADDING_TOP * 2,
				Bottom: DEFAULT_PADDING_BOTTOM,
			}
			opt.ValueFormatter = func(f float64) string {
				if isMini {
					return "-"
				}
				return fmt.Sprintf("%s %s", NumberToK(&f), req.Metric)

			}
			idx := len(opt.SeriesList) - 1
			if len(opt.SeriesList) > 1 {
				idx = 1
			}
			for series := range opt.SeriesList {
				opt.SeriesList[series].RoundRadius = 10
			}
			opt.SeriesList[idx].MarkPoint = charts.NewMarkPoint(
				charts.SeriesMarkDataTypeMax,
				charts.SeriesMarkDataTypeMin,
			)
			opt.YAxisOptions = []charts.YAxisOption{
				{
					SplitLineShow: showGrid,
				},
			}
			if !isMini {
				opt.SeriesList[idx].MarkLine = charts.NewMarkLine(
					charts.SeriesMarkDataTypeAverage,
				)
			}
		},
	)
	if err != nil {
		return nil, err
	}

	buf, err := p.Bytes()
	if err != nil {
		return nil, err
	}
	return buf, err
}

func (c *BarChart) GetStacked(xData []string, yData [][]float64, zData [][]float64, names []string, req *ChartRequest) ([]byte, error) {
	showGrid := charts.TrueFlag()
	if req.Grid == "hide" {
		showGrid = charts.FalseFlag()
	}
	series := make([]charts.Series, 0)
	for _, y := range yData {
		s := charts.Series{
			Type: charts.ChartTypeBar,
			Data: charts.NewSeriesDataFromValues(y),
		}
		series = append(series, s)
	}

	for idx, z := range zData {
		s := charts.Series{
			Data:      charts.NewSeriesDataFromValues(z),
			AxisIndex: idx,
		}
		series = append(series, s)
	}
	titleSizes := GetTitleSizes(req)
	opt := charts.ChartOption{
		Title:      titleSizes,
		XAxis:      charts.NewXAxisOption(xData),
		Legend:     charts.NewLegendOption(names),
		SeriesList: series,
	}
	opt.YAxisOptions = []charts.YAxisOption{
		{
			SplitLineShow: showGrid,
		},
	}
	opt.ValueFormatter = func(f float64) string {
		return fmt.Sprintf("%s %s", NumberToK(&f), req.Metric)
	}
	opt.Type = req.Output
	opt.Theme = req.Theme
	opt.Legend.Padding = charts.Box{
		Top:    DEFAULT_PADDING_TOP * 2,
		Bottom: DEFAULT_PADDING_BOTTOM,
	}
	opt.Legend.Orient = charts.OrientHorizontal
	opt.Legend.Left = charts.PositionLeft
	opt.Width = req.Width
	opt.Height = req.Height

	idx := len(opt.SeriesList) - 1
	if len(opt.SeriesList) > 1 {
		idx = 1
	}
	for series := range opt.SeriesList {
		opt.SeriesList[series].RoundRadius = 10
	}
	opt.SeriesList[idx].MarkPoint = charts.NewMarkPoint(
		charts.SeriesMarkDataTypeMax,
		charts.SeriesMarkDataTypeMin,
	)
	opt.SeriesList[idx].MarkLine = charts.NewMarkLine(
		charts.SeriesMarkDataTypeAverage,
	)

	d, err := charts.Render(opt)
	if err != nil {
		return nil, err
	}
	return d.Bytes()
}

func (c *BarChart) GetHorizontal(xData []string, yData [][]float64, names []string, req *ChartRequest) ([]byte, error) {
	isMini := IsMiniChart(req)
	showLegend := true
	paddings := GetPaddings(req)
	titleSizes := GetTitleSizes(req)

	showGrid := charts.TrueFlag()
	if req.Grid == "hide" {
		showGrid = charts.FalseFlag()
	}
	if isMini {
		showLegend = false
	}
	p, err := charts.HorizontalBarRender(
		yData,
		charts.TitleOptionFunc(titleSizes),
		charts.HeightOptionFunc(req.Height),
		charts.WidthOptionFunc(req.Width),
		charts.PaddingOptionFunc(paddings),
		charts.YAxisDataOptionFunc(xData),
		charts.LegendOptionFunc(charts.LegendOption{
			Orient: charts.OrientHorizontal,
			Data:   names,
			Left:   charts.PositionLeft,
			Show:   &showLegend,
		}),
		func(opt *charts.ChartOption) {
			opt.Theme = req.Theme
			opt.Type = req.Output
			opt.YAxisOptions = []charts.YAxisOption{
				{
					SplitLineShow: showGrid,
				},
			}
			opt.ValueFormatter = func(f float64) string {
				if isMini {
					return "-"
				}
				return fmt.Sprintf("%s %s", NumberToK(&f), req.Metric)
			}
		},
	)

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	buf, err := p.Bytes()
	return buf, err
}
