package main

import (
	"embed"
	"log"
	"os"

	"github.com/yael-castro/survey-renderer-api/internal/dependency"
	"github.com/yael-castro/survey-renderer-api/internal/handler"
)

const defaultPort = "8080"

//go:embed templates/index.gohtml
var defaultFileSystem embed.FS

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Initializing default log
	log.SetFlags(log.Flags() | log.Lshortfile)

	h := handler.New()

	// Loading file system for embed files
	dependency.FileSystem = defaultFileSystem

	err := dependency.NewInjector(dependency.Default).Inject(h)
	if err != nil {
		log.Fatal(err)
	}

	// Building an initialized instance of *echo.Echo
	e := handler.NewEcho(*h)

	log.Fatal(e.Start(":" + port))
}
