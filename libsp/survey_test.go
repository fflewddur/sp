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

	// TODO test QID16's timing data
	var tests = [][]string{
		[]string{"id", "finished", "progress", "duration",
			"QID1", "QID3",
			"QID4_1", "QID4_2", "QID4_3", "QID4_5", "QID4_5_TEXT", "QID4_6", "QID4_6_TEXT", "QID4_4",
			"QID11", "QID11_3_TEXT",
			"QID5_1", "QID5_2", "QID5_3", "QID5_4", "QID5_4_TEXT",
			"QID13_1_1", "QID13_1_2", "QID13_1_3", "QID13_2_1", "QID13_2_2", "QID13_2_3", "QID13_3_1", "QID13_3_2", "QID13_3_3", "QID13_5_1", "QID13_5_2", "QID13_5_3", "QID13_4_1", "QID13_4_2", "QID13_4_3", "QID13_4_TEXT",
			"QID6_1", "QID6_2", "QID6_3", "QID7_TEXT", "QID8_TEXT",
			"QID9_1", "QID9_2", "QID9_3",
			"QID10_1", "QID10_2", "QID10_3",
			"QID17_1_GROUP", "QID17_1_RANK", "QID17_2_GROUP", "QID17_2_RANK", "QID17_3_GROUP", "QID17_3_RANK", "QID17_4_GROUP", "QID17_4_RANK", "QID17_5_GROUP", "QID17_5_RANK", "QID17_5_TEXT",
			"QID15", "QID15_NPS_GROUP"},
		[]string{"R_1dtWhiBDD96nfyk", "true", "100", "122",
			"Click to write Choice 1", "Click to write Choice 2",
			"", "", "Click to write Choice 3", "Other1", "other response 1", "Other2", "other response 2", "",
			"Click to write Choice 2", "",
			"Click to write Scale point 1", "Click to write Scale point 2", "Click to write Scale point 3", "Click to write Scale point 3", "other matrix row",
			"Col 1", "Col 2", "", "", "Col 2", "Col 3", "", "Col 2", "", "", "Col 2", "", "", "", "Col 3", "other matrix multiple row",
			"Click to write Scale point 1", "", "Click to write Scale point 2", "one line of text", "multiple\nlines\nof\ntext?",
			"field 1", "field 2", "field 3",
			"3", "2", "1",
			"Click to write Group 1", "3", "Click to write Group 2", "1", "Click to write Group 1", "2", "Click to write Group 1", "1", "Click to write Group 3", "1", "in group 3",
			"6", "Detractor"},
		[]string{"R_z72KJQMnr3lxZGp", "true", "100", "104",
			"Click to write Choice 3", "Click to write Choice 2",
			"", "", "", "", "", "", "", "None",
			"Click to write Choice 3", "other text",
			"Click to write Scale point 3", "Click to write Scale point 2", "", "", "",
			"Col 1", "", "", "", "", "", "Col 1", "Col 2", "Col 3", "", "", "", "", "", "", "",
			"", "Click to write Scale point 2", "Click to write Scale point 1", "foo", "bar",
			"name", "email", "job role",
			"1", "2", "3",
			"Click to write Group 1", "2", "Click to write Group 1", "1", "Click to write Group 2", "2", "Click to write Group 2", "1", "", "", "",
			"10", "Promoter"},
		[]string{"R_3MPTb9vwnCBmijR", "false", "33", "22",
			"Click to write Choice 2", "Click to write Choice 2",
			"", "Click to write Choice 2", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
			"", "", "", "", "", "", "", "", "", "", "",
			"", "", "", ""},
	}
	row := 0
	var record []string
	for _, test := range tests {
		record, err = csvReader.Read()
		if err != nil {
			t.Errorf("row %d, err = %s", row, err)
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
