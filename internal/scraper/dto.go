package scraper

type ScraperRequest struct {
	Urls []string `json:"urls"`
}

type ScraperResponse struct {
	Url          string  `json:"url"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	StatusCode   int     `json:"status_code"`
	ResponseTime float64 `json:"response_time"`
	Error        string  `json:"error"`
}
