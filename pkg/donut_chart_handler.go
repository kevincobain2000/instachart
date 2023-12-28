package pkg

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wcharczuk/go-chart/v2"
)

type DonutChartHandler struct {
	chart *DonutChart
}

func NewDonutChartHandler() *DonutChartHandler {
	return &DonutChartHandler{
		chart: NewDonutChart(),
	}
}

type DonutChartData struct {
	Names  []string  `json:"names"`
	Values []float64 `json:"values"`
}

func (h *DonutChartHandler) Get(c echo.Context) ([]byte, error) {
	req := new(ChartRequest)
	if err := BindRequest(c, req); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	var data DonutChartData
	if err := json.Unmarshal([]byte(req.ChartData), &data); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if len(data.Values) == 0 || len(data.Values) != len(data.Names) {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, "data provided is invalid")
	}

	graph := chart.DonutChart{
		Title:  req.ChartTitle,
		Height: req.Height,
		Width:  req.Width,
		Values: h.chart.GetValues(data.Names, data.Values),
	}

	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	SetHeaders(c.Response().Header())
	return buffer.Bytes(), err
}
