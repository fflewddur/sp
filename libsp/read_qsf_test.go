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
	if s.Description != "" {
		t.Errorf("Description = '%s'; want empty string", s.Description)
	}
	if s.CreatedOn.String() != "2019-02-10 21:50:52 +0000 UTC" {
		t.Errorf("CreatedOn = '%s'; want '2019-02-10 21:50:52 +0000 UTC'", s.CreatedOn)
	}
	if s.ModifiedOn.String() != "2019-12-17 14:37:30 +0000 UTC" {
		t.Errorf("ModifiedOn = '%s'; want '2019-12-17 14:37:30 +0000 UTC'", s.ModifiedOn)
	}
	if s.LaunchedOn.String() != "2019-08-01 02:23:00 +0000 UTC" {
		t.Errorf("LaunchedOn = '%s'; want '2019-08-01 02:23:00 +0000 UTC'", s.LaunchedOn)
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
		{"s", true}, // embedded data
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
		{"QID18", true},
		{"QID19", true},
		{"QID20", true},
		{"QID21", true},
		{"QID22", true},
		{"QID23", true},
		{"QID26", true},
		{"QID27", true},
		{"QID28", true},
		{"QID29", true},
	}
	if len(s.Questions) != 26 {
		t.Errorf("len(Questions) = %d; want 26", len(s.Questions))
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
		return
	}
	if len(s.Questions) != 21 {
		t.Errorf("len(Questions) = %d; want 21", len(s.Questions))
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
		{"QID16", Timing},
		{"QID17", PickGroupRank},
		{"QID18", MultipleChoiceSingleResponse},
		{"QID19", MultipleChoiceMultiResponse},
		{"QID20", MultipleChoiceSingleResponse},
		{"QID21", MultipleChoiceSingleResponse},
		{"QID22", Meta},
		{"QID23", ConstantSum},
		{"QID29", RankOrder},
	}
	for _, test := range tests {
		if s.Questions[test.id] == nil {
			t.Errorf("Questions[%s] = nil", test.id)
		}
		if s.Questions[test.id].Type() != test.want {
			t.Errorf("Questions[%s].Type() = %v; wanted %v", test.id, s.Questions[test.id].Type(), test.want)
		}
	}
}

func TestForNilQuestions(t *testing.T) {
	r := bufio.NewReader(strings.NewReader(qsfTestContent))
	s, err := ReadQsf(r)
	if err != nil {
		t.Errorf("err = %s", err)
	}
	if s == nil {
		t.Error("survey = nil")
		return
	}

	for _, i := range s.QuestionOrder {
		if s.Questions[i] == nil {
			t.Errorf("Questions[%s] = nil", i)
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

func TestReadQsfChoiceExportTags(t *testing.T) {
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
		{"QID5", []tChoice{{0, "statement1"}, {1, "statement2"}, {2, "statement3"}, {3, "other"}}},
		{"QID19", []tChoice{{0, "choice1"}, {1, "choice3"}, {2, "choice2"}, {3, "none"}}},
	}
	for _, test := range tests {
		q, ok := s.Questions[test.id]
		if !ok {
			t.Errorf("%s not found in Questions", test.id)
		}

		var rc []Choice
		if q.qType.choicesAreQuestions() {
			rc = q.SubQuestions()
		} else {
			rc = q.ResponseChoices()
		}
		for i, c := range rc {
			if test.choices[i].label != c.VarName {
				t.Errorf("Questions[%s].Choices[%d] = '%s'; wanted '%s'", test.id, i, c.VarName, test.choices[i].label)
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

func TestQuestionOrder(t *testing.T) {
	r := bufio.NewReader(strings.NewReader(qsfTestContent))
	s, err := ReadQsf(r)
	if err != nil {
		t.Errorf("err = %s", err)
	}
	if s == nil {
		t.Error("survey = nil")
		return
	}

	want := []string{
		"QID14",
		"QID22",
		"QID1",
		"QID18",
		"QID3",
		"QID4",
		"QID19",
		"QID11",
		"QID5",
		"QID13",
		"QID6",
		"QID16",
		"QID7",
		"QID8",
		"QID9",
		"QID10",
		"QID17",
		"QID15",
		"QID29",
		"QID20",
		"QID21",
		"QID28",
		"QID23",
		"QID26",
		"QID27",
		"s",
	}

	if len(want) != len(s.QuestionOrder) {
		t.Errorf("len(s.QuestionOrder) = %d; wanted %d", len(s.QuestionOrder), len(want))
	}

	for i := range s.QuestionOrder {
		if want[i] != s.QuestionOrder[i] {
			t.Errorf("s.QuestionOrder[%d] = %s; wanted %s", i, s.QuestionOrder[i], want[i])
		}
	}
}
