package libsp

import (
	"bufio"
	"encoding/json"
	"errors"
	"log"
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
	Questions   map[string]*Question
	Responses   []*Response
}

const timeFormat = "2006-01-02 15:04:05"

// WriteCSV saves the parsed survey questions and responses in comma-separated value format
func (s *Survey) WriteCSV(w *bufio.Writer) error {
	log.Print("Sorry, Survey.WriteCSV() isn't implemented yet :(")
	return nil
}

// WriteR saves an R script suitable for importing the survey questions to R
func (s *Survey) WriteR(w *bufio.Writer) error {
	log.Print("Sorry, Survey.WriteR() isn't implemented yet :(")
	return nil
}

// ReadXML reads a Qualtrics XML file of participant responses
func (s *Survey) ReadXML(r *bufio.Reader) error {
	return nil
}

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
			q := newQuestion(e.Payload)
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
	// Survey questions have a Payload object, other elements have an array of Payload objects
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
	ChoiceOrder   []json.Number
	Answers       map[int]qsfChoice
	AnswerOrder   []json.Number
}

func (p *qsfPayload) OrderedChoices() []string {
	ordered := []string{}
	for _, s := range p.ChoiceOrder {
		i, err := s.Int64()
		if err != nil {
			log.Fatalf("could not convert '%s' to int: %s", s, err)
		}
		ordered = append(ordered, p.Choices[int(i)].Display)
	}

	return ordered
}

func (p *qsfPayload) OrderedAnswers() []string {
	ordered := []string{}
	for _, s := range p.AnswerOrder {
		i, err := s.Int64()
		if err != nil {
			log.Fatalf("could not convert '%s' to int: %s", s, err)
		}
		ordered = append(ordered, p.Answers[int(i)].Display)
	}

	return ordered
}

type qsfChoice struct {
	Display string
}
