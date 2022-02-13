// Package routes contains everything related to endpoints
package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yael-castro/survey-renderer-api/internal/handler"
)

// SetAll use to make and set all endpoint using all handlers contained in handler.Handler structure
func SetAll(e *echo.Echo, h handler.Handler) {
	e.GET("survey/v1/:id", h.RenderSurvey)
}
