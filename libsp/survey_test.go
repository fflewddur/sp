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
			"q1label", "q18label", "q3label",
			"q4label_1", "q4label_2", "q4label_3", "q4label_5", "q4label_5_text", "q4label_6", "q4label_6_text", "q4label_4",
			"q19_choice1", "q19_choice3", "q19_choice2", "q19_none", // 19
			"q11label", "q11label_3_text",
			"q5label_statement1", "q5label_statement2", "q5label_statement3", "q5label_other", "q5label_other_text",
			"q13label_1_1", "q13label_1_2", "q13label_1_3", "q13label_2_1", "q13label_2_2", "q13label_2_3", "q13label_3_1", "q13label_3_2", "q13label_3_3", "q13label_5_1", "q13label_5_2", "q13label_5_3", "q13label_4_1", "q13label_4_2", "q13label_4_3", "q13label_4_text", // 42
			"q6label_1", "q6label_2", "q6label_3",
			"q16_first_click", "q16_last_click", "q16_page_submit", "q16_click_count",
			"q7label_text", "q8label_text",
			"q9label_1", "q9label_2", "q9label_3",
			"q10label_1", "q10label_2", "q10label_3",
			"pgr_item_1_group", "pgr_item_1_rank", "pgr_item_2_group", "pgr_item_2_rank", "pgr_item_3_group", "pgr_item_3_rank", "pgr_item_4_group", "pgr_item_4_rank", "pgr_other_group", "pgr_other_rank", "pgr_other_text", // 68
			"q15label", "q15label_group",
			"q29_choice1", "q29_choice2", "q29_choice3_w_txt", "q29_choice3_w_txt_text", "q29_choice5",
			"q20", "q21",
			"q28_statement1", "q28_statement2", "q28_statement3", "q28_other", "q28_other_text", // 82
			"q23_1", "q23_2", "q23_3", "q23_4", "q23_4_text",
			"q26", "q27",
			"loop_base_1", "loop_base_2", "loop_base_3",
			"q31", "q32_1", "q32_2", "q32_3", "q33_text",
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
			"And let's try a smart quote for good measure", "Click to write Choice 3 (ordered 2nd)", "Click to write Choice 3",
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
		"library(tidyr)",
		"library(dplyr)",
		"library(readr)",
		"",
		"input_path <- \"test.csv\"",
		`scale_08ff20f7e2e706f820599f8da8a74b96eb33aa1f <- c("Click to write Choice 1", "Click to write Choice 2", "Click to write Choice 3", "Let's include a choice with an apostrophe", "And let's try a smart quote for good measure", "No response")`,
		`scale_0d33bdb7dd7ad7e7644895dab595541b141f5b39 <- c("Click to write Choice 1", "Click to write Choice 2", "Click to write Choice 3", "No response")`,
		`scale_30f77603f9a26644196c8b5400b99945c2c294a5 <- c("Click to write Choice 2 (ordered 1st)", "Click to write Choice 1 (ordered 2nd)", "Click to write Choice 3", "No response")`,
		`scale_37e352849a3d8bb86e939a98337b2c6229d54634 <- c("1", "2", "3", "4", "5", "Not grouped", "No response")`,
		`scale_39cfaa6cbb1b24d4b91d28093c45805a7fcb8e05 <- c("1", "2", "3", "No response")`,
		`scale_8a3feac0fed5c3348d223a4c4a52d9b74e347e0a <- c("Click to write Choice 2 (ordered 1st)", "Click to write Choice 3 (ordered 2nd)", "Click to write Choice 1 (ordered 3rd)", "No response")`,
		`scale_9bc0385ea2c175f3341306637ae392b35bd86573 <- c("Click to write Scale point 1", "Click to write Scale point 2", "Click to write Scale point 3", "n/a", "No response")`,
		`scale_a0e99a824c3c578ebc8d6823906c6e48d95cd2ae <- c("1", "2", "3", "4", "No response")`, // 10
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
		"q1label = col_factor(levels = scale_08ff20f7e2e706f820599f8da8a74b96eb33aa1f),",
		"q18label = col_factor(levels = scale_8a3feac0fed5c3348d223a4c4a52d9b74e347e0a),",
		"q3label = col_factor(levels = scale_0d33bdb7dd7ad7e7644895dab595541b141f5b39),",
		"q4label_1 = col_logical(),",
		"q4label_2 = col_logical(),",
		"q4label_3 = col_logical(),",
		"q4label_5 = col_logical(),",
		"q4label_6 = col_logical(),",
		"q4label_4 = col_logical(),",
		"q19_choice1 = col_logical(),",
		"q19_choice3 = col_logical(),",
		"q19_choice2 = col_logical(),",
		"q19_none = col_logical(),",
		"q11label = col_factor(levels = scale_30f77603f9a26644196c8b5400b99945c2c294a5),",
		"q5label_statement1 = col_factor(levels = scale_9bc0385ea2c175f3341306637ae392b35bd86573, ordered = TRUE),",
		"q5label_statement2 = col_factor(levels = scale_9bc0385ea2c175f3341306637ae392b35bd86573, ordered = TRUE),",
		"q5label_statement3 = col_factor(levels = scale_9bc0385ea2c175f3341306637ae392b35bd86573, ordered = TRUE),",
		"q5label_other = col_factor(levels = scale_9bc0385ea2c175f3341306637ae392b35bd86573, ordered = TRUE),",
		"q13label_1_1 = col_logical(),",
		"q13label_1_2 = col_logical(),",
		"q13label_1_3 = col_logical(),",
		"q13label_2_1 = col_logical(),",
		"q13label_2_2 = col_logical(),",
		"q13label_2_3 = col_logical(),",
		"q13label_3_1 = col_logical(),",
		"q13label_3_2 = col_logical(),",
		"q13label_3_3 = col_logical(),",
		"q13label_5_1 = col_logical(),",
		"q13label_5_2 = col_logical(),",
		"q13label_5_3 = col_logical(),",
		"q13label_4_1 = col_logical(),",
		"q13label_4_2 = col_logical(),",
		"q13label_4_3 = col_logical(),",
		"q6label_1 = col_logical(),",
		"q6label_2 = col_logical(),",
		"q6label_3 = col_logical(),",
		"q16_first_click = col_double(),",
		"q16_last_click = col_double(),",
		"q16_page_submit = col_double(),",
		"q16_click_count = col_integer(),",
		"q9label_1 = col_logical(),",
		"q9label_2 = col_logical(),",
		"q9label_3 = col_logical(),",
		"q10label_1 = col_factor(levels = scale_39cfaa6cbb1b24d4b91d28093c45805a7fcb8e05, ordered = TRUE),",
		"q10label_2 = col_factor(levels = scale_39cfaa6cbb1b24d4b91d28093c45805a7fcb8e05, ordered = TRUE),",
		"q10label_3 = col_factor(levels = scale_39cfaa6cbb1b24d4b91d28093c45805a7fcb8e05, ordered = TRUE),",
		"pgr_item_1_group = col_factor(levels = scale_dfbadf501868c43fd508372a48f65f9327d3c676),",
		"pgr_item_1_rank = col_factor(levels = scale_37e352849a3d8bb86e939a98337b2c6229d54634, ordered = TRUE),",
		"pgr_item_2_group = col_factor(levels = scale_dfbadf501868c43fd508372a48f65f9327d3c676),",
		"pgr_item_2_rank = col_factor(levels = scale_37e352849a3d8bb86e939a98337b2c6229d54634, ordered = TRUE),",
		"pgr_item_3_group = col_factor(levels = scale_dfbadf501868c43fd508372a48f65f9327d3c676),",
		"pgr_item_3_rank = col_factor(levels = scale_37e352849a3d8bb86e939a98337b2c6229d54634, ordered = TRUE),",
		"pgr_item_4_group = col_factor(levels = scale_dfbadf501868c43fd508372a48f65f9327d3c676),",
		"pgr_item_4_rank = col_factor(levels = scale_37e352849a3d8bb86e939a98337b2c6229d54634, ordered = TRUE),",
		"pgr_other_group = col_factor(levels = scale_dfbadf501868c43fd508372a48f65f9327d3c676),",
		"pgr_other_rank = col_factor(levels = scale_37e352849a3d8bb86e939a98337b2c6229d54634, ordered = TRUE),",
		"q15label = col_factor(),",
		"q15label_group = col_factor(),",
		"q29_choice1 = col_factor(levels = scale_a0e99a824c3c578ebc8d6823906c6e48d95cd2ae, ordered = TRUE),",
		"q29_choice2 = col_factor(levels = scale_a0e99a824c3c578ebc8d6823906c6e48d95cd2ae, ordered = TRUE),",
		"q29_choice3_w_txt = col_factor(levels = scale_a0e99a824c3c578ebc8d6823906c6e48d95cd2ae, ordered = TRUE),",
		"q29_choice5 = col_factor(levels = scale_a0e99a824c3c578ebc8d6823906c6e48d95cd2ae, ordered = TRUE),",
		"q20 = col_factor(levels = scale_ae8733afbe88aee2192428ceea072703a0de0e4e),",
		"q21 = col_factor(levels = scale_ae8733afbe88aee2192428ceea072703a0de0e4e),",
		"q28_statement1 = col_factor(levels = scale_a51a95e35530472ee800821ae86ba1bf3ff20b00),",
		"q28_statement2 = col_factor(levels = scale_a51a95e35530472ee800821ae86ba1bf3ff20b00),",
		"q28_statement3 = col_factor(levels = scale_a51a95e35530472ee800821ae86ba1bf3ff20b00),",
		"q28_other = col_factor(levels = scale_a51a95e35530472ee800821ae86ba1bf3ff20b00),",
		"q23_1 = col_double(),",
		"q23_2 = col_double(),",
		"q23_3 = col_double(),",
		"q23_4 = col_double(),",
		"q26 = col_factor(levels = scale_0d33bdb7dd7ad7e7644895dab595541b141f5b39),",
		"q27 = col_factor(levels = scale_0d33bdb7dd7ad7e7644895dab595541b141f5b39),",
		"loop_base_1 = col_logical(),",
		"loop_base_2 = col_logical(),",
		"loop_base_3 = col_logical(),",
		"q31 = col_factor(levels = scale_0d33bdb7dd7ad7e7644895dab595541b141f5b39),",
		"q32_1 = col_logical(),",
		"q32_2 = col_logical(),",
		"q32_3 = col_logical(),",
		"s = col_factor()",
		"))",
		"",
		"rm(input_path)",
		"rm(scale_08ff20f7e2e706f820599f8da8a74b96eb33aa1f)",
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
