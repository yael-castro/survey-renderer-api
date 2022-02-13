// Package service contains the business logic
package service

import (
	"bytes"
	"io"
	"io/fs"
	"sync"
	"text/template"

	"github.com/yael-castro/survey-renderer-api/internal/model"
)

// surveyTemplate use to survey an template
var surveyTemplate *template.Template

// onceParseTemplate used to initialize the survey template once
var onceParseTemplate sync.Once

// NewSurveyTemplate returns a survey template (singleton).
//
// Initialize singleton with ParseSurveyTemplate
func NewSurveyTemplate() *template.Template {
	return surveyTemplate
}

// ParseSurveyTemplate use to parse and save the survey template to can be use NewSurveyTemplate.
//
// Important: must be called before of the function NewSurveyTemplate
func ParseSurveyTemplate(fileSystem fs.FS, name string) (err error) {
	onceParseTemplate.Do(func() {
		surveyTemplate, err = template.ParseFS(fileSystem, name)
	})
	return
}

// SurveyRenderer use to build a survey template
type SurveyRenderer interface {
	// RenderSurvey use to render survey and save it in io.Reader
	RenderSurvey(model.Survey) (io.Reader, error)
}

// NewSurveyRenderer constructs a new SurveyRenderer
func NewSurveyRenderer() SurveyRenderer {
	return &surveyRenderer{
		Template: NewSurveyTemplate(),
	}
}

type surveyRenderer struct {
	*template.Template
}

func (sr *surveyRenderer) RenderSurvey(survey model.Survey) (io.Reader, error) {
	buffer := bytes.NewBuffer([]byte{})

	err := surveyTemplate.Execute(buffer, survey)

	return buffer, err
}
