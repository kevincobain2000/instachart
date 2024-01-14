package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mcuadros/go-defaults"
)

type FunnelChartHandler struct {
	chart                *FunnelChart
	allowedRemoteDomains string
}

func NewFunnelChartHandler(allowedRemoteDomains string) *FunnelChartHandler {
	return &FunnelChartHandler{
		chart:                NewFunnelChart(),
		allowedRemoteDomains: allowedRemoteDomains,
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
	defaults.SetDefaults(req)
	msgs, err := ValidateRequest(req)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
	}
	err = SetDataIfRemoteURL(req, h.allowedRemoteDomains)
	if err != nil {
		msgs := map[string]string{
			"data": err.Error(),
		}
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
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

	SetHeadersResponseImage(c.Response().Header(), req.Output)
	return h.chart.Get(data.Values, data.Names, req)
}
