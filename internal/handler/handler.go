// Package handler contains everything related to handle http requests
package handler

import (
	"github.com/yael-castro/survey-renderer-api/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler contains all handlers (main handler struct)
type Handler struct {
	TemplateProvider
}

// New returns a empty Handler
func New() *Handler {
	return &Handler{}
}

// TemplateProvider defines a provider of html templates
type TemplateProvider interface {
	// ProvideTemplate handles http requests made to obtain a html template (page)
	ProvideTemplate(echo.Context) error
}

// _ implement constraint for SurveyTemplateProvider
var _ TemplateProvider = SurveyTemplateProvider{}

// SurveyTemplateProvider build html pages for surveys based on a survey template
type SurveyTemplateProvider struct {
	service.SurveyProvider
}

// ProvideTemplate build a html page based on the model.Survey found by id
func (sr SurveyTemplateProvider) ProvideTemplate(c echo.Context) error {
	id := c.Param("id")
	// TODO search template
	return c.HTML(http.StatusOK, id)
}
