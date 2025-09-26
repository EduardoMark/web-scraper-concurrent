package scraper

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHandlerConcurrent(t *testing.T) {
	e := echo.New()
	ctx := context.TODO()

	body := ScraperRequest{
		Urls: []string{
			"https://google.com",
			"https://golang.org/",
			"https://youtube.com",
			"https://www.wikipedia.org/",
			"https://httpbin.org/delay/2",
			"https://httpbin.org/delay/3",
		},
	}

	data, _ := json.Marshal(body)

	req := httptest.NewRequestWithContext(ctx, http.MethodPost, "/scraper",
		strings.NewReader(string(data)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := NewScraperHandler()

	if assert.NoError(t, handler.ScraperConcurrent(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotEmpty(t, rec.Body.String())
	}
}

func TestHandlerSynchronous(t *testing.T) {
	e := echo.New()
	ctx := context.TODO()

	body := ScraperRequest{
		Urls: []string{
			"https://google.com",
			"https://golang.org/",
			"https://youtube.com",
			"https://www.wikipedia.org/",
			"https://httpbin.org/delay/2",
			"https://httpbin.org/delay/3",
		},
	}

	data, _ := json.Marshal(body)

	req := httptest.NewRequestWithContext(ctx, http.MethodPost, "/scraper/synchronous",
		strings.NewReader(string(data)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := NewScraperHandler()

	if assert.NoError(t, handler.ScraperConcurrent(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotEmpty(t, rec.Body.String())
	}
}
