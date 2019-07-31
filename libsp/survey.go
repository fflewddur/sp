package libsp

import (
	"encoding/json"
	"regexp"
	"time"
)

// Survey represents a survey, includings its questions, potential responses, and meta-data
type Survey struct {
	Title       string
	Description string
	Status      string
	CreatedOn   time.Time
	LaunchedOn  time.Time
	ModifiedOn  time.Time
	Questions   []*Question
}

// UnmarshalJSON fills the fields of s with the data found in b
func (s *Survey) UnmarshalJSON(b []byte) error {
	var qs qsf
	if err := json.Unmarshal(b, &qs); err != nil {
		return err
	}

	s.Title = qs.SurveyEntry.SurveyName
	s.Description = qs.SurveyEntry.SurveyDescription
	s.Status = qs.SurveyEntry.SurveyStatus
	if t, err := time.Parse(time.RFC3339, qs.SurveyEntry.SurveyCreationDate); err == nil {
		s.CreatedOn = t
	}
	if t, err := time.Parse(time.RFC3339, qs.SurveyEntry.SurveyStartDate); err == nil {
		s.LaunchedOn = t
	}
	if t, err := time.Parse(time.RFC3339, qs.SurveyEntry.LastModified); err == nil {
		s.ModifiedOn = t
	}

	for _, e := range qs.SurveyElements {
		if e.Element == "SQ" {
			q := new(Question)
			q.ID = e.Payload.DataExportTag
			q.Wording = e.Payload.QuestionText
			q.Type = e.Payload.QuestionType

			s.Questions = append(s.Questions, q)
		}
	}

	return nil
}

type qsf struct {
	SurveyEntry    *qsfSurveyEntry
	SurveyElements []*qsfSurveyElement
}

type qsfSurveyEntry struct {
	SurveyID           string
	SurveyName         string
	SurveyDescription  string
	SurveyStatus       string
	SurveyStartDate    string
	SurveyCreationDate string
	LastModified       string
}

type qsfSurveyElement struct {
	SurveyID           string
	Element            string
	PrimaryAttribute   string
	SecondaryAttribute string
	Payload            *qsfPayload
}

func (e *qsfSurveyElement) UnmarshalJSON(b []byte) error {
	// Survey Questions have a Payload object, other elements have an array of Payload objects
	if ok, err := regexp.Match("\"Element\":[\\s*]\"SQ\"", b); err != nil {
		return err
	} else if ok {
		var q qsfSurveyElementQuestion
		json.Unmarshal(b, &q)
		e.Element = q.Element
		e.PrimaryAttribute = q.PrimaryAttribute
		e.SecondaryAttribute = q.SecondaryAttribute
		e.Payload = q.Payload
	}

	return nil
}

type qsfSurveyElementQuestion qsfSurveyElement

type qsfPayload struct {
	QuestionText  string
	DataExportTag string
	QuestionType  string
	Selector      string
}
