package main

import (
	"embed"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/yael-castro/survey-renderer-api/internal/handler"
	"github.com/yael-castro/survey-renderer-api/internal/routes"
	"github.com/yael-castro/survey-renderer-api/internal/service"
)

//go:embed templates
var templates embed.FS

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	err := service.ParseSurveyTemplate(templates, "templates/index.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	routes.SetAll(e, handler.New())

	log.Fatal(e.Start(":" + port))
}
