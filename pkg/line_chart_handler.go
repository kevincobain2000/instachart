package pkg

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	charts "github.com/vicanso/go-charts/v2"
	"github.com/wcharczuk/go-chart/v2"
)

type LineChartHandler struct {
	chart *LineChart
}

func NewLineChartHandler() *LineChartHandler {
	return &LineChartHandler{
		chart: NewLineChart(),
	}
}

type LineChartRequest struct {
	ChartData     string `json:"data" query:"data" form:"data" validate:"required" message:"data is required"`
	XAxisLabel    string `json:"x_label" query:"x_label" form:"x_label"`
	YAxisLabel    string `json:"y_label" query:"y_label" form:"y_label"`
	ChartTitle    string `json:"title" query:"title" form:"title"`
	ChartSubtitle string `json:"subtitle" query:"subtitle" form:"subtitle"`
	Fill          bool   `json:"fill" query:"fill" form:"fill"`
	Color         string `json:"color" query:"color" form:"color"`
	Height        int    `json:"height" query:"height" form:"height"`
	Width         int    `json:"width" query:"width" form:"width"`
}

type LineChartData struct {
	XData [][]string  `json:"x"`
	YData [][]float64 `json:"y"`
	Names []string    `json:"names"`
}

func (h *LineChartHandler) Get(c echo.Context) ([]byte, error) {
	req := new(LineChartRequest)
	if err := BindRequest(c, req); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	var data LineChartData
	if err := json.Unmarshal([]byte(req.ChartData), &data); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	isSeries := h.chart.ContainsDateOrFloat(data.XData)

	if isSeries {
		if len(data.XData) == 0 || len(data.XData) != len(data.YData) {
			return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, "data provided is invalid")
		}
		graph := chart.Chart{
			Title:      req.ChartTitle,
			Height:     req.Height,
			Width:      req.Width,
			Background: h.chart.GetBackground(),
			XAxis:      h.chart.GetXAxis(req.XAxisLabel),
			YAxis:      h.chart.GetYAxis(req.YAxisLabel),
			Series:     h.chart.GetSeries(data.XData, data.YData, data.Names, req.Color, req.Fill),
		}
		graph.Elements = []chart.Renderable{
			chart.Legend(&graph),
		}
		buffer := bytes.NewBuffer([]byte{})
		err := graph.Render(chart.PNG, buffer)
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}
		return buffer.Bytes(), err
	}
	p, err := charts.LineRender(
		data.YData,
		charts.HeightOptionFunc(req.Height),
		charts.WidthOptionFunc(req.Width),
		charts.TitleOptionFunc(charts.TitleOption{
			Text:            req.ChartTitle,
			Subtext:         req.ChartSubtitle,
			SubtextFontSize: 9,
			Left:            charts.PositionCenter,
		}),
		charts.XAxisDataOptionFunc(data.XData[0]),
		charts.LegendOptionFunc(charts.LegendOption{
			Orient: charts.OrientVertical,
			Data:   data.Names,
			Left:   charts.PositionLeft,
		}),
		func(opt *charts.ChartOption) {
			opt.FillArea = req.Fill
			opt.SeriesList[1].MarkPoint = charts.NewMarkPoint(
				charts.SeriesMarkDataTypeMax,
				charts.SeriesMarkDataTypeMin,
			)
			opt.SeriesList[1].MarkLine = charts.NewMarkLine(
				charts.SeriesMarkDataTypeAverage,
			)
		},
	)

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	buf, err := p.Bytes()
	if err != nil {
		return nil, err
	}
	return buf, err
}
