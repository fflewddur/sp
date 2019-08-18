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
	r = bufio.NewReader(strings.NewReader(xmlTestContent))
	if err = s.ReadXML(r); err != nil {
		t.Errorf("could not parse xml: %s", err)
	}
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	err = s.WriteCSV(w)
	if err != nil {
		t.Errorf("err = %s", err)
	}

	r = bufio.NewReader(&b)
	csvReader := csv.NewReader(r)

	var tests = []struct {
		row  int
		col  int
		want string
	}{
		{0, 0, "id"}, {0, 1, "finished"}, {0, 2, "progress"}, {0, 3, "duration"},
		{0, 4, "QID2_1"}, {0, 5, "QID2_2"}, {0, 6, "QID2_3"},
		{0, 7, "QID9_3"}, {0, 8, "QID9_2"}, {0, 9, "QID9_1"},
		{0, 10, "QID4_1"}, {0, 11, "QID4_2"}, {0, 12, "QID4_3"},
		{0, 13, "QID3"}, {0, 14, "QID1"},
		{1, 0, "R_eg2X4t8Notm1zeV"}, {1, 1, "true"}, {1, 2, "100"}, {1, 3, "62"},
		{1, 4, ""}, {1, 5, ""}, {1, 6, ""},
		{1, 7, "line three"}, {1, 8, "line two"}, {1, 9, "line one"},
		{1, 10, "Click to write Choice 1"}, {1, 11, "Click to write Choice 2"}, {1, 12, ""},
		{1, 13, "Click to write Choice 3"}, {1, 14, "Click to write Choice 1"},
		{2, 0, "R_6EBZzqhSZOghMWt"}, {2, 1, "false"}, {2, 2, "80"}, {2, 3, "69"},
		{2, 4, ""}, {2, 5, ""}, {2, 6, ""},
		{2, 7, "c"}, {2, 8, "b"}, {2, 9, "a"},
		{2, 10, ""}, {2, 11, "Click to write Choice 2"}, {2, 12, "Click to write Choice 3"},
		{2, 13, "Click to write Choice 1"}, {2, 14, "Click to write Choice 3"},
	}
	row := -1
	var record []string
	for _, test := range tests {
		if row != test.row {
			record, err = csvReader.Read()
			if err != nil {
				t.Errorf("err = %s", err)
			}
			row = test.row
		}
		if record[test.col] != test.want {
			t.Errorf("row %d, record[%d] = '%s'; want '%s'", row, test.col, record[test.col], test.want)
		}
	}

	_, err = csvReader.Read()
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

	var tests = []struct {
		index    int
		id       string
		progress int
		duration int
		finished bool
	}{
		{0, "R_eg2X4t8Notm1zeV", 100, 62, true},
		{1, "R_6EBZzqhSZOghMWt", 80, 69, false},
	}
	for _, test := range tests {
		r := s.Responses[test.index]
		if r.ID != test.id {
			t.Errorf("Responses[%d].ID = '%s'; wanted '%s'", test.index, r.ID, test.id)
		}
		if r.Progress != test.progress {
			t.Errorf("Responses[%d].Progress = %d; wanted %d", test.index, r.Progress, test.progress)
		}
		if r.Duration != test.duration {
			t.Errorf("Responses[%d].Duration = %d; wanted %d", test.index, r.Duration, test.duration)
		}
		if r.Finished != test.finished {
			t.Errorf("Responses[%d].Finished = %t", test.index, r.Finished)
		}
	}
}
