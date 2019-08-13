package libsp

import (
	"bufio"
	"strings"
	"testing"
)

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
	if len(s.Questions) != 11 {
		t.Errorf("len(Questions) = %d; want 11", len(s.Questions))
	}
	if _, ok := s.Questions["QID1"]; !ok {
		t.Error("QID1 not found in s.Questions")
	}
	if _, ok := s.Questions["QID2"]; !ok {
		t.Error("QID2 not found in s.Questions")
	}
	if _, ok := s.Questions["QID3"]; !ok {
		t.Error("QID3 not found in s.Questions")
	}
	if _, ok := s.Questions["QID4"]; !ok {
		t.Error("QID4 not found in s.Questions")
	}
	if _, ok := s.Questions["QID5"]; !ok {
		t.Error("QID5 not found in s.Questions")
	}
	if _, ok := s.Questions["QID6"]; !ok {
		t.Error("QID6 not found in s.Questions")
	}
	if _, ok := s.Questions["QID7"]; !ok {
		t.Error("QID7 not found in s.Questions")
	}
	if _, ok := s.Questions["QID8"]; !ok {
		t.Error("QID8 not found in s.Questions")
	}
	if _, ok := s.Questions["QID9"]; !ok {
		t.Error("QID9 not found in s.Questions")
	}
	if _, ok := s.Questions["QID10"]; !ok {
		t.Error("QID10 not found in s.Questions")
	}
	if _, ok := s.Questions["QID11"]; !ok {
		t.Error("QID11 not found in s.Questions")
	}
	if _, ok := s.Questions["QID12"]; ok {
		t.Error("QID11 found in s.Questions")
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

	if s.Questions["QID1"].Type() != MultipleChoiceSingleResponse {
		t.Errorf("Type = %v; wanted %v", s.Questions["QID1"].Type(), MultipleChoiceSingleResponse)
	}
	if s.Questions["QID4"].Type() != MultipleChoiceMultiResponse {
		t.Errorf("Type = %v; wanted %v", s.Questions["QID4"].Type(), MultipleChoiceMultiResponse)
	}
	if s.Questions["QID5"].Type() != MatrixSingleResponse {
		t.Errorf("Type = %v; wanted %v", s.Questions["QID5"].Type(), MatrixSingleResponse)
	}
	if s.Questions["QID6"].Type() != MaxDiff {
		t.Errorf("Type = %v; wanted %v", s.Questions["QID6"].Type(), MaxDiff)
	}
	if s.Questions["QID7"].Type() != TextEntry {
		t.Errorf("Type = %v; wanted %v", s.Questions["QID7"].Type(), TextEntry)
	}
	if s.Questions["QID10"].Type() != RankOrder {
		t.Errorf("Type = %v; wanted %v", s.Questions["QID10"].Type(), RankOrder)
	}
	if s.Questions["QID11"].Type() != Description {
		t.Errorf("Type = %v; wanted %v", s.Questions["QID11"].Type(), Description)
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

	// Ascending order
	q, ok := s.Questions["QID2"]
	if !ok {
		t.Error("QID2 not found in s.Questions")
	}
	rc := q.ResponseChoices()
	if rc[0] != "Click to write Choice 1" {
		t.Errorf("Choices[0] = '%s'; wanted 'Click to write Choice 1'", rc[0])
	}
	if rc[1] != "Click to write Choice 2" {
		t.Errorf("Choices[1] = '%s'; wanted 'Click to write Choice 2'", rc[1])
	}
	if rc[2] != "Click to write Choice 3" {
		t.Errorf("Choices[2] = '%s'; wanted 'Click to write Choice 3'", rc[2])
	}

	// Descending order
	q, ok = s.Questions["QID9"]
	if !ok {
		t.Error("QID9 not found in s.Questions")
	}
	rc = q.ResponseChoices()
	if rc[0] != "Click to write Choice 3 (ordered first)" {
		t.Errorf("Choices[0] = '%s'; wanted 'Click to write Choice 3 (ordered first)'", rc[0])
	}
	if rc[1] != "Click to write Choice 2 (ordered second)" {
		t.Errorf("Choices[1] = '%s'; wanted 'Click to write Choice 2 (ordered second)'", rc[1])
	}
	if rc[2] != "Click to write Choice 1 (ordered third)" {
		t.Errorf("Choices[2] = '%s'; wanted 'Click to write Choice 1 (ordered third)'", rc[2])
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

	q, ok := s.Questions["QID5"]
	if !ok {
		t.Error("QID5 not found in s.Questions")
	}
	rc := q.ResponseChoices()
	if rc[0] != "Click to write Scale point 1" {
		t.Errorf("Choices[0] = '%s'; wanted 'Click to write Scale point 1'", rc[0])
	}
	if rc[1] != "Click to write Scale point 2" {
		t.Errorf("Choices[1] = '%s'; wanted 'Click to write Scale point 2'", rc[1])
	}
	if rc[2] != "Click to write Scale point 3" {
		t.Errorf("Choices[2] = '%s'; wanted 'Click to write Scale point 3'", rc[2])
	}
	sq := q.SubQuestions()
	if sq[0] != "Click to write Statement 1" {
		t.Errorf("Choices[0] = '%s'; wanted 'Click to write Statement 1'", sq[0])
	}
	if sq[1] != "Click to write Statement 2" {
		t.Errorf("Choices[1] = '%s'; wanted 'Click to write Statement 2'", sq[1])
	}
	if sq[2] != "Click to write Statement 3" {
		t.Errorf("Choices[2] = '%s'; wanted 'Click to write Statement 3'", sq[2])
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
