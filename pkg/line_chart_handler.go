package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LineChartHandler struct {
	chart *LineChart
}

func NewLineChartHandler() *LineChartHandler {
	return &LineChartHandler{
		chart: NewLineChart(),
	}
}

type LineChartData struct {
	XData [][]string  `json:"x"`
	YData [][]float64 `json:"y"`
	Names []string    `json:"names"`
}

func (h *LineChartHandler) Get(c echo.Context) ([]byte, error) {
	req := new(ChartRequest)
	if err := BindRequest(c, req); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}

	var data LineChartData
	if err := json.Unmarshal([]byte(req.ChartData), &data); err != nil {
		msgs := map[string]string{
			"data": "Invalid JSON",
		}
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
	}

	SetHeaders(c.Response().Header())
	return h.chart.Get(data.XData, data.YData, data.Names, req)
}
