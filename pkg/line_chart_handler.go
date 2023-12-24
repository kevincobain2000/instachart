package pkg

import (
	"bytes"
	"encoding/json"
	"net/http"

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
	Height     int    `json:"height" query:"height" form:"height"`
	Width      int    `json:"width" query:"width" form:"width"`
}

type LineChartData struct {
	XData [][]string  `json:"x"`
	YData [][]float64 `json:"y"`
	Names []string    `json:"names"`
}

func (h *LineChartHandler) Get(c echo.Context) ([]byte, error) {
	req := new(LineChartRequest)
	if err := BindRequest(c, req); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	var data LineChartData
	if err := json.Unmarshal([]byte(req.ChartData), &data); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if len(data.XData) == 0 || len(data.XData) != len(data.YData) {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, "data provided is invalid")
	}

	graph := chart.Chart{
		Title:      req.ChartTitle,
		Height:     req.Height,
		Width:      req.Width,
		Background: h.chart.GetBackground(),
		XAxis:      h.chart.GetXAxis(req.XAxisLabel),
		YAxis:      h.chart.GetYAxis(req.YAxisLabel),
		Series:     h.chart.GetSeries(data.XData, data.YData, data.Names),
	}
	graph.Elements = []chart.Renderable{
		chart.Legend(&graph),
	}
	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return buffer.Bytes(), err
}
