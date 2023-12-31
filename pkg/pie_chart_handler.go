package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mcuadros/go-defaults"
)

type PieChartHandler struct {
	chart                *PieChart
	allowedRemoteDomains string
}

func NewPieChartHandler(allowedRemoteDomains string) *PieChartHandler {
	return &PieChartHandler{
		chart:                NewPieChart(),
		allowedRemoteDomains: allowedRemoteDomains,
	}
}

type PieChartData struct {
	Names  []string  `json:"names"`
	Values []float64 `json:"values"`
}

func (h *PieChartHandler) Get(c echo.Context) ([]byte, error) {
	req := new(ChartRequest)
	defaults.SetDefaults(req)
	if err := BindRequest(c, req); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}
	err := SetDataIfRemoteURL(req, h.allowedRemoteDomains)
	if err != nil {
		msgs := map[string]string{
			"data": err.Error(),
		}
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
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
	SetHeadersResponseImage(c.Response().Header(), req.Output)
	return buf, err
}
