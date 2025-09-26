package main

import (
	"fmt"
	"log"
	"os"

	"github.com/EduardoMark/web-scraper-concurrent/internal/scraper"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load envs")
	}

	e := echo.New()
	apiV1 := e.Group("/api/v1")

	scraperHandler := scraper.NewScraperHandler()
	scraperHandler.ScraperRoutes(apiV1)

	serverPort := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	e.Logger.Fatal(e.Start(serverPort))
}
