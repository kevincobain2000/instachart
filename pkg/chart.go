package pkg

import (
	"net/http"
)

const (
	DEFAULT_PADDING_TOP        = 20
	DEFAULT_PADDING_RIGHT      = 20
	DEFAULT_PADDING_BOTTOM     = 20
	DEFAULT_PADDING_LEFT       = 20
	DEFAULT_SUBTITLE_FONT_SIZE = 9
)

type Chart struct {
}

func NewChart() *Chart {
	return &Chart{}
}

type ChartRequest struct {
	ChartData     string `json:"data" query:"data" form:"data" validate:"required" message:"data is required"`
	ChartTitle    string `json:"title" query:"title" form:"title"`
	ChartSubtitle string `json:"subtitle" query:"subtitle" form:"subtitle"`
	Metric        string `json:"metric" query:"metric" form:"metric"`
	Height        int    `json:"height" query:"height" form:"height"`
	Theme         string `json:"theme" query:"theme" form:"theme"`
	Type          string `json:"type" query:"type" form:"type"`
	Width         int    `json:"width" query:"width" form:"width"`
	Horizontal    bool   `json:"horizontal" query:"horizontal" form:"horizontal"`
	Fill          bool   `json:"fill" query:"fill" form:"fill"`
}

func SetHeaders(header http.Header) {
	header.Set("Cache-Control", "max-age=31536000")
	header.Set("Expires", "31536000")
	header.Set("Content-Type", "image/png")
}
