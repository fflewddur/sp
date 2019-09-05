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
		t.Errorf("err = %s", err)
	}
	if s == nil {
		t.Error("survey = nil")
		return
	}
	if s.Title != "Test survey" {
		t.Errorf("Title = '%s'; want 'Test survey'", s.Title)
	}
	if s.Status != "Active" {
		t.Errorf("Status = '%s'; want 'Active'", s.Status)
	}
	if s.Description != "Test description" {
		t.Errorf("Description = '%s'; want 'Test description'", s.Description)
	}
	if s.CreatedOn.String() != "2019-02-10 21:50:52 +0000 UTC" {
		t.Errorf("CreatedOn = '%s'; want '2019-02-10 21:50:52 +0000 UTC'", s.CreatedOn)
	}
	if s.ModifiedOn.String() != "2019-08-20 12:23:35 +0000 UTC" {
		t.Errorf("ModifiedOn = '%s'; want '2019-08-20 12:23:35 +0000 UTC'", s.ModifiedOn)
	}
	if s.LaunchedOn.String() != "2019-08-01 01:23:45 +0000 UTC" {
		t.Errorf("LaunchedOn = '%s'; want '2019-08-01 01:23:45 +0000 UTC'", s.LaunchedOn)
	}
}

func TestReadQsfQuestions(t *testing.T) {
	r := bufio.NewReader(strings.NewReader(qsfTestContent))
	s, err := ReadQsf(r)
	if err != nil {
		t.Errorf("err = %s", err)
	}
	if s == nil {
		t.Error("survey = nil")
		return
	}
	tests := []struct {
		id   string
		want bool
	}{
		{"", false},
		{"QID0", false},
		{"QID1", true},
		{"QID2", false},
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
		{"QID13", true},
		{"QID14", true},
		{"QID15", true},
		{"QID16", true},
		{"QID17", true},
		{"QID18", false},
	}
	if len(s.Questions) != 15 {
		t.Errorf("len(Questions) = %d; want 15", len(s.Questions))
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
		t.Errorf("err = %s", err)
	}
	if s == nil {
		t.Error("survey = nil")
	}
	if len(s.Questions) != 9 {
		t.Errorf("len(Questions) = %d; want 9", len(s.Questions))
	}
}

func TestReadQsfTypes(t *testing.T) {
	r := bufio.NewReader(strings.NewReader(qsfTestContent))
	s, err := ReadQsf(r)
	if err != nil {
		t.Errorf("err = %s", err)
	}
	if s == nil {
		t.Error("survey = nil")
		return
	}

	// TODO test for timing question type
	tests := []struct {
		id   string
		want QType
	}{
		{"QID1", MultipleChoiceSingleResponse},
		{"QID3", MultipleChoiceSingleResponse},
		{"QID4", MultipleChoiceMultiResponse},
		{"QID5", MatrixSingleResponse},
		{"QID6", MaxDiff},
		{"QID7", TextEntry},
		{"QID8", TextEntry},
		{"QID9", Form},
		{"QID10", RankOrder},
		{"QID11", MultipleChoiceSingleResponse},
		{"QID13", MatrixMultiResponse},
		{"QID14", Description},
		{"QID15", NPS},
		{"QID17", PickGroupRank},
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
		t.Errorf("err = %s", err)
	}
	if s == nil {
		t.Error("survey = nil")
		return
	}

	tests := []struct {
		id      string
		choices []tChoice
	}{
		{"QID13", []tChoice{{0, "Col 1"}, {1, "Col 2"}, {2, "Col 3"}}},
		{"QID11", []tChoice{{0, "Click to write Choice 2 (ordered 1st)"}, {1, "Click to write Choice 1 (ordered 2nd)"}, {2, "Click to write Choice 3"}}},
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
	}
}

func TestReadQsfAnswerOrder(t *testing.T) {
	r := bufio.NewReader(strings.NewReader(qsfTestContent))
	s, err := ReadQsf(r)
	if err != nil {
		t.Errorf("err = %s", err)
	}
	if s == nil {
		t.Error("survey = nil")
		return
	}

	tests := []struct {
		id      string
		choices []tChoice
		answers []tChoice
	}{
		{"QID13", []tChoice{{0, "Col 1"}, {1, "Col 2"}, {2, "Col 3"}}, []tChoice{{0, "Row 1"}, {1, "Row 2"}, {2, "Row 3"}, {3, ""}, {4, "Other:"}}},
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
		t.Error("err = nil")
	}
	if s != nil {
		t.Error("survey != nil")
	}
	if err.Error() != "could not parse: json had no SurveyEntry object" {
		t.Errorf("err = '%s'; want 'could not parse: json had no SurveyEntry object'", err)
	}
}

func TestNilReader(t *testing.T) {
	s, err := ReadQsf(nil)
	if err == nil {
		t.Error("err = nil")
	}
	if err.Error() != "r cannot be nil" {
		t.Errorf("err.Error() = '%s'; want 'r cannot be nil'", err)
	}
	if s != nil {
		t.Error("survey != nil")
	}
}
