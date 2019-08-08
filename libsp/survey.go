package libsp

import (
	"encoding/json"
	"errors"
	"log"
	"regexp"
	"strconv"
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
	Questions   map[string]*Question
}

const timeFormat = "2006-01-02 15:04:05"

// UnmarshalJSON fills the fields of s with the data found in b
func (s *Survey) UnmarshalJSON(b []byte) error {
	var qs qsf
	if err := json.Unmarshal(b, &qs); err != nil {
		return err
	}
	if qs.SurveyEntry == nil {
		return errors.New("json had no SurveyEntry object")
	}

	s.Title = qs.SurveyEntry.SurveyName
	s.Description = qs.SurveyEntry.SurveyDescription
	s.Status = qs.SurveyEntry.SurveyStatus
	if t, err := time.Parse(timeFormat, qs.SurveyEntry.SurveyCreationDate); err == nil {
		s.CreatedOn = t
	}
	if t, err := time.Parse(timeFormat, qs.SurveyEntry.SurveyStartDate); err == nil {
		s.LaunchedOn = t
	}
	if t, err := time.Parse(timeFormat, qs.SurveyEntry.LastModified); err == nil {
		s.ModifiedOn = t
	}

	s.Questions = make(map[string]*Question)
	for _, e := range qs.SurveyElements {
		if e.Element == "SQ" {
			q := new(Question)
			q.ID = e.Payload.QuestionID
			q.Wording = e.Payload.QuestionText
			q.Type = QTypeFromString(e.Payload.QuestionType)
			q.Choices = e.Payload.OrderedChoices()

			s.Questions[q.ID] = q
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
	QuestionID    string
	Choices       map[int]qsfChoice
	ChoiceOrder   []string
}

func (p *qsfPayload) OrderedChoices() []string {
	orderedChoices := []string{}
	for _, s := range p.ChoiceOrder {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("could not convert '%s' to int: %s", s, err)
		}
		orderedChoices = append(orderedChoices, p.Choices[i].Display)
	}

	return orderedChoices
}

type qsfChoice struct {
	Display string
}
