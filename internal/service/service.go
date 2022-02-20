// Package service contains the business logic
package service

import (
	"bytes"
	"github.com/yael-castro/survey-renderer-api/internal/model"
	"github.com/yael-castro/survey-renderer-api/internal/repository"
	"html/template"
	"io"
)

// TODO comment
type SurveyProvider interface {
	ProvideSurvey(string) (io.Reader, error)
}

// TODO comment
type SurveyTemplateProvider struct {
	repository.SurveyFinder
	SurveyRenderer
}

// TODO comment
func (SurveyTemplateProvider) ProvideSurvey(string) (io.Reader, error) {
	panic(`implement me`)
}

// SurveyRenderer use to build a survey template
type SurveyRenderer interface {
	// RenderSurvey use to render survey and save it in io.Reader
	RenderSurvey(model.Survey) (io.Reader, error)
}

// SurveyTemplateRenderer using a *template.Template to constructs an html page based on model.Survey
type SurveyTemplateRenderer struct {
	*template.Template
}

// RenderSurvey takes the data from model.Survey to render a html page in a io.Reader
func (sr SurveyTemplateRenderer) RenderSurvey(survey model.Survey) (io.Reader, error) {
	buffer := bytes.NewBuffer([]byte{})

	err := sr.Execute(buffer, survey)

	return buffer, err
}
