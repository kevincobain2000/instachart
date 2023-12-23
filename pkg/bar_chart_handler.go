package pkg

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wcharczuk/go-chart/v2"
)

type BarChartHandler struct {
	chart *BarChart
}

func NewBarChartHandler() *BarChartHandler {
	return &BarChartHandler{
		chart: NewBarChart(),
	}
}

type BarChartRequest struct {
	ChartData  string  `json:"data" query:"data" form:"data" validate:"required" message:"data is required"`
	YAxisLabel string  `json:"y_label" query:"y_label" form:"y_label"`
	ChartTitle string  `json:"title" query:"title" form:"title"`
	BaseValue  float64 `json:"base_value" query:"base_value" form:"base_value"`
	Height     int     `json:"height" query:"height" form:"height"`
	Width      int     `json:"width" query:"width" form:"width"`
}

type BarChartData struct {
	XData []string  `json:"x"`
	YData []float64 `json:"y"`
}

func (h *BarChartHandler) Get(c echo.Context) ([]byte, error) {
	req := new(BarChartRequest)
	if err := BindRequest(c, req); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	var data BarChartData
	if err := json.Unmarshal([]byte(req.ChartData), &data); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if len(data.XData) == 0 || len(data.XData) != len(data.YData) {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, errors.New("data: invalid data"))
	}

	baseValue := req.BaseValue
	useBaseValue := false
	if baseValue > 0.0 {
		baseValue = req.BaseValue
		useBaseValue = true
	}

	graph := chart.BarChart{
		Title:      req.ChartTitle,
		Height:     req.Height,
		Width:      req.Width,
		Background: h.chart.GetBackground(),
		YAxis: chart.YAxis{
			Name:     req.YAxisLabel,
			AxisType: chart.YAxisSecondary,
		},
		UseBaseValue: useBaseValue,
		BaseValue:    baseValue,
		Bars:         h.chart.GetValues(data.XData, data.YData),
	}
	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return buffer.Bytes(), err
}
