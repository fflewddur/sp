package libsp

import (
	"fmt"
	"regexp"
	"strings"
)

// Question represents a survey question
type Question struct {
	ID             string
	Wording        string
	label          string
	qType          QType
	choices        []Choice
	subQuestions   []Choice
	groups         []string
	orderedChoices bool
	dataExportTag  string
	dynChoices     *dynamicChoices
}

// Choice represents one possible response to a survey question
type Choice struct {
	ID      string
	Label   string
	VarName string // short variable name for use in analysis scripts
	HasText bool
}

type dynamicChoices struct {
	Source string
	Type   string
}

// Type returns a QType representing the type of this survey question
func (q *Question) Type() QType {
	return q.qType
}

// ResponseChoices returns a slice of Choice holding the ordered response choices available to survey respondents
func (q *Question) ResponseChoices() []Choice {
	return q.choices
}

// SubQuestions returns a slice of Choice holding the subquestions asked as part of this question (e.g., rows of a matrix question)
func (q *Question) SubQuestions() []Choice {
	return q.subQuestions
}

// OrderedChoices returns true if this question has RecodeValues set, indicating that order matters
func (q *Question) OrderedChoices() bool {
	return q.orderedChoices
}

// CSVCols returns a slice of string holding the ordered CSV column names for this question
func (q *Question) CSVCols() []string {
	cols := make([]string, 0)
	suffixes := []string{}
	if q.qType == PickGroupRank {
		for _, c := range q.choices {
			label := strings.ToLower(c.Label)
			r := regexp.MustCompile(`[\s]+`)
			label = r.ReplaceAllString(label, ".")
			label = strings.ReplaceAll(label, ":", "")
			suffixes = append(suffixes, fmt.Sprintf("_%s_GROUP", label))
			suffixes = append(suffixes, fmt.Sprintf("_%s_RANK", label))
			if c.HasText {
				suffixes = append(suffixes, "_"+label+"_text")
			}
		}
	} else {
		suffixes = q.qType.semanticSuffixes(q)
	}

	prefix := q.csvPrefix()
	for _, s := range suffixes {
		cols = append(cols, prefix+s)
	}

	// replace all non-R-compatible chars with '.'
	r := regexp.MustCompile(`[^a-zA-Z0-9_.]`)
	for i, c := range cols {
		cols[i] = r.ReplaceAllString(c, ".")
	}

	return cols
}

// CSVPrefix returns a string prefix for all CSV column names for this question
func (q *Question) csvPrefix() string {
	// FIXME len <= 30 is a hack; need to check if q.label is an ellipsized variant of q.Wording
	if q.label == "" || q.label == q.Wording || len(q.label) > 30 {
		if q.dataExportTag != "" {
			return q.dataExportTag
		}
		return q.ID
	}
	return q.label
}

// RColType returns the R type of columns associated with this question
func (q *Question) RColType() string {
	switch q.qType {
	case ConstantSum:
		return "col_integer()"
	case Embedded, MatrixSingleResponse, MultipleChoiceSingleResponse, NPS, PickGroupRank, RankOrder:
		return "col_factor()"
	}
	return "col_logical()"
}

// ResponseCols returns a slice of string holding the ordered responses in r for this question
func (q *Question) ResponseCols(r *Response) []string {
	cols := make([]string, 0)
	if q.qType == PickGroupRank {
		cols = q.groupsAndRanks(r)
	} else {
		suffixes := q.qType.internalSuffixes(q)
		allEmpty := true
		// First, check to see if the user answered any part of this question
		for _, s := range suffixes {
			a := r.answers[q.ID+s]
			if a != "" && a != noResponseCode {
				allEmpty = false
			}
		}

		for _, s := range suffixes {
			var col string
			if allEmpty {
				// If the user didn't answer this question, all of its columns should be NA
				col = ""
			} else {
				// If the user answered this question, any unchecked options should be FALSE
				col = q.formatResponseForCol(r.answers[q.ID+s])
			}
			cols = append(cols, col)
		}
	}

	return cols
}

// groupsAndRanks returns the groups and ranks for each choice in this Question.
// PickGroupRank responses are structured differently than other question types,
// so it needs its own logic.
func (q *Question) groupsAndRanks(r *Response) []string {
	cols := make([]string, 0)
	allEmpty := true
	for _, c := range q.choices {
		var group, rank string
		for gi, g := range q.groups {
			gk := fmt.Sprintf("%s_%d_GROUP_%s", q.ID, gi, c.ID)
			if v, ok := r.answers[gk]; ok && len(v) > 0 && v != noResponseCode {
				allEmpty = false
				group = g
				rk := fmt.Sprintf("%s_G%d_%s_RANK", q.ID, gi, c.ID)
				if v, ok := r.answers[rk]; ok && len(v) > 0 {
					rank = v
				}
			}
		}
		cols = append(cols, group)
		cols = append(cols, rank)

		if c.HasText {
			cols = append(cols, r.answers[q.ID+"_"+c.ID+"_TEXT"])
		}
	}
	if !allEmpty {
		// This response includes non-NA values, so change all "" to "Not grouped"
		// (this lets us differentiate "not selected" from "NA" during analysis)
		colsWithNA := []string{}
		for _, v := range cols {
			if v == "" {
				v = notGrouped
			}
			colsWithNA = append(colsWithNA, v)
		}
		cols = colsWithNA
	}
	return cols
}

func (q *Question) formatResponseForCol(userAnswer string) string {
	retval := userAnswer
	freeText := userAnswer != noResponseCode // true unless the user skipped this question
	for _, c := range q.choices {
		if c.Label == userAnswer || c.VarName == userAnswer {
			// set to false if the user's answer matches one of the question's choices
			freeText = false
			break
		}
	}

	if q.qType.exportAsBools() && !freeText {
		if userAnswer == noResponseCode || userAnswer == "" {
			retval = "FALSE"
		} else {
			retval = "TRUE"
		}
	} else if userAnswer == noResponseCode && q.qType != Timing {
		retval = noResponseConst
	}

	return retval
}

func newQuestionFromPayload(p *qsfPayload) (*Question, error) {
	q := new(Question)
	q.ID = p.QuestionID
	q.Wording = p.QuestionText
	q.dataExportTag = p.DataExportTag
	if p.QuestionDescription != "" {
		q.label = p.QuestionDescription
	}
	q.qType = newQTypeFromString(p.QuestionType, p.Selector, p.SubSelector)
	var err error
	if q.qType.choicesAreQuestions() {
		q.subQuestions, err = p.OrderedChoices(q.qType.choicesAreQuestions())
		if err != nil {
			return nil, fmt.Errorf("could not parse choices: %s", err)
		}
		q.choices, err = p.OrderedAnswers()
		if err != nil {
			return nil, fmt.Errorf("could not parse answers: %s", err)
		}
	} else {
		q.choices, err = p.OrderedChoices(q.qType.choicesAreQuestions())
		if err != nil {
			return nil, fmt.Errorf("could not parse choices: %s", err)
		}
	}
	q.groups = p.Groups
	// If we've recoded values in Qualtrics, that means order matters
	q.orderedChoices = p.RecodeValues != nil

	// If this question uses dynamic choices, save that information.
	// We'll parse it on a second pass, since it might refer to a question
	// we haven't parsed yet.
	if p.DynamicChoices != nil {
		q.dynChoices = new(dynamicChoices)
		q.dynChoices.Source = getSourceFromLocator(p.DynamicChoices.Locator)
		q.dynChoices.Type = getTypeFromLocator(p.DynamicChoices.Locator)
	}

	return q, nil
}

func getSourceFromLocator(l string) string {
	re := regexp.MustCompile(`(QID[0-9]+)`)
	matches := re.FindStringSubmatch(l)
	if matches != nil {
		return matches[1]
	}
	return ""
}

func getTypeFromLocator(l string) string {
	re := regexp.MustCompile(`/QID[0-9]+\/.+\/(.+)`)
	matches := re.FindStringSubmatch(l)
	if matches != nil {
		return matches[1]
	}
	return ""
}

func newQuestionFromEmbeddedData(d *qsfEmbeddedData) (*Question, error) {
	q := new(Question)
	q.qType = Embedded
	q.ID = d.Field
	q.label = d.Field

	return q, nil
}

// QType represents the type of a survey question
type QType int

// Types of supported survey questions
const (
	Unknown QType = iota
	ConstantSum
	Description
	Embedded
	Form
	Meta
	MultipleChoiceSingleResponse
	MultipleChoiceMultiResponse
	MatrixSingleResponse
	MatrixMultiResponse
	MaxDiff
	NPS
	PickGroupRank
	RankOrder
	TextEntry
	Timing
)

// newQTypeFromString returns the corresponding QType value for the given string
func newQTypeFromString(t, s, ss string) QType {
	switch t {
	case "CS":
		return ConstantSum
	case "DB":
		return Description
	case "Matrix":
		if s == "MaxDiff" {
			return MaxDiff
		} else if ss == "MultipleAnswer" {
			return MatrixMultiResponse
		}
		return MatrixSingleResponse
	case "Meta":
		return Meta
	case "MC":
		if s == "MAVR" || s == "MAHR" || s == "MACOL" || s == "MSB" {
			return MultipleChoiceMultiResponse
		} else if s == "NPS" {
			return NPS
		}
		return MultipleChoiceSingleResponse
	case "PGR":
		return PickGroupRank
	case "RO":
		return RankOrder
	case "TE":
		if s == "FORM" {
			return Form
		}
		return TextEntry
	case "Timing":
		return Timing
	}

	return Unknown
}

func (qt QType) String() string {
	s := []string{
		"Unknown",
		"ConstantSum",
		"Description",
		"Embedded",
		"Form",
		"Meta",
		"MultipleChoiceSingleResponse",
		"MultipleChoiceMultiResponse",
		"MatrixSingleResponse",
		"MatrixMultiResponse",
		"MaxDiff",
		"NPS",
		"PickGroupRank",
		"RankOrder",
		"TextEntry",
		"Timing",
	}
	return s[qt]
}

func (qt QType) exportAsBools() bool {
	switch qt {
	case MultipleChoiceMultiResponse, MatrixMultiResponse:
		return true
	}
	return false
}

// choicesAreQuestions returns true if the survey definition for this question uses the Choices field to hold question wording
func (qt QType) choicesAreQuestions() bool {
	retval := false
	switch qt {
	case ConstantSum, MatrixMultiResponse, MatrixSingleResponse:
		retval = true
	}
	return retval
}

func (qt QType) semanticSuffixes(q *Question) []string {
	return qt.suffixes(q, true)
}

func (qt QType) internalSuffixes(q *Question) []string {
	return qt.suffixes(q, false)
}

func (qt QType) suffixes(q *Question, useExportTags bool) []string {
	// These are the formats for the keys to lookup responses for each question type
	// ConstantSum: [question id]_[choice id]
	// Form: [question id]_[choice id]
	// MultipleChoiceSingleResponse: [question id]
	// MultipleChoiceMultiResponse: [question id]_[choice id]
	// MatrixSingleResponse: [question id]_[subquestion id]
	// MatrixMultiResponse: [question id]_[subquestion id]_[choice id]
	// MaxDiff: [question id]_[choice id]
	// PickGroupRank groupings: [question id]_[group index]_[choice id] (handled by groupsAndRanks())
	// PickGroupRank rankings: [question id]_G[group index]_choice id]_RANK (handled by groupsAndRanks())
	// RankOrder: [question id]_[choice id]
	// TextEntry: [question id]_TEXT
	// NPS: [question id] and [question id]_NPS_GROUP

	textSuffix := "_TEXT"
	npsSuffix := "_NPS_GROUP"
	timingSuffixes := []string{"_FIRST_CLICK", "_LAST_CLICK", "_PAGE_SUBMIT", "_CLICK_COUNT"}

	if useExportTags {
		textSuffix = strings.ToLower(textSuffix)
		npsSuffix = "_group"
		for i, s := range timingSuffixes {
			timingSuffixes[i] = strings.ToLower(s)
		}
	}

	suffixes := []string{}
	switch qt {
	case Embedded:
		suffixes = append(suffixes, "")
	case Form, MaxDiff, MultipleChoiceMultiResponse, RankOrder:
		for _, c := range q.choices {
			s := suffix(c, useExportTags)
			suffixes = append(suffixes, "_"+s)
			if c.HasText {
				suffixes = append(suffixes, "_"+s+textSuffix)
			}
		}
	case MatrixMultiResponse:
		for _, sq := range q.subQuestions {
			for _, c := range q.choices {
				suffixes = append(suffixes, "_"+sq.ID+"_"+c.ID)
			}
			if sq.HasText {
				suffixes = append(suffixes, "_"+sq.ID+textSuffix)
			}
		}
	case ConstantSum, MatrixSingleResponse:
		for _, sq := range q.subQuestions {
			s := suffix(sq, useExportTags)
			suffixes = append(suffixes, "_"+s)
			if sq.HasText {
				suffixes = append(suffixes, "_"+s+textSuffix)
			}
		}
	case MultipleChoiceSingleResponse:
		suffixes = append(suffixes, "")
		for _, c := range q.choices {
			if c.HasText {
				suffixes = append(suffixes, "_"+c.ID+textSuffix)
			}
		}
	case NPS:
		suffixes = append(suffixes, "")
		suffixes = append(suffixes, npsSuffix)
	case TextEntry:
		suffixes = append(suffixes, textSuffix)
	case Timing:
		suffixes = append(suffixes, timingSuffixes...)
	}

	return suffixes
}

func suffix(c Choice, useExportTags bool) string {
	if useExportTags && c.VarName != "" {
		return c.VarName
	}
	return c.ID
}
