package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mcuadros/go-defaults"
)

type RadarChartHandler struct {
	chart                *RadarChart
	allowedRemoteDomains string
}

func NewRadarChartHandler(allowedRemoteDomains string) *RadarChartHandler {
	return &RadarChartHandler{
		chart:                NewRadarChart(),
		allowedRemoteDomains: allowedRemoteDomains,
	}
}

type RadarChartData struct {
	Labels []string    `json:"labels"`
	Names  []string    `json:"names"`
	Values [][]float64 `json:"values"`
}

func (h *RadarChartHandler) Get(c echo.Context) ([]byte, error) {
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

	var data RadarChartData
	if err := json.Unmarshal([]byte(req.ChartData), &data); err != nil {
		msgs := map[string]string{
			"data": "Invalid JSON",
		}
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
	}

	if len(data.Names) == 0 || len(data.Names) != len(data.Values[0]) {
		msgs := map[string]string{
			"data": "Counts are invalid",
		}
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
	}

	SetHeadersResponseImage(c.Response().Header(), req.Output)
	return h.chart.Get(data.Values, data.Names, data.Labels, req)
}
