package main

import (
	"log"
	"os"
	"surver-renderer-api/internal/handler"
	"surver-renderer-api/internal/routes"

	"github.com/labstack/echo/v4"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	e := echo.New()

	routes.SetAll(e, handler.New())

	log.Fatal(e.Start(":" + port))
}
