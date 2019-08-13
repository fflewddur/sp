package libsp

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestWriteCSV(t *testing.T) {
	r := bufio.NewReader(strings.NewReader(qsfTestContent))
	s, err := ReadQsf(r)
	if s == nil {
		t.Error("s = nil; want s != nil")
	}
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	err = s.WriteCSV(w)
	if err != nil {
		t.Errorf("err = %s; want err = nil", err)
	}
	line, err := b.ReadString('\n')
	line = strings.TrimSpace(line)
	if err != nil && err != io.EOF {
		t.Errorf("err = %s; want err = nil", err)
	}
	if line != "id,finished,progress,duration" {
		t.Errorf("line = '%s'; wanted 'id,finished,progress,duration'", line)
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
