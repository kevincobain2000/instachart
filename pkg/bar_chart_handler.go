package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mcuadros/go-defaults"
)

type BarChartHandler struct {
	chart                *BarChart
	allowedRemoteDomains string
}

func NewBarChartHandler(allowedRemoteDomains string) *BarChartHandler {
	return &BarChartHandler{
		chart:                NewBarChart(),
		allowedRemoteDomains: allowedRemoteDomains,
	}
}

type BarChartData struct {
	XData []string    `json:"x"`
	YData [][]float64 `json:"y"`
	ZData [][]float64 `json:"z"`
	Names []string    `json:"names"`
}

func (h *BarChartHandler) Get(c echo.Context) ([]byte, error) {
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

	var data BarChartData
	if err := json.Unmarshal([]byte(req.ChartData), &data); err != nil {
		msgs := map[string]string{
			"data": "Invalid JSON",
		}
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
	}

	if len(data.XData) == 0 || len(data.YData) == 0 {
		msgs := map[string]string{
			"data": "X or Y Counts are invalid",
		}
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
	}
	for _, y := range data.YData {
		if len(y) != len(data.XData) {
			msgs := map[string]string{
				"data": "Y Counts are invalid",
			}
			return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
		}
	}
	if len(data.ZData) != 0 {
		for _, z := range data.ZData {
			if len(z) != len(data.XData) {
				msgs := map[string]string{
					"data": "Z Counts are invalid",
				}
				return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
			}
		}
		req.Style = BAR_STYLE_STACKED
	}

	SetHeadersResponseImage(c.Response().Header(), req.Output)
	switch req.Style {
	case BAR_STYLE_VERTICAL:
		return h.chart.GetVertical(data.XData, data.YData, data.Names, req)
	case BAR_STYLE_HORIZONTAL:
		return h.chart.GetHorizontal(data.XData, data.YData, data.Names, req)
	case BAR_STYLE_STACKED:
		return h.chart.GetStacked(data.XData, data.YData, data.ZData, data.Names, req)
	}

	return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid style")

}
