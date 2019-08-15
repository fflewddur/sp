package libsp

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"io"
	"strings"
	"testing"
)

func TestWriteCSV(t *testing.T) {
	r := bufio.NewReader(strings.NewReader(qsfTestContent))
	s, err := ReadQsf(r)
	if s == nil {
		t.Error("s = nil")
	}
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	err = s.WriteCSV(w)
	if err != nil {
		t.Errorf("err = %s", err)
	}

	r = bufio.NewReader(&b)
	csvReader := csv.NewReader(r)
	record, err := csvReader.Read()
	if err != nil {
		t.Errorf("err = %s", err)
	}
	if record[0] != "id" {
		t.Errorf("line 0, record[0] = '%s'; want 'id'", record[0])
	}
	if record[1] != "finished" {
		t.Errorf("line 0, record[1] = '%s'; want 'finished'", record[1])
	}
	if record[2] != "progress" {
		t.Errorf("line 0, record[2] = '%s'; want 'progress'", record[2])
	}
	if record[3] != "duration" {
		t.Errorf("line 0, record[3] = '%s'; want 'duration'", record[3])
	}
	if record[4] != "QID2_1" {
		t.Errorf("line 0, record[4] = '%s'; want 'QID2_1'", record[4])
	}
	if record[5] != "QID2_2" {
		t.Errorf("line 0, record[5] = '%s'; want 'QID2_2'", record[5])
	}
	if record[6] != "QID2_3" {
		t.Errorf("line 0, record[6] = '%s'; want 'QID2_3'", record[6])
	}
	if record[7] != "QID9_3" {
		t.Errorf("line 0, record[7] = '%s'; want 'QID9_3'", record[7])
	}
	if record[8] != "QID9_2" {
		t.Errorf("line 0, record[8] = '%s'; want 'QID9_2'", record[8])
	}
	if record[9] != "QID9_1" {
		t.Errorf("line 0, record[9] = '%s'; want 'QID9_1'", record[9])
	}
	record, err = csvReader.Read()
	if err != io.EOF {
		t.Errorf("err = %v; want EOF", err)
	}
}

func TestWriteCSVNil(t *testing.T) {
	r := bufio.NewReader(strings.NewReader(qsfTestContent))
	s, err := ReadQsf(r)
	if s == nil {
		t.Error("s = nil; want s != nil")
	}
	err = s.WriteCSV(nil)
	if err != nil && err.Error() != "bw cannot be nil" {
		t.Errorf("err = %s; want err = 'bw cannot be nil'", err)
	}
}

func TestWriteR(t *testing.T) {
	r := bufio.NewReader(strings.NewReader(qsfTestContent))
	s, err := ReadQsf(r)
	if s == nil {
		t.Error("s = nil; want s != nil")
	}
	err = s.WriteR(nil)
	if err != nil {
		t.Errorf("err = %s; want err = nil", err)
	}
}

func TestReadXML(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(qsfTestContent))
	s, err := ReadQsf(reader)
	if s == nil {
		t.Error("s = nil; want s != nil")
	}

	reader = bufio.NewReader(strings.NewReader(xmlTestContent))
	err = s.ReadXML(reader)
	if err != nil {
		t.Errorf("err = %s; want err = nil", err)
	}

	if len(s.Responses) != 2 {
		t.Errorf("len(Responses) = %d; wanted 2", len(s.Responses))
	}

	r := s.Responses[0]
	if r.ID != "R_eg2X4t8Notm1zeV" {
		t.Errorf("ID = '%s'; wanted 'R_eg2X4t8Notm1zeV'", r.ID)
	}
	if r.Progress != 100 {
		t.Errorf("Progress = %d; wanted 100", r.Progress)
	}
	if r.Duration != 62 {
		t.Errorf("Duration = %d; wanted 62", r.Duration)
	}
	if r.Finished != true {
		t.Errorf("Finished = %t; wanted true", r.Finished)
	}

	r = s.Responses[1]
	if r.ID != "R_6EBZzqhSZOghMWt" {
		t.Errorf("ID = '%s'; wanted 'R_6EBZzqhSZOghMWt'", r.ID)
	}
	if r.Progress != 80 {
		t.Errorf("Progress = %d; wanted 100", r.Progress)
	}
	if r.Duration != 69 {
		t.Errorf("Duration = %d; wanted 62", r.Duration)
	}
	if r.Finished != false {
		t.Errorf("Finished = %t; wanted true", r.Finished)
	}
}
