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
	if err != nil {
		t.Errorf("err = %s", err)
	}
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

	var tests = [][]string{
		[]string{"id", "finished", "progress", "duration", "QID2_1", "QID2_2", "QID2_3", "QID12", "QID9_1", "QID9_2", "QID9_3", "QID5_1", "QID5_2", "QID5_3", "QID5_4", "QID6_1", "QID6_2", "QID6_3", "QID8_TEXT", "QID4_1", "QID4_2", "QID4_3", "QID4_5", "QID4_6", "QID4_4", "QID15", "QID11", "QID10_1", "QID10_2", "QID10_3", "QID3", "QID1", "QID7_TEXT"},
		[]string{"R_1dtWhiBDD96nfyk", "true", "100", "122", "", "", "", "", "field 1", "field 2", "field 3", "Click to write Scale point 1", "Click to write Scale point 2", "Click to write Scale point 3", "Click to write Scale point 3", "Click to write Scale point 1", "", "Click to write Scale point 2", "multiple\nlines\nof\ntext?", "", "", "Click to write Choice 3", "Other1", "Other2", "", "6", "Click to write Choice 2", "3", "2", "1", "Click to write Choice 2", "Click to write Choice 1", "one line of text"},
		[]string{"R_z72KJQMnr3lxZGp", "true", "100", "104", "", "", "", "", "name", "email", "job role", "Click to write Scale point 3", "Click to write Scale point 2", "", "", "", "Click to write Scale point 2", "Click to write Scale point 1", "bar", "", "", "", "", "", "None", "10", "Click to write Choice 3", "1", "2", "3", "Click to write Choice 2", "Click to write Choice 3", "foo"},
		[]string{"R_3MPTb9vwnCBmijR", "false", "33", "22", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "Click to write Choice 2", "", "", "", "", "", "", "", "", "", "Click to write Choice 2", "Click to write Choice 2", ""},
	}
	row := 0
	var record []string
	for _, test := range tests {
		record, err = csvReader.Read()
		if err != nil {
			t.Errorf("err = %s", err)
		}

		if len(record) != len(test) {
			t.Errorf("row %d, len(record) = %d; wanted %d", row, len(record), len(test))
		}
		for i, want := range test {
			if record[i] != want {
				t.Errorf("row %d, record[%d] = '%s'; want '%s'", row, i, record[i], want)
			}
		}
		row++
	}

	_, err = csvReader.Read()
	if err != io.EOF {
		t.Errorf("err = %v; want EOF", err)
	}
}

func TestWriteCSVNil(t *testing.T) {
	r := bufio.NewReader(strings.NewReader(qsfTestContent))
	s, err := ReadQsf(r)
	if err != nil {
		t.Errorf("err = %s", err)
	}
	if s == nil {
		t.Error("s = nil")
	}
	err = s.WriteCSV(nil)
	if err != nil && err.Error() != "bw cannot be nil" {
		t.Errorf("err = %s; want err = 'bw cannot be nil'", err)
	}
}

func TestWriteR(t *testing.T) {
	r := bufio.NewReader(strings.NewReader(qsfTestContent))
	s, err := ReadQsf(r)
	if err != nil {
		t.Errorf("err = %s", err)
	}
	if s == nil {
		t.Error("s = nil")
	}
	err = s.WriteR(nil)
	if err != nil {
		t.Errorf("err = %s", err)
	}
}

func TestReadXML(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(qsfTestContent))
	s, err := ReadQsf(reader)
	if err != nil {
		t.Errorf("err = %s", err)
	}
	if s == nil {
		t.Error("s = nil")
	}

	reader = bufio.NewReader(strings.NewReader(xmlTestContent))
	err = s.ReadXML(reader)
	if err != nil {
		t.Errorf("err = %s", err)
	}

	if len(s.Responses) != 3 {
		t.Errorf("len(Responses) = %d; wanted 3", len(s.Responses))
	}

	var tests = []struct {
		index    int
		id       string
		progress int
		duration int
		finished bool
	}{
		{0, "R_1dtWhiBDD96nfyk", 100, 122, true},
		{1, "R_z72KJQMnr3lxZGp", 100, 104, true},
		{2, "R_3MPTb9vwnCBmijR", 33, 22, false},
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
