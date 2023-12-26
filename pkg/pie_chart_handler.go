package pkg

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wcharczuk/go-chart/v2"
)

type PieChartHandler struct {
	chart *PieChart
}

func NewPieChartHandler() *PieChartHandler {
	return &PieChartHandler{
		chart: NewPieChart(),
	}
}

type PieChartRequest struct {
	ChartData  string `json:"data" query:"data" form:"data" validate:"required" message:"data is required"`
	ChartTitle string `json:"title" query:"title" form:"title"`
	Height     int    `json:"height" query:"height" form:"height"`
	Width      int    `json:"width" query:"width" form:"width"`
}

type PieChartData struct {
	Names  []string  `json:"names"`
	Values []float64 `json:"values"`
}

func (h *PieChartHandler) Get(c echo.Context) ([]byte, error) {
	req := new(PieChartRequest)
	if err := BindRequest(c, req); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	var data PieChartData
	if err := json.Unmarshal([]byte(req.ChartData), &data); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if len(data.Values) == 0 || len(data.Values) != len(data.Names) {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, "data provided is invalid")
	}

	graph := chart.PieChart{
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
	return buffer.Bytes(), err
}
