package scraper

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/go-shiori/go-readability"
)

func WebScraperData(ctx context.Context, url string) *ScraperResponse {
	parsedUrl, isValid := isValidURL(url)
	if !isValid {
		return &ScraperResponse{
			Url:   url,
			Error: "invalid url",
		}
	}

	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		return &ScraperResponse{
			Url:   url,
			Error: "failed to scraper web page",
		}
	}
	defer resp.Body.Close()

	elapsed := time.Since(start)

	result, err := readability.FromReader(resp.Body, parsedUrl)
	if err != nil {
		return &ScraperResponse{
			Url:          url,
			StatusCode:   resp.StatusCode,
			ResponseTime: elapsed.Seconds(),
			Error:        "failed to get data from web page",
		}
	}

	return &ScraperResponse{
		Url:          url,
		Title:        result.Title,
		Description:  result.Excerpt,
		StatusCode:   resp.StatusCode,
		ResponseTime: elapsed.Seconds(),
	}
}

func isValidURL(rawURL string) (*url.URL, bool) {
	url, err := url.Parse(rawURL)
	if err != nil {
		return nil, false
	}

	return url, true
}
