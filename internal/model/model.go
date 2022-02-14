// Package model contains domain objects, errors, application data types
package model

type (
	// Survey struct used to build a survey
	Survey struct {
		// Action indicates where send data
		Action string `json:"action"`
		// Method http method used to send data
		Method string `json:"method"`
		// Title survey title
		Title string `json:"title"`
		// Description survey prompt
		Description string `json:"description"`
		// Language defines the language that use the survey
		Language string `json:"language"`
		Questions
	}

	// Questions alias for []Question
	Questions = []Question

	// Question contains information about question of survey
	Question struct {
		Text string `json:"text"`
		// Suggestions array of string used to show suggestions to answer an open question (show only if Options field is nil)
		Suggestions []string `json:"suggestions"`
		// Options array of string used to show unique options to answer the question
		Options []string `json:"options"`
	}
)
