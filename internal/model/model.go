// Package model contains all data transfer objects (dto)
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
		// Prompt survey prompt (description)
		Description string `json:"description"`
		// Language language of survey
		Language string `json:"language"`
		Questions
	}

	// Questions array of questions
	Questions []Question

	// Question question of survey
	Question struct {
		Text        string   `json:"text"`
		// Suggestions array of string used to show suggestions to answer a open question (show only if Options field is nil)
		Suggestions []string `json:"suggestions"`
		// Options array of string used to show unique options to answer the question
		Options     []string `json:"options"`
	}
)