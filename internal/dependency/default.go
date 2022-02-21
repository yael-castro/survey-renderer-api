package dependency

import (
	"fmt"
	"github.com/yael-castro/survey-renderer-api/internal/business"
	"github.com/yael-castro/survey-renderer-api/internal/handler"
	"github.com/yael-castro/survey-renderer-api/internal/repository"
	"html/template"
	"io/fs"
	"os"
	"strconv"
)

// FileSystem file system used to get embedded files
var FileSystem fs.FS

// defaultProfile is an InjectorFunc for *handler.Handler that uses a Default Profile
func defaultProfile(i interface{}) (err error) {
	h, ok := i.(*handler.Handler)
	if !ok {
		return fmt.Errorf(`required a "%T" not a "%T"`, h, i)
	}

	// Initializing MongoDB connections
	mongoPort, err := strconv.Atoi(os.Getenv("MONGO_HOST"))
	if err != nil {
		return
	}

	mongoConfig := repository.Configuration{
		Type:     repository.NoSQL,
		Host:     os.Getenv("MONGO_HOST"),
		Port:     mongoPort,
		Database: os.Getenv("MONGO_HOST"),
		User:     os.Getenv("MONGO_USER"),
		Password: os.Getenv("MONGO_PASS"),
		Secure:   os.Getenv("MONGO_SRV") == "true",
	}

	mongoDatabase, err := repository.NewMongoDatabase(mongoConfig)
	if err != nil {
		return
	}

	mongoCollection := mongoDatabase.Collection(os.Getenv("SURVEY_COLLECTION"))

	// Initializing html template for render surveys
	surveyTemplate, err := template.ParseFS(FileSystem, "index.gohtml")
	if err != nil {
		return
	}

	// Initializing handler.TemplateProvider
	h.TemplateProvider = handler.SurveyTemplateProvider{
		SurveyProvider: business.SurveyTemplateProvider{
			SurveyFinder: repository.SurveyStorageNoSQL{
				Collection: mongoCollection,
			},
			SurveyRenderer: business.SurveyTemplateRenderer{
				Template: surveyTemplate,
			},
		},
	}

	return
}
