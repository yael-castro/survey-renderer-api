// Package routes contains everything related to endpoints
package routes

import (
	"surver-renderer-api/internal/handler"

	"github.com/labstack/echo/v4"
)

// SetAll use to make and set all endpoint using all handlers contained in handler.Handler structure
func SetAll(e *echo.Echo, h handler.Handler) {
	v1 := e.Group("survey/v1/")

	v1.GET(":id", h.RenderSurvey)
}
