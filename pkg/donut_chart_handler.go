package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mcuadros/go-defaults"
)

type DonutChartHandler struct {
	chart                *DonutChart
	allowedRemoteDomains string
}

func NewDonutChartHandler(allowedRemoteDomains string) *DonutChartHandler {
	return &DonutChartHandler{
		chart:                NewDonutChart(),
		allowedRemoteDomains: allowedRemoteDomains,
	}
}

type DonutChartData struct {
	Names  []string  `json:"names"`
	Values []float64 `json:"values"`
}

func (h *DonutChartHandler) Get(c echo.Context) ([]byte, error) {
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

	SetHeadersResponseImage(c.Response().Header(), req.Output)
	return h.chart.Get(data.Values, data.Names, req)
}
