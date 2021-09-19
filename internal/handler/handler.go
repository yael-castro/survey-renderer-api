// Package handler contains everything related to handle http requests
package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler contains all handlers (main handler struct)
type Handler struct {
	SurveyRenderer
}

// New returns a initialized Handler
func New() Handler {
	return Handler{
		SurveyRenderer: NewSurveyRenderer(),
	}
}

// SurveyRenderer implement to handle request for render (draw) surveys
type SurveyRenderer interface {
	// RenderSurvey use to handle a request and return renderized survey in html format
	RenderSurvey(echo.Context) error
}

// NewSurveyRenderer returns an implementation of survey renderer
func NewSurveyRenderer() SurveyRenderer {
	return &surveyRenderer{}
}

type surveyRenderer struct{}

func (sr *surveyRenderer) RenderSurvey(c echo.Context) error {
	id := c.Param("id")

	return c.HTML(http.StatusOK, id)
}
