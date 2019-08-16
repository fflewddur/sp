package libsp

import (
	"bufio"
	"strings"
	"testing"
)

type tChoice struct {
	order int
	label string
}

func TestReadQsfMetadata(t *testing.T) {
	r := bufio.NewReader(strings.NewReader(qsfTestContent))
	s, err := ReadQsf(r)
	if err != nil {
		t.Error("err != nil; want err = nil")
	}
	if s == nil {
		t.Error("survey = nil; want survey != nil")
	}
	if s.Title != "Test survey" {
		t.Errorf("Title = '%s'; want 'Test survey'", s.Title)
	}
	if s.Status != "Inactive" {
		t.Errorf("Status = '%s'; want 'Inactive'", s.Status)
	}
	if s.Description != "Test description" {
		t.Errorf("Description = '%s'; want 'Test description'", s.Description)
	}
	if s.CreatedOn.String() != "2019-02-10 21:50:52 +0000 UTC" {
		t.Errorf("CreatedOn = '%s'; want '2019-02-10 21:50:52 +0000 UTC'", s.CreatedOn)
	}
	if s.ModifiedOn.String() != "2019-02-10 21:55:31 +0000 UTC" {
		t.Errorf("ModifiedOn = '%s'; want '2019-02-10 21:55:31 +0000 UTC'", s.ModifiedOn)
	}
	if s.LaunchedOn.String() != "2019-08-01 01:23:45 +0000 UTC" {
		t.Errorf("LaunchedOn = '%s'; want '2019-08-01 01:23:45 +0000 UTC'", s.LaunchedOn)
	}
}

func TestReadQsfQuestions(t *testing.T) {
	r := bufio.NewReader(strings.NewReader(qsfTestContent))
	s, err := ReadQsf(r)
	if err != nil {
		t.Error("err != nil; want err = nil")
	}
	if s == nil {
		t.Error("survey = nil; want survey != nil")
	}
	tests := []struct {
		id   string
		want bool
	}{
		{"", false},
		{"QID1", true},
		{"QID2", true},
		{"QID3", true},
		{"QID4", true},
		{"QID5", true},
		{"QID6", true},
		{"QID7", true},
		{"QID8", true},
		{"QID9", true},
		{"QID10", true},
		{"QID11", true},
		{"QID12", false},
	}
	if len(s.Questions) != 11 {
		t.Errorf("len(Questions) = %d; want 11", len(s.Questions))
	}
	for _, test := range tests {
		if _, ok := s.Questions[test.id]; test.want != ok {
			t.Errorf("Questions[%s] = %t", test.id, ok)
		}
	}
}

func TestReadQsfMinified(t *testing.T) {
	r := bufio.NewReader(strings.NewReader(qsfTestContentMin))
	s, err := ReadQsf(r)
	if err != nil {
		t.Error("err != nil; want err = nil")
	}
	if s == nil {
		t.Error("survey = nil; want survey != nil")
	}
	if len(s.Questions) != 10 {
		t.Errorf("len(Questions) = %d; want 10", len(s.Questions))
	}
}

func TestReadQsfTypes(t *testing.T) {
	r := bufio.NewReader(strings.NewReader(qsfTestContent))
	s, err := ReadQsf(r)
	if err != nil {
		t.Error("err != nil; want err = nil")
	}
	if s == nil {
		t.Error("survey = nil; want survey != nil")
	}

	tests := []struct {
		id   string
		want QType
	}{
		{"QID1", MultipleChoiceSingleResponse},
		{"QID2", MultipleChoiceMultiResponse},
		{"QID3", MultipleChoiceSingleResponse},
		{"QID4", MultipleChoiceMultiResponse},
		{"QID5", MatrixSingleResponse},
		{"QID6", MaxDiff},
		{"QID7", TextEntry},
		{"QID8", TextEntry},
		{"QID9", Form},
		{"QID10", RankOrder},
		{"QID11", Description},
	}
	for _, test := range tests {
		if s.Questions[test.id].Type() != test.want {
			t.Errorf("Questions[%s].Type() = %v; wanted %v", test.id, s.Questions[test.id].Type(), test.want)
		}
	}
}

func TestReadQsfChoiceOrder(t *testing.T) {
	r := bufio.NewReader(strings.NewReader(qsfTestContent))
	s, err := ReadQsf(r)
	if err != nil {
		t.Error("err != nil; want err = nil")
	}
	if s == nil {
		t.Error("survey = nil; want survey != nil")
	}

	tests := []struct {
		id      string
		choices []tChoice
	}{
		{"QID2", []tChoice{{0, "Click to write Choice 1"}, {1, "Click to write Choice 2"}, {2, "Click to write Choice 3"}}},
		{"QID9", []tChoice{{0, "Click to write Choice 3 (ordered first)"}, {1, "Click to write Choice 2 (ordered second)"}, {2, "Click to write Choice 1 (ordered third)"}}},
	}
	for _, test := range tests {
		q, ok := s.Questions[test.id]
		if !ok {
			t.Errorf("%s not found in Questions", test.id)
		}
		rc := q.ResponseChoices()
		for i, c := range rc {
			if test.choices[i].label != c.Label {
				t.Errorf("Choices[%d] = '%s'; wanted '%s'", i, c.Label, test.choices[i].label)
			}
		}
	}
}

func TestReadQsfAnswerOrder(t *testing.T) {
	r := bufio.NewReader(strings.NewReader(qsfTestContent))
	s, err := ReadQsf(r)
	if err != nil {
		t.Error("err != nil; want err = nil")
	}
	if s == nil {
		t.Error("survey = nil; want survey != nil")
	}

	tests := []struct {
		id      string
		choices []tChoice
		answers []tChoice
	}{
		{"QID5", []tChoice{{0, "Click to write Scale point 1"}, {1, "Click to write Scale point 2"}, {2, "Click to write Scale point 3"}}, []tChoice{{0, "Click to write Statement 1"}, {1, "Click to write Statement 2"}, {2, "Click to write Statement 3"}}},
	}
	for _, test := range tests {
		q, ok := s.Questions[test.id]
		if !ok {
			t.Errorf("%s not found in Questions", test.id)
		}
		rc := q.ResponseChoices()
		for i, c := range rc {
			if test.choices[i].label != c.Label {
				t.Errorf("Questions[%s].Choices[%d] = '%s'; wanted '%s'", test.id, i, c.Label, test.choices[i].label)
			}
		}
		sq := q.SubQuestions()
		for i, c := range sq {
			if test.answers[i].label != c.Label {
				t.Errorf("Questions[%s].Answers[%d] = '%s'; wanted '%s'", test.id, i, c.Label, test.answers[i].label)
			}
		}
	}
}

func TestIncompleteQSF(t *testing.T) {
	r := bufio.NewReader(strings.NewReader(qsfContentIncomplete))
	s, err := ReadQsf(r)
	if err == nil {
		t.Error("err = nil; want err != nil")
	}
	if s != nil {
		t.Error("survey != nil; want survey = nil")
	}
	if err.Error() != "could not parse: json had no SurveyEntry object" {
		t.Errorf("err = '%s'; want 'could not parse: json had no SurveyEntry object'", err)
	}
}

func TestNilReader(t *testing.T) {
	s, err := ReadQsf(nil)
	if err == nil {
		t.Error("err = nil; want err != nil")
	}
	if err.Error() != "r cannot be nil" {
		t.Errorf("err.Error() = '%s'; want 'r cannot be nil'", err)
	}
	if s != nil {
		t.Error("got != nil; want got = nil")
	}
}
