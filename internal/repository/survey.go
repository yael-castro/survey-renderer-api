package repository

import (
	"context"

	"github.com/yael-castro/survey-renderer-api/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// SurveyFinder defines the survey finder
type SurveyFinder interface {
	// FindSurvey search survey by id
	FindSurvey(string) (model.Survey, error)
}

// _ implements constraint for SurveyStorageNoSQL
var _ SurveyFinder = SurveyStorageNoSQL{}

// SurveyStorageNoSQL implementation of SurveyFinder that search a survey in a NoSQL storage oriented to documents
type SurveyStorageNoSQL struct {
	*mongo.Collection
}

// FindSurveyContext search a survey in an MongoDB collection using the surveyId passed as paramenter and
// the search is canceled if the context is canceled
func (s SurveyStorageNoSQL) FindSurveyContext(ctx context.Context, surveyId string) (survey model.Survey, err error) {
	err = s.FindOne(ctx, bson.M{"id": surveyId}).Decode(&survey)
	return
}

// FindSurvey works the same that FindSurveyContext but uses as default context the context.TODO()
func (s SurveyStorageNoSQL) FindSurvey(surveyId string) (survey model.Survey, err error) {
	return s.FindSurveyContext(context.TODO(), surveyId)
}
