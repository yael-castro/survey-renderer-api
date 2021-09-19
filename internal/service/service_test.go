package service

import (
	"bytes"
	"errors"
	"io"
	"reflect"
	"strconv"
	"testing"

	"github.com/yael-castro/survey-renderer-api/internal/model"
)

var (
	errFailedRendering = errors.New("rendering failed")

	renderTests = []struct {
		// Input
		model.Survey
		// Evaluate object
		SurveyRenderer
		// Expected output
		ExpectedError error
		io.Reader
	}{
		{
			SurveyRenderer: surveyRendererMock{
				Reader: bytes.NewBuffer([]byte(`<!DOCTYPE html><html></html>`)),
			},
			Survey: model.Survey{
				Title: "Title 1",
			},
			Reader: bytes.NewBuffer([]byte(`<!DOCTYPE html><html></html>`)),
		},
		{
			SurveyRenderer: surveyRendererMock{
				Reader: bytes.NewBuffer([]byte(`<!DOCTYPE html><html></html>`)),
			},
			Survey: model.Survey{
				Title: "Title 2",
			},
			Reader: bytes.NewBuffer([]byte(`<!DOCTYPE html><html></html>`)),
		},
		{
			SurveyRenderer: surveyRendererMock{
				Reader: bytes.NewBuffer([]byte(`<!DOCTYPE hmtl><html>`)),
				error:  errFailedRendering,
			},
			ExpectedError: errFailedRendering,
		},
	}
)

// SurveyRenderer mock used to testing
type surveyRendererMock struct {
	io.Reader
	error
}

func (rm surveyRendererMock) RenderSurvey(model.Survey) (io.Reader, error) {
	return rm.Reader, rm.error
}

// TestSurveyRenderer unit test for survey rendering
func TestSurveyRenderer_RenderSurvey(t *testing.T) {
	for i, renderTest := range renderTests {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			renderer := renderTest.SurveyRenderer

			gotReader, gotErr := renderer.RenderSurvey(renderTest.Survey)
			if gotErr != renderTest.ExpectedError {
				t.Fatalf("expected error '%v' got '%v'", renderTest.ExpectedError, gotErr)
			}

			if renderTest.ExpectedError != nil {
				return
			}

			if gotReader == nil {
				t.Fatal("nil data")
			}

			gotData, err := io.ReadAll(gotReader)
			if err != nil {
				t.Fatalf("error reading got data: %v", err)
			}

			expectedData, err := io.ReadAll(renderTest)
			if err != nil {
				t.Fatalf("error reading expected data: %v", err)
			}

			if !reflect.DeepEqual(expectedData, gotData) {
				t.Fatalf("expected data '%v' got '%v'", expectedData, gotData)
			}
		})
	}
}
