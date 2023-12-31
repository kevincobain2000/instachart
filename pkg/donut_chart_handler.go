package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
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
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}

	var data DonutChartData
	if err := json.Unmarshal([]byte(req.ChartData), &data); err != nil {
		msgs := map[string]string{
			"data": "Invalid JSON",
		}
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
	}

	if len(data.Values) == 0 || len(data.Values) != len(data.Names) {
		msgs := map[string]string{
			"data": "Counts are invalid",
		}
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
	}

	SetHeadersResponseImage(c.Response().Header())
	return h.chart.Get(data.Values, data.Names, req)
}
