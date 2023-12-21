package pkg

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/wcharczuk/go-chart/v2"
)

type LineChartHandler struct {
	chart *LineChart
}

func NewLineChartHandler() *LineChartHandler {
	return &LineChartHandler{
		chart: NewLineChart(),
	}
}

type LineChartRequest struct {
	ChartData  string `json:"data" query:"data" form:"data" validate:"required" message:"data is required"`
	XAxisLabel string `json:"x_label" query:"x_label" form:"x_label"`
	YAxisLabel string `json:"y_label" query:"y_label" form:"y_label"`
	ChartTitle string `json:"title" query:"title" form:"title"`
}

type LineChartData struct {
	XData [][]string  `json:"x"`
	YData [][]float64 `json:"y"`
	Names []string    `json:"names"`
}

func (h *LineChartHandler) Get(c echo.Context) ([]byte, error) {
	req := new(LineChartRequest)
	if err := BindRequest(c, req); err != nil {
		return nil, err
	}

	var data LineChartData
	if err := json.Unmarshal([]byte(req.ChartData), &data); err != nil {
		return nil, err
	}

	if len(data.XData) == 0 || len(data.XData) != len(data.YData) {
		return nil, errors.New("data: invalid data")
	}

	isTimeSeries := h.chart.IsTimeseries(data.XData[0][0])

	var series []chart.Series
	for i := 0; i < len(data.XData); i++ {
		name := "Series " + strconv.Itoa(i+1) + " "
		if len(data.Names) > i {
			name = data.Names[i]
		}

		if isTimeSeries {
			series = append(series, chart.TimeSeries{
				Name:    name,
				Style:   h.chart.GetChartStroke(i),
				XValues: h.chart.GetXValuesAsTime(data.XData[i]),
				YValues: h.chart.GetYValues(data.YData[i]),
			})
		} else {
			series = append(series, chart.ContinuousSeries{
				Name:    name,
				Style:   h.chart.GetChartStroke(i),
				XValues: h.chart.GetXValuesAsFloat(data.XData[i]),
				YValues: h.chart.GetYValues(data.YData[i]),
			})
		}

	}

	graph := chart.Chart{
		Title:      req.ChartTitle,
		Background: h.chart.GetBackground(),
		XAxis:      h.chart.GetXAxis(req.XAxisLabel),
		YAxis:      h.chart.GetYAxis(req.YAxisLabel),
		Series:     series,
	}
	graph.Elements = []chart.Renderable{
		chart.Legend(&graph),
	}
	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	return buffer.Bytes(), err
}
