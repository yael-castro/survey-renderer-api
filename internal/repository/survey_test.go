package repository

import (
	"errors"
	"reflect"
	"strconv"
	"testing"

	"github.com/yael-castro/survey-renderer-api/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	// ttFindSurvey tatble testing used to prove the interface SurveyFinder
	ttFindSurvey = []struct {
		SurveyFinder
		surveyId       string
		expectedSurvey model.Survey
		expectedError  error
	}{
		// This case tests the search for existent surveyId
		{
			surveyId: "1000",
			SurveyFinder: SurveyStorageNoSQL{
				Collection: NewMD(localSettings).Collection("sandbox"),
			},
			expectedSurvey: model.Survey{
				Title: "ITSOEH - Status vacunación",
			},
		},
		// This case tests the search for non-existent surveyId
		{
			SurveyFinder: SurveyStorageNoSQL{
				Collection: NewMD(localSettings).Collection("no-collection"),
			},
			surveyId:      "",
			expectedError: mongo.ErrNoDocuments,
		},
	}
)

// TestSurveyFinder_FindSurvey test the method FindSurvey of SurveyFinder
func TestSurveyFinder_FindSurvey(t *testing.T) {
	for i, v := range ttFindSurvey {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			survey, err := v.FindSurvey(v.surveyId)
			if !errors.Is(err, v.expectedError) {
				t.Fatalf(`expected error "%v" got "%v"`, v.expectedError, err)
			}

			if err != nil {
				t.Skip(err)
			}

			if !reflect.DeepEqual(survey, v.expectedSurvey) {
				t.Fatalf(`expected survey "%+v" got "%+v"`, v.expectedSurvey, survey)
			}

			t.Logf(`%+v`, survey)
		})
	}
}
