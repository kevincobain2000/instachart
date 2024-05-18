package pkg

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetBarChart(t *testing.T) {
	e := echo.New()

	e.GET("/bar", func(c echo.Context) error {
		img, err := NewBarChartHandler().Get(c)
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
			QueryParams:    `{"x": ["2022-12-23","2022-12-28","2023-12-30"], "y": [[1,2,3]]}`,
			ExpectedStatus: http.StatusOK,
		},
		{
			QueryParams:    `{"x": ["2022-12-23","2022-12-28","2023-12-30"], "y": [1,2,3]}`,
			ExpectedStatus: http.StatusUnprocessableEntity,
		},
		{
			QueryParams:    `{"x": ["2022-12-23,"2022-12-28","2023-12-30"], "y": [1,2,3]}`,
			ExpectedStatus: http.StatusUnprocessableEntity,
		},
		{
			QueryParams:    `{"x": [["2022-12-23","2022-12-28","2023-12-30"]], "y": [1,2,3]}`,
			ExpectedStatus: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		url := server.URL + "/bar?data=" + url.QueryEscape(tc.QueryParams)
		resp, err := http.Get(url)
		assert.NoError(t, err)
		assert.Equal(t, tc.ExpectedStatus, resp.StatusCode)
		if resp.StatusCode == http.StatusOK {
			assert.Equal(t, "image/png", resp.Header.Get("Content-Type"))
		} else {
			assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
		}
	}
}
