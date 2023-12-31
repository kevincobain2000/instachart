package pkg

import (
	"net/http"
)

const (
	DEFAULT_PADDING_TOP        = 20
	DEFAULT_PADDING_RIGHT      = 20
	DEFAULT_PADDING_BOTTOM     = 20
	DEFAULT_PADDING_LEFT       = 20
	DEFAULT_SUBTITLE_FONT_SIZE = 10
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

func SetHeadersResponseImage(header http.Header) {
	header.Set("Cache-Control", "max-age=31536000")
	header.Set("Expires", "31536000")
	header.Set("Content-Type", "image/png")
	// security headers
	header.Set("X-Content-Type-Options", "nosniff")
	header.Set("X-Frame-Options", "DENY")
	header.Set("X-XSS-Protection", "1; mode=block")
	// content policy
	header.Set("Content-Security-Policy", "default-src 'none'; img-src 'self'; style-src 'self'; font-src 'self'; connect-src 'self'; script-src 'self';")
}
func SetHeadersResponseHTML(header http.Header) {
	header.Set("Cache-Control", "max-age=31536000")
	header.Set("Expires", "31536000")
	header.Set("Content-Type", "text/html; charset=utf-8")
	// security headers
	header.Set("X-Content-Type-Options", "nosniff")
	header.Set("X-Frame-Options", "DENY")
	header.Set("X-XSS-Protection", "1; mode=block")
}
