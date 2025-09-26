package scraper

import (
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

type ScraperHandler struct{}

func NewScraperHandler() ScraperHandler {
	return ScraperHandler{}
}

func (h *ScraperHandler) ScraperRoutes(c *echo.Group) {
	scraperAPI := c.Group("/scraper")

	scraperAPI.POST("", h.ScraperConcurrent)
}

func (h *ScraperHandler) ScraperConcurrent(c echo.Context) error {
	start := time.Now()
	ctx := c.Request().Context()

	var body ScraperRequest
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid body request",
		})
	}

	if len(body.Urls) == 0 {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": "urls field must be grater than 0",
		})
	}

	ch := make(chan ScraperResponse, len(body.Urls))

	var wg sync.WaitGroup
	for _, url := range body.Urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			result := WebScraperData(ctx, url)
			ch <- *result
		}(url)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	response := make([]ScraperResponse, 0, len(body.Urls))
	for r := range ch {
		response = append(response, r)
	}

	elapsed := time.Since(start)
	return c.JSON(http.StatusOK, map[string]any{
		"handler_time": elapsed.Seconds(),
		"results":      response,
	})
}
