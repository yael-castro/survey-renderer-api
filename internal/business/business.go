// Package business contains the business logic
package business

import (
	"bytes"
	"github.com/yael-castro/survey-renderer-api/internal/repository"
	"html/template"
	"io"
)

// SurveyProvider defines a provider of serialized survey data
type SurveyProvider interface {
	// ProvideSurvey search for a survey by id and save it in the io.Reader
	ProvideSurvey(string) (io.Reader, error)
}

// _ "implement" constraint for SurveyTemplateProvider
var _ SurveyProvider = SurveyTemplateProvider{}

// SurveyTemplateProvider provider of serialized survey data in html
type SurveyTemplateProvider struct {
	repository.SurveyFinder
	*template.Template
}

// ProvideSurvey search a model.Survey to then render it in a html page returned in io.Reader
func (p SurveyTemplateProvider) ProvideSurvey(surveyId string) (io.Reader, error) {
	survey, err := p.FindSurvey(surveyId)
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer([]byte{})
	err = p.Execute(buffer, survey)
	return buffer, err
}

// _ "implement" constraint for SurveyProviderMock
var _ SurveyProvider = SurveyProviderMock{}

// SurveyProviderMock mock of SurveyProvider used to testing
type SurveyProviderMock struct {
	reader io.Reader
	error
}

// ProvideSurvey returns m.reader and m.error (the string parameter is ignored)
func (m SurveyProviderMock) ProvideSurvey(string) (io.Reader, error) {
	return m.reader, m.error
}
