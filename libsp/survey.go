package libsp

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/beevik/etree"
)

// Survey represents a survey, including its questions, potential responses, and meta-data
type Survey struct {
	Title         string
	Description   string
	Status        string
	CreatedOn     time.Time
	LaunchedOn    time.Time
	ModifiedOn    time.Time
	QuestionOrder []string
	Questions     map[string]*Question
	Responses     []*Response
}

const timeFormat = "2006-01-02 15:04:05"

// WriteCSV saves the parsed survey questions and responses in comma-separated value format
func (s *Survey) WriteCSV(bw *bufio.Writer) error {
	if bw == nil {
		return errors.New("bw cannot be nil")
	}

	w := csv.NewWriter(bw)
	w.Write(s.csvCols())
	for _, r := range s.Responses {
		row := []string{r.ID, fmt.Sprintf("%t", r.Finished), fmt.Sprintf("%d", r.Progress), fmt.Sprintf("%d", r.Duration)}

		for _, id := range s.QuestionOrder {
			q := s.Questions[id]
			row = append(row, q.ResponseCols(r)...)
		}
		w.Write(row)
	}
	w.Flush()

	if err := w.Error(); err != nil {
		return fmt.Errorf("error writing csv: %s", err)
	}

	return nil
}

func (s *Survey) csvCols() []string {
	cols := []string{"id", "finished", "progress", "duration"}
	for _, id := range s.QuestionOrder {
		q := s.Questions[id]
		cols = append(cols, q.CSVCols()...)
	}
	return cols
}

// WriteR saves an R script suitable for importing the survey questions to R
func (s *Survey) WriteR(w *bufio.Writer) error {
	log.Print("Sorry, Survey.WriteR() isn't implemented yet :(")
	return nil
}

// ReadXML reads a Qualtrics XML file of participant responses
func (s *Survey) ReadXML(r *bufio.Reader) error {
	doc := etree.NewDocument()
	_, err := doc.ReadFrom(r)
	if err != nil {
		return fmt.Errorf("could not parse xml: %s", err)
	}
	responses := []*Response{}
	root := doc.SelectElement("Responses")
	for _, resp := range root.SelectElements("Response") {
		r := NewResponse()
		r.ID = getStringElement("_recordId", resp)
		r.Progress = getIntElement("progress", resp)
		r.Duration = getIntElement("duration", resp)
		r.Finished = getBoolElement("finished", resp)

		for _, e := range resp.ChildElements() {
			if strings.HasPrefix(e.Tag, "QID") {
				r.AddAnswer(e.Tag, e.Text())
			}
		}

		responses = append(responses, r)
	}
	s.Responses = responses
	return nil
}

func getStringElement(name string, e *etree.Element) string {
	var retval string
	if v := e.SelectElement(name); v != nil {
		retval = v.Text()
	}
	return retval
}

func getIntElement(name string, e *etree.Element) int {
	var retval int
	if v := e.SelectElement(name); v != nil {
		var err error
		retval, err = strconv.Atoi(v.Text())
		if err != nil {
			log.Printf("error converting '%s' to int: %s", v.Text(), err)
		}
	}
	return retval
}

func getBoolElement(name string, e *etree.Element) bool {
	var retval bool
	if v := e.SelectElement(name); v != nil {
		var err error
		retval, err = strconv.ParseBool(v.Text())
		if err != nil {
			log.Printf("error converting '%s' to bool: %s", v.Text(), err)
		}
	}
	return retval
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
			q, err := newQuestion(e.Payload)
			if err != nil {
				return fmt.Errorf("could not create question from JSON: %s", err)
			}
			s.QuestionOrder = append(s.QuestionOrder, q.ID) // TODO get question order from FLOW elements in QSF
			s.Questions[q.ID] = q
		}
		// TODO parse survery blocks from BL element
		// TODO parse Trash block in BL and remove those questions from our output
		// TODO parse survey question count to verify we didn't miss any questions
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
	if ok, err := regexp.Match(`"Element"\s*:\s*"SQ"`, b); err != nil {
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
	SubSelector   string
	QuestionID    string
	Choices       map[int]qsfChoice
	ChoiceOrder   []json.Number
	Answers       map[int]qsfChoice
	AnswerOrder   []json.Number
	Groups        []string
}

func (p *qsfPayload) OrderedChoices() ([]Choice, error) {
	ordered := []Choice{}
	for _, s := range p.ChoiceOrder {
		i64, err := s.Int64()
		if err != nil {
			log.Fatalf("could not convert '%s' to int: %s", s, err)
		}
		i := int(i64)

		hasText := false
		if len(p.Choices[i].TextEntry) > 0 {
			hasText, err = strconv.ParseBool(p.Choices[i].TextEntry)
			if err != nil {
				log.Fatalf("could not convert '%s' to bool: %s", p.Choices[i].TextEntry, err)
			}
		}

		c := Choice{ID: s.String(), Label: p.Choices[i].Display, HasText: hasText}
		ordered = append(ordered, c)
	}

	return ordered, nil
}

func (p *qsfPayload) OrderedAnswers() ([]Choice, error) {
	ordered := []Choice{}
	for _, s := range p.AnswerOrder {
		i64, err := s.Int64()
		if err != nil {
			log.Fatalf("could not convert '%s' to int: %s", s, err)
		}
		i := int(i64)

		hasText := false
		if len(p.Answers[i].TextEntry) > 0 {
			hasText, err = strconv.ParseBool(p.Answers[i].TextEntry)
			if err != nil {
				log.Fatalf("could not convert '%s' to bool: %s", p.Answers[i].TextEntry, err)
			}
		}

		c := Choice{ID: s.String(), Label: p.Answers[i].Display, HasText: hasText}
		ordered = append(ordered, c)
	}

	return ordered, nil
}

type qsfChoice struct {
	Display   string
	TextEntry string
}
