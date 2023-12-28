package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FunnelChartHandler struct {
	chart *FunnelChart
}

func NewFunnelChartHandler() *FunnelChartHandler {
	return &FunnelChartHandler{
		chart: NewFunnelChart(),
	}
}

type FunnelChartData struct {
	Names  []string  `json:"names"`
	Values []float64 `json:"values"`
}

func (h *FunnelChartHandler) Get(c echo.Context) ([]byte, error) {
	req := new(ChartRequest)
	if err := BindRequest(c, req); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}

	var data FunnelChartData
	if err := json.Unmarshal([]byte(req.ChartData), &data); err != nil {
		msgs := map[string]string{
			"data": "Invalid JSON",
		}
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
	}

	if len(data.Names) == 0 || len(data.Names) != len(data.Values) {
		msgs := map[string]string{
			"data": "Counts are invalid",
		}
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
	}

	SetHeaders(c.Response().Header())
	return h.chart.Get(data.Values, data.Names, req)
}
