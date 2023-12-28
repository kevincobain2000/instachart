package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BarChartHandler struct {
	chart *BarChart
}

func NewBarChartHandler() *BarChartHandler {
	return &BarChartHandler{
		chart: NewBarChart(),
	}
}

type BarChartData struct {
	XData []string    `json:"x"`
	YData [][]float64 `json:"y"`
	Names []string    `json:"names"`
}

func (h *BarChartHandler) Get(c echo.Context) ([]byte, error) {
	req := new(ChartRequest)
	if err := BindRequest(c, req); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}

	var data BarChartData
	if err := json.Unmarshal([]byte(req.ChartData), &data); err != nil {
		msgs := map[string]string{
			"data": "Invalid JSON",
		}
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
	}

	if len(data.XData) == 0 || len(data.XData) != len(data.YData[0]) {
		msgs := map[string]string{
			"data": "Counts are invalid",
		}
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
	}
	SetHeaders(c.Response().Header())
	if req.Horizontal {
		return h.chart.GetHorizontal(data.XData, data.YData, data.Names, req)
	}

	return h.chart.GetVertical(data.XData, data.YData, data.Names, req)
}
