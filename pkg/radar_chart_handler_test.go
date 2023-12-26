package pkg

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetRadarChart(t *testing.T) {
	e := echo.New()

	e.GET("/radar", func(c echo.Context) error {
		img, err := NewRadarChartHandler().Get(c)
		if err != nil {
			return err
		}
		return c.Blob(http.StatusOK, "image/png", img)
	})

	// Start a test HTTP server
	server := httptest.NewServer(e)
	defer server.Close()

	type TestCase struct {
		QueryParams    string
		ExpectedStatus int
	}
	testCases := []TestCase{
		{
			QueryParams:    `{"names": ["Mon","Tue","Wed"], "labels": ["Work", "Relax"], "values": [[1,2,4]]}`,
			ExpectedStatus: http.StatusOK,
		},
		// invalid json
		{
			QueryParams:    `data={"names": [Mon,"Tue"], "labels": ["Work", "Relax"], "values": [[1,2,4]]}`,
			ExpectedStatus: http.StatusUnprocessableEntity,
		},
		// bad count
		{
			QueryParams:    `data={"names": ["Mon","Tue"], "labels": ["Work", "Relax"], "values": [[1,2,4]]}`,
			ExpectedStatus: http.StatusUnprocessableEntity,
		},
		// bad count
		{
			QueryParams:    `data={"names": ["Mon","Tue", "Wed"], "labels": ["Work", "Relax"], "values": [[1,2]]}`,
			ExpectedStatus: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		url := server.URL + "/radar?data=" + url.QueryEscape(tc.QueryParams)
		resp, err := http.Get(url)
		assert.NoError(t, err)
		assert.Equal(t, tc.ExpectedStatus, resp.StatusCode)
	}
}
