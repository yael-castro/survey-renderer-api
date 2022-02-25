package business

import (
	"bytes"
	"errors"
	"io"
	"reflect"
	"strconv"
	"testing"
)

var (
	errFailedRendering = errors.New("rendering failed")

	ttProvideSurvey = []struct {
		// Evaluate object
		SurveyProvider
		// surveyId Input
		surveyId string
		// Expected output
		expectedError error
		expectedData  io.Reader
	}{
		{
			SurveyProvider: SurveyProviderMock{reader: bytes.NewBuffer([]byte(`<!DOCTYPE html><html></html>`))},
			expectedData:   bytes.NewBuffer([]byte(`<!DOCTYPE html><html></html>`)),
		},
		{
			SurveyProvider: SurveyProviderMock{reader: bytes.NewBuffer([]byte(`<!DOCTYPE html><html></html>`))},
			expectedData:   bytes.NewBuffer([]byte(`<!DOCTYPE html><html></html>`)),
		},
		{
			SurveyProvider: SurveyProviderMock{error: errFailedRendering},
			expectedError:  errFailedRendering,
		},
	}
)

// TestSurveyRenderer_RenderSurvey unit test for survey rendering
func TestSurveyRenderer_RenderSurvey(t *testing.T) {
	for i, v := range ttProvideSurvey {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			provider := v.SurveyProvider

			gotReader, gotErr := provider.ProvideSurvey(v.surveyId)
			if gotErr != v.expectedError {
				t.Fatalf("expected error '%v' got '%v'", v.expectedError, gotErr)
			}

			if v.expectedError != nil {
				return
			}

			if gotReader == nil {
				t.Fatal("nil data")
			}

			gotData, err := io.ReadAll(gotReader)
			if err != nil {
				t.Fatalf("error reading got data: %v", err)
			}

			expectedData, err := io.ReadAll(v.expectedData)
			if err != nil {
				t.Fatalf("error reading expected data: %v", err)
			}

			if !reflect.DeepEqual(expectedData, gotData) {
				t.Fatalf("expected data '%v' got '%v'", expectedData, gotData)
			}
		})
	}
}
