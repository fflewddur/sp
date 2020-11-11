package libsp

// spell-checker: disable

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
			"Q19_choice1", "Q19_choice3", "Q19_choice2", "Q19_none", // 19
			"Q11Label", "Q11Label_3_text",
			"Q5Label_statement1", "Q5Label_statement2", "Q5Label_statement3", "Q5Label_other", "Q5Label_other_text",
			"Q13Label_1_1", "Q13Label_1_2", "Q13Label_1_3", "Q13Label_2_1", "Q13Label_2_2", "Q13Label_2_3", "Q13Label_3_1", "Q13Label_3_2", "Q13Label_3_3", "Q13Label_5_1", "Q13Label_5_2", "Q13Label_5_3", "Q13Label_4_1", "Q13Label_4_2", "Q13Label_4_3", "Q13Label_4_text", // 42
			"Q6Label_1", "Q6Label_2", "Q6Label_3",
			"Q16_first_click", "Q16_last_click", "Q16_page_submit", "Q16_click_count",
			"Q7Label_text", "Q8Label_text",
			"Q9Label_1", "Q9Label_2", "Q9Label_3",
			"Q10Label_1", "Q10Label_2", "Q10Label_3",
			"pgr_item.1_GROUP", "pgr_item.1_RANK", "pgr_item.2_GROUP", "pgr_item.2_RANK", "pgr_item.3_GROUP", "pgr_item.3_RANK", "pgr_item.4_GROUP", "pgr_item.4_RANK", "pgr_other_GROUP", "pgr_other_RANK", "pgr_other_text", // 68
			"Q15Label", "Q15Label_group",
			"Q29_choice1", "Q29_choice2", "Q29_choice3_w_txt", "Q29_choice3_w_txt_text", "Q29_choice5",
			"Q20", "Q21",
			"Q28_statement1", "Q28_statement2", "Q28_statement3", "Q28_other", "Q28_other_text", // 82
			"Q23_1", "Q23_2", "Q23_3", "Q23_4", "Q23_4_text",
			"Q26", "Q27",
			"loop.base_1", "loop.base_2", "loop.base_3",
			"Q31", "Q32_1", "Q32_2", "Q32_3", "Q33_text",
			"s"},
		{"R_1dtWhiBDD96nfyk", "true", "100", "122", "2019-08-20 12:44:31",
			"Click to write Choice 1", "", "Click to write Choice 2",
			"FALSE", "FALSE", "TRUE", "TRUE", "other response 1", "TRUE", "other response 2", "FALSE",
			"FALSE", "TRUE", "FALSE", "FALSE", // 19
			"Click to write Choice 2 (ordered 1st)", "",
			"scale1", "scale2", "scale3", "scale3", "other matrix row",
			"TRUE", "TRUE", "FALSE", "FALSE", "TRUE", "TRUE", "FALSE", "TRUE", "FALSE", "FALSE", "TRUE", "FALSE", "FALSE", "FALSE", "TRUE", "other matrix multiple row", // 42
			"Click to write Scale point 1", "No response", "Click to write Scale point 2",
			"1.313", "32.89", "33.84", "15",
			"one line of text", "multiple\nlines\nof\ntext?",
			"field 1", "field 2", "field 3",
			"3", "2", "1",
			"Group 1", "3", "Group 2", "1", "Group 1", "2", "Group 1", "1", "Group 3", "1", "in group 3", // 68
			"6", "Detractor",
			"1", "2", "3", "testing", "4",
			"", "",
			"Click to write Scale Point 1", "Click to write Scale Point 2", "Click to write Scale Point 3", "Click to write Scale Point 1", "", // 82
			"", "", "", "", "",
			"", "",
			"", "", "",
			"", "", "", "", "",
			"g",
		},
		{"R_z72KJQMnr3lxZGp", "true", "100", "104", "2019-08-20 12:46:35",
			"Click to write Choice 3", "", "Click to write Choice 2",
			"FALSE", "FALSE", "FALSE", "FALSE", "", "FALSE", "", "TRUE",
			"FALSE", "FALSE", "FALSE", "TRUE", // 19
			"Click to write Choice 3", "other text",
			"scale3", "scale2", "", "", "",
			"TRUE", "FALSE", "FALSE", "FALSE", "FALSE", "FALSE", "TRUE", "TRUE", "TRUE", "FALSE", "FALSE", "FALSE", "FALSE", "FALSE", "FALSE", "", // 42
			"", "Click to write Scale point 2", "Click to write Scale point 1",
			"3.172", "25.605", "26.387", "9",
			"foo", "bar",
			"name", "email", "job role",
			"1", "2", "3",
			"Group 1", "2", "Group 1", "1", "Group 2", "2", "Group 2", "1", "Not grouped", "Not grouped", "Not grouped", // 68
			"10", "Promoter",
			"", "", "", "", "",
			"", "",
			"", "", "", "", "", // 82
			"", "", "", "", "",
			"", "",
			"", "", "",
			"", "", "", "", "",
			"e",
		},
		{"R_3MPTb9vwnCBmijR", "false", "33", "22", "2019-08-20 12:52:35",
			"Click to write Choice 2", "", "Click to write Choice 2",
			"FALSE", "TRUE", "FALSE", "FALSE", "", "FALSE", "", "FALSE",
			"TRUE", "FALSE", "TRUE", "FALSE", // 19
			"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
			"", "", "", "",
			"", "", "", "", "", "", "", "", "", "", "",
			"", "", "", "",
			"", "", "", "", "",
			"", "",
			"", "", "", "", "", // 82
			"", "", "", "", "",
			"", "",
			"", "", "",
			"", "", "", "", "",
			"",
		},
		{"R_2EzY1K5pqRpzi0n", "true", "100", "140", "2020-11-09 13:12:11",
			"Click to write Choice 1", "choice3", "Click to write Choice 3",
			"TRUE", "FALSE", "TRUE", "FALSE", "", "TRUE", "other 2 text", "FALSE",
			"TRUE", "FALSE", "TRUE", "FALSE", // 19
			"Click to write Choice 3", "other text",
			"scale1", "scale2", "scale3", "No response", "",
			"TRUE", "TRUE", "FALSE", "TRUE", "FALSE", "TRUE", "FALSE", "TRUE", "FALSE", "FALSE", "FALSE", "FALSE", "FALSE", "FALSE", "FALSE", "", // 42
			"Click to write Scale point 1", "Click to write Scale point 2", "No response",
			"1.717", "15.783", "16.628", "10",
			"line of text", "", // 51
			"form 1", "form 2", "form 3",
			"3", "2", "1",
			"Group 1", "1", "Group 2", "1", "Group 3", "1", "Not grouped", "Not grouped", "Not grouped", "Not grouped", "", // 68
			"5", "Detractor",
			"1", "2", "3", "", "",
			"Dyna choice 2", "Dyna choice 2",
			"Click to write Scale Point 1", "Click to write Scale Point 2", "Click to write Scale Point 3", "Click to write Scale Point 1", "other text", // 82
			"10", "20", "30", "0", "",
			"", "Click to write Choice 2",
			"TRUE", "TRUE", "TRUE", // 92
			"Click to write Choice 1", "TRUE", "TRUE", "FALSE", "choice 3 text",
			"",
		},
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
		"# Generated by sp " + Version + " (https://github.com/fflewddur/sp)",
		"library(tidyverse)",
		"",
		"input_path <- \"test.csv\"",
		`scale_0d33bdb7dd7ad7e7644895dab595541b141f5b39 <- c("Click to write Choice 1", "Click to write Choice 2", "Click to write Choice 3", "No response")`,
		`scale_30f77603f9a26644196c8b5400b99945c2c294a5 <- c("Click to write Choice 2 (ordered 1st)", "Click to write Choice 1 (ordered 2nd)", "Click to write Choice 3", "No response")`,
		`scale_37e352849a3d8bb86e939a98337b2c6229d54634 <- c("1", "2", "3", "4", "5", "Not grouped", "No response")`,
		`scale_39cfaa6cbb1b24d4b91d28093c45805a7fcb8e05 <- c("1", "2", "3", "No response")`,
		`scale_8a3feac0fed5c3348d223a4c4a52d9b74e347e0a <- c("choice2", "choice3", "choice1", "No response")`,
		`scale_9bc0385ea2c175f3341306637ae392b35bd86573 <- c("scale1", "scale2", "scale3", "scale.na", "No response")`,
		`scale_a0e99a824c3c578ebc8d6823906c6e48d95cd2ae <- c("1", "2", "3", "4", "No response")`,
		`scale_a51a95e35530472ee800821ae86ba1bf3ff20b00 <- c("Click to write Scale Point 1", "Click to write Scale Point 2", "Click to write Scale Point 3", "No response")`,
		`scale_ae8733afbe88aee2192428ceea072703a0de0e4e <- c("Dyna choice 1", "Dyna choice 2", "Dyna choice 3", "No response")`,
		`scale_dfbadf501868c43fd508372a48f65f9327d3c676 <- c("Group 1", "Group 2", "Group 3", "Not grouped", "No response")`,
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
		"Q10Label_1 = col_factor(levels = scale_39cfaa6cbb1b24d4b91d28093c45805a7fcb8e05, ordered = TRUE),",
		"Q10Label_2 = col_factor(levels = scale_39cfaa6cbb1b24d4b91d28093c45805a7fcb8e05, ordered = TRUE),",
		"Q10Label_3 = col_factor(levels = scale_39cfaa6cbb1b24d4b91d28093c45805a7fcb8e05, ordered = TRUE),",
		"pgr_item.1_GROUP = col_factor(levels = scale_dfbadf501868c43fd508372a48f65f9327d3c676),",
		"pgr_item.1_RANK = col_factor(levels = scale_37e352849a3d8bb86e939a98337b2c6229d54634, ordered = TRUE),",
		"pgr_item.2_GROUP = col_factor(levels = scale_dfbadf501868c43fd508372a48f65f9327d3c676),",
		"pgr_item.2_RANK = col_factor(levels = scale_37e352849a3d8bb86e939a98337b2c6229d54634, ordered = TRUE),",
		"pgr_item.3_GROUP = col_factor(levels = scale_dfbadf501868c43fd508372a48f65f9327d3c676),",
		"pgr_item.3_RANK = col_factor(levels = scale_37e352849a3d8bb86e939a98337b2c6229d54634, ordered = TRUE),",
		"pgr_item.4_GROUP = col_factor(levels = scale_dfbadf501868c43fd508372a48f65f9327d3c676),",
		"pgr_item.4_RANK = col_factor(levels = scale_37e352849a3d8bb86e939a98337b2c6229d54634, ordered = TRUE),",
		"pgr_other_GROUP = col_factor(levels = scale_dfbadf501868c43fd508372a48f65f9327d3c676),",
		"pgr_other_RANK = col_factor(levels = scale_37e352849a3d8bb86e939a98337b2c6229d54634, ordered = TRUE),",
		"Q15Label = col_factor(),",
		"Q15Label_group = col_factor(),",
		"Q29_choice1 = col_factor(levels = scale_a0e99a824c3c578ebc8d6823906c6e48d95cd2ae, ordered = TRUE),",
		"Q29_choice2 = col_factor(levels = scale_a0e99a824c3c578ebc8d6823906c6e48d95cd2ae, ordered = TRUE),",
		"Q29_choice3_w_txt = col_factor(levels = scale_a0e99a824c3c578ebc8d6823906c6e48d95cd2ae, ordered = TRUE),",
		"Q29_choice5 = col_factor(levels = scale_a0e99a824c3c578ebc8d6823906c6e48d95cd2ae, ordered = TRUE),",
		"Q20 = col_factor(levels = scale_ae8733afbe88aee2192428ceea072703a0de0e4e),",
		"Q21 = col_factor(levels = scale_ae8733afbe88aee2192428ceea072703a0de0e4e),",
		"Q28_statement1 = col_factor(levels = scale_a51a95e35530472ee800821ae86ba1bf3ff20b00),",
		"Q28_statement2 = col_factor(levels = scale_a51a95e35530472ee800821ae86ba1bf3ff20b00),",
		"Q28_statement3 = col_factor(levels = scale_a51a95e35530472ee800821ae86ba1bf3ff20b00),",
		"Q28_other = col_factor(levels = scale_a51a95e35530472ee800821ae86ba1bf3ff20b00),",
		"Q23_1 = col_double(),",
		"Q23_2 = col_double(),",
		"Q23_3 = col_double(),",
		"Q23_4 = col_double(),",
		"Q26 = col_factor(levels = scale_0d33bdb7dd7ad7e7644895dab595541b141f5b39),",
		"Q27 = col_factor(levels = scale_0d33bdb7dd7ad7e7644895dab595541b141f5b39),",
		"loop.base_1 = col_logical(),",
		"loop.base_2 = col_logical(),",
		"loop.base_3 = col_logical(),",
		"Q31 = col_factor(levels = scale_0d33bdb7dd7ad7e7644895dab595541b141f5b39),",
		"Q32_1 = col_logical(),",
		"Q32_2 = col_logical(),",
		"Q32_3 = col_logical(),",
		"s = col_factor()",
		"))",
		"",
		"rm(input_path)",
		"rm(scale_0d33bdb7dd7ad7e7644895dab595541b141f5b39)",
		"rm(scale_30f77603f9a26644196c8b5400b99945c2c294a5)",
		"rm(scale_37e352849a3d8bb86e939a98337b2c6229d54634)",
		"rm(scale_39cfaa6cbb1b24d4b91d28093c45805a7fcb8e05)",
		"rm(scale_8a3feac0fed5c3348d223a4c4a52d9b74e347e0a)",
		"rm(scale_9bc0385ea2c175f3341306637ae392b35bd86573)",
		"rm(scale_a0e99a824c3c578ebc8d6823906c6e48d95cd2ae)",
		"rm(scale_a51a95e35530472ee800821ae86ba1bf3ff20b00)",
		"rm(scale_ae8733afbe88aee2192428ceea072703a0de0e4e)",
		"rm(scale_dfbadf501868c43fd508372a48f65f9327d3c676)",
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

	if len(s.Responses) != 4 {
		t.Errorf("len(Responses) = %d; wanted 4", len(s.Responses))
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

// spell-checker: enable
