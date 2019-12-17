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
		return
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
		{"id", "finished", "progress", "duration", "recorded",
			"Q1Label", "Q18Label", "Q3Label",
			"Q4Label_1", "Q4Label_2", "Q4Label_3", "Q4Label_5", "Q4Label_5_text", "Q4Label_6", "Q4Label_6_text", "Q4Label_4",
			"Q19_choice1", "Q19_choice3", "Q19_choice2", "Q19_none",
			"Q11Label", "Q11Label_3_text",
			"Q5Label_statement1", "Q5Label_statement2", "Q5Label_statement3", "Q5Label_other", "Q5Label_other_text",
			"Q13Label_1_1", "Q13Label_1_2", "Q13Label_1_3", "Q13Label_2_1", "Q13Label_2_2", "Q13Label_2_3", "Q13Label_3_1", "Q13Label_3_2", "Q13Label_3_3", "Q13Label_5_1", "Q13Label_5_2", "Q13Label_5_3", "Q13Label_4_1", "Q13Label_4_2", "Q13Label_4_3", "Q13Label_4_text",
			"Q6Label_1", "Q6Label_2", "Q6Label_3",
			"Q16_first_click", "Q16_last_click", "Q16_page_submit", "Q16_click_count",
			"Q7Label_text", "Q8Label_text",
			"Q9Label_1", "Q9Label_2", "Q9Label_3",
			"RO_1", "RO_2", "RO_3",
			"PGR_item.1_GROUP", "PGR_item.1_RANK", "PGR_item.2_GROUP", "PGR_item.2_RANK", "PGR_item.3_GROUP", "PGR_item.3_RANK", "PGR_item.4_GROUP", "PGR_item.4_RANK", "PGR_other_GROUP", "PGR_other_RANK", "PGR_other_text",
			"NPS", "NPS_group",
			"s"},
		{"R_1dtWhiBDD96nfyk", "true", "100", "122", "2019-08-20 12:44:31",
			"Click to write Choice 1", "", "Click to write Choice 2",
			"FALSE", "FALSE", "TRUE", "TRUE", "other response 1", "TRUE", "other response 2", "FALSE",
			"", "TRUE", "", "",
			"Click to write Choice 2 (ordered 1st)", "No response",
			"scale1", "scale2", "scale3", "scale3", "other matrix row",
			"TRUE", "TRUE", "FALSE", "FALSE", "TRUE", "TRUE", "FALSE", "TRUE", "FALSE", "FALSE", "TRUE", "FALSE", "FALSE", "FALSE", "TRUE", "other matrix multiple row",
			"Click to write Scale point 1", "No response", "Click to write Scale point 2",
			"1.313", "32.89", "33.84", "15",
			"one line of text", "multiple\nlines\nof\ntext?",
			"field 1", "field 2", "field 3",
			"3", "2", "1",
			"Click to write Group 1", "3", "Click to write Group 2", "1", "Click to write Group 1", "2", "Click to write Group 1", "1", "Click to write Group 3", "1", "in group 3",
			"6", "Detractor",
			"g"},
		{"R_z72KJQMnr3lxZGp", "true", "100", "104", "2019-08-20 12:46:35",
			"Click to write Choice 3", "", "Click to write Choice 2",
			"", "", "", "", "", "", "", "TRUE",
			"", "", "", "TRUE",
			"Click to write Choice 3", "other text",
			"scale3", "scale2", "", "", "",
			"TRUE", "", "", "", "", "", "TRUE", "TRUE", "TRUE", "", "", "", "", "", "", "",
			"", "Click to write Scale point 2", "Click to write Scale point 1",
			"3.172", "25.605", "26.387", "9",
			"foo", "bar",
			"name", "email", "job role",
			"1", "2", "3",
			"Click to write Group 1", "2", "Click to write Group 1", "1", "Click to write Group 2", "2", "Click to write Group 2", "1", "", "", "",
			"10", "Promoter",
			"e"},
		{"R_3MPTb9vwnCBmijR", "false", "33", "22", "2019-08-20 12:52:35",
			"Click to write Choice 2", "", "Click to write Choice 2",
			"", "TRUE", "", "", "", "", "", "",
			"TRUE", "", "TRUE", "",
			"No response", "No response", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
			"", "", "", "",
			"", "", "", "", "", "", "", "", "", "", "",
			"", "", "", "", ""},
	}

	for row, test := range tests {
		record, err := csvReader.Read()
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
		return
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
		return
	}
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	err = s.WriteR(w, "test.csv")
	if err != nil {
		t.Errorf("err = %s", err)
	}

	r = bufio.NewReader(&b)
	var tests = []string{
		"# Generated by sp (https://github.com/fflewddur/survey_parser)",
		"library(readr)",
		"",
		"input_path <- \"test.csv\"",
		`scale_0d33bdb7dd7ad7e7644895dab595541b141f5b39 <- c("Click to write Choice 1", "Click to write Choice 2", "Click to write Choice 3", "No response")`,
		`scale_30f77603f9a26644196c8b5400b99945c2c294a5 <- c("Click to write Choice 2 (ordered 1st)", "Click to write Choice 1 (ordered 2nd)", "Click to write Choice 3", "No response")`,
		`scale_6e0893b06a434e2872d7c7f9430394c9318f3537 <- c("1", "2", "3", "4", "5", "No response")`,
		`scale_6fe6de0859f881e87e5656b0b361e2830bbbc058 <- c("Click to write Group 1", "Click to write Group 2", "Click to write Group 3", "No response")`,
		`scale_8a3feac0fed5c3348d223a4c4a52d9b74e347e0a <- c("choice2", "choice3", "choice1", "No response")`,
		`scale_9bc0385ea2c175f3341306637ae392b35bd86573 <- c("scale1", "scale2", "scale3", "scale.na", "No response")`,
		"",
		"message(sprintf(\"Reading %s...\", input_path))",
		"data <- read_csv(input_path, col_types = cols(",
		"finished = col_logical(),",
		"progress = col_integer(),",
		"duration = col_integer(),",
		"recorded = col_datetime(),",
		"Q1Label = col_factor(levels = scale_0d33bdb7dd7ad7e7644895dab595541b141f5b39),",
		"Q18Label = col_factor(levels = scale_8a3feac0fed5c3348d223a4c4a52d9b74e347e0a),",
		"Q3Label = col_factor(levels = scale_0d33bdb7dd7ad7e7644895dab595541b141f5b39),",
		"Q4Label_1 = col_logical(),",
		"Q4Label_2 = col_logical(),",
		"Q4Label_3 = col_logical(),",
		"Q4Label_5 = col_logical(),",
		"Q4Label_6 = col_logical(),",
		"Q4Label_4 = col_logical(),",
		"Q19_choice1 = col_logical(),",
		"Q19_choice3 = col_logical(),",
		"Q19_choice2 = col_logical(),",
		"Q19_none = col_logical(),",
		"Q11Label = col_factor(levels = scale_30f77603f9a26644196c8b5400b99945c2c294a5),",
		"Q5Label_statement1 = col_factor(levels = scale_9bc0385ea2c175f3341306637ae392b35bd86573, ordered = TRUE),",
		"Q5Label_statement2 = col_factor(levels = scale_9bc0385ea2c175f3341306637ae392b35bd86573, ordered = TRUE),",
		"Q5Label_statement3 = col_factor(levels = scale_9bc0385ea2c175f3341306637ae392b35bd86573, ordered = TRUE),",
		"Q5Label_other = col_factor(levels = scale_9bc0385ea2c175f3341306637ae392b35bd86573, ordered = TRUE),",
		"Q13Label_1_1 = col_logical(),",
		"Q13Label_1_2 = col_logical(),",
		"Q13Label_1_3 = col_logical(),",
		"Q13Label_2_1 = col_logical(),",
		"Q13Label_2_2 = col_logical(),",
		"Q13Label_2_3 = col_logical(),",
		"Q13Label_3_1 = col_logical(),",
		"Q13Label_3_2 = col_logical(),",
		"Q13Label_3_3 = col_logical(),",
		"Q13Label_5_1 = col_logical(),",
		"Q13Label_5_2 = col_logical(),",
		"Q13Label_5_3 = col_logical(),",
		"Q13Label_4_1 = col_logical(),",
		"Q13Label_4_2 = col_logical(),",
		"Q13Label_4_3 = col_logical(),",
		"Q6Label_1 = col_logical(),",
		"Q6Label_2 = col_logical(),",
		"Q6Label_3 = col_logical(),",
		"Q16_first_click = col_double(),",
		"Q16_last_click = col_double(),",
		"Q16_page_submit = col_double(),",
		"Q16_click_count = col_integer(),",
		"Q9Label_1 = col_logical(),",
		"Q9Label_2 = col_logical(),",
		"Q9Label_3 = col_logical(),",
		"RO_1 = col_logical(),",
		"RO_2 = col_logical(),",
		"RO_3 = col_logical(),",
		"PGR_item.1_GROUP = col_factor(levels = scale_6fe6de0859f881e87e5656b0b361e2830bbbc058),",
		"PGR_item.1_RANK = col_factor(levels = scale_6e0893b06a434e2872d7c7f9430394c9318f3537, ordered = TRUE),",
		"PGR_item.2_GROUP = col_factor(levels = scale_6fe6de0859f881e87e5656b0b361e2830bbbc058),",
		"PGR_item.2_RANK = col_factor(levels = scale_6e0893b06a434e2872d7c7f9430394c9318f3537, ordered = TRUE),",
		"PGR_item.3_GROUP = col_factor(levels = scale_6fe6de0859f881e87e5656b0b361e2830bbbc058),",
		"PGR_item.3_RANK = col_factor(levels = scale_6e0893b06a434e2872d7c7f9430394c9318f3537, ordered = TRUE),",
		"PGR_item.4_GROUP = col_factor(levels = scale_6fe6de0859f881e87e5656b0b361e2830bbbc058),",
		"PGR_item.4_RANK = col_factor(levels = scale_6e0893b06a434e2872d7c7f9430394c9318f3537, ordered = TRUE),",
		"PGR_other_GROUP = col_factor(levels = scale_6fe6de0859f881e87e5656b0b361e2830bbbc058),",
		"PGR_other_RANK = col_factor(levels = scale_6e0893b06a434e2872d7c7f9430394c9318f3537, ordered = TRUE),",
		"NPS = col_factor(),",
		"NPS_group = col_factor(),",
		"s = col_factor()",
		"))",
		"",
		"rm(input_path)",
		"rm(scale_0d33bdb7dd7ad7e7644895dab595541b141f5b39)",
		"rm(scale_30f77603f9a26644196c8b5400b99945c2c294a5)",
		"rm(scale_6e0893b06a434e2872d7c7f9430394c9318f3537)",
		"rm(scale_6fe6de0859f881e87e5656b0b361e2830bbbc058)",
		"rm(scale_8a3feac0fed5c3348d223a4c4a52d9b74e347e0a)",
		"rm(scale_9bc0385ea2c175f3341306637ae392b35bd86573)",
	}
	for row, test := range tests {
		line, err := r.ReadString('\n')
		if err != nil {
			t.Errorf("row %d, err = %s", row, err)
		}
		line = strings.TrimSpace(line)
		if line != test {
			t.Errorf("row %d, line = '%s'; wanted '%s'", row, line, test)
		}
	}
	line, err := r.ReadString('\n')
	if err != io.EOF {
		if err == nil {
			t.Errorf("line = %s; expected EOF", strings.TrimSpace(line))
		} else {
			t.Errorf("err = %s; expected EOF", err)
		}
	}
}

func TestWriteRNil(t *testing.T) {
	r := bufio.NewReader(strings.NewReader(qsfTestContent))
	s, err := ReadQsf(r)
	if err != nil {
		t.Errorf("err = %s", err)
	}
	if s == nil {
		t.Error("s = nil")
		return
	}
	err = s.WriteR(nil, "")
	if err != nil && err.Error() != "w cannot be nil" {
		t.Errorf("err = %s; want err = 'w cannot be nil'", err)
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
		return
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
