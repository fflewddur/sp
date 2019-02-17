package libsp

import (
	"encoding/json"
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
	s.CreatedOn = qs.SurveyEntry.SurveyCreationDate
	s.LaunchedOn = qs.SurveyEntry.SurveyStartDate
	s.ModifiedOn = qs.SurveyEntry.LastModified

	return nil
}

type qsf struct {
	SurveyEntry *qsfSurveyEntry
}

type qsfSurveyEntry struct {
	SurveyName         string
	SurveyDescription  string
	SurveyStatus       string
	SurveyStartDate    time.Time
	SurveyCreationDate time.Time
	LastModified       time.Time
}
