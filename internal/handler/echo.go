package handler

import "github.com/labstack/echo/v4"

// NewEcho build and initialize an instance of *echo.Echo using the Handler passed as parameter
func NewEcho(h Handler) (e *echo.Echo) {
	e = echo.New()

	e.GET("survey/v1/:id", h.ProvideTemplate)

	return
}
