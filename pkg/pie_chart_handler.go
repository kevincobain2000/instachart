package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PieChartHandler struct {
	chart *PieChart
}

func NewPieChartHandler() *PieChartHandler {
	return &PieChartHandler{
		chart: NewPieChart(),
	}
}

type PieChartData struct {
	Names  []string  `json:"names"`
	Values []float64 `json:"values"`
}

func (h *PieChartHandler) Get(c echo.Context) ([]byte, error) {
	req := new(ChartRequest)
	if err := BindRequest(c, req); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}

	var data PieChartData
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

	buf, err := h.chart.Get(data.Values, data.Names, req)
	SetHeadersResponseImage(c.Response().Header())
	return buf, err
}
