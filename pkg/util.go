package pkg

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/imroc/req/v3"
)

func IsURL(urlStr string) bool {
	parsedURL, err := url.ParseRequestURI(urlStr)
	return err == nil && parsedURL.Scheme != "" && parsedURL.Host != ""
}

func IsAllowedDomain(urlStr string, allowedDomains string) bool {
	if allowedDomains == "" {
		return false // default do not allow any urls
	}

	// Parse the URL to extract the domain
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return false // If the URL is invalid, do not allow
	}
	domain := parsedURL.Hostname()

	// Split the allowedDomains into a slice
	domains := strings.Split(allowedDomains, ",")

	// Check if the domain is in the list of allowed domains
	for _, d := range domains {
		if domain == d {
			return true
		}
	}

	return false
}

func GetURL(urlStr string) (string, error) {
	resp, err := req.Get(urlStr)
	if err != nil {
		return "", err
	}
	return resp.ToString()
}

// NumberToK converts a number to a string with 'k' for thousands and 'm' for millions.
func NumberToK(num *float64) string {
	if num == nil {
		return "0"
	}

	formatNumber := func(n float64) string {
		if n == float64(int64(n)) {
			// If n is an integer, format without decimal places.
			return strconv.FormatFloat(n, 'f', 0, 64)
		}
		// Otherwise, format with one decimal place.
		return strconv.FormatFloat(n, 'f', 1, 64)
	}

	if *num < 1000 {
		return formatNumber(*num)
	}

	if *num < 1000000 {
		return formatNumber(*num/1000) + "k"
	}

	return formatNumber(*num/1000000) + "m"
}

func Truncate(s string, max int) string {
	if len(s) > max {
		return s[:max] + "..."
	}
	return s
}
