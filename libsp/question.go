package libsp

import "fmt"

// Question represents a survey question
type Question struct {
	ID           string
	Wording      string
	qType        QType
	choices      []Choice
	subQuestions []Choice
	groups       []string
}

// Choice represents one possible response to a survey question
type Choice struct {
	ID      string
	Label   string
	HasText bool
}

// Type returns a QType representing the type of this survey question
func (q *Question) Type() QType {
	return q.qType
}

// ResponseChoices returns a slice of Choice holding the response choices available to survey respondents
func (q *Question) ResponseChoices() []Choice {
	return q.choices
}

// SubQuestions returns a slice of Choice holding the subquestions asked as part of this question (e.g., rows of a matrix question)
func (q *Question) SubQuestions() []Choice {
	return q.subQuestions
}

// CSVCols returns a slice of string holding the ordered CSV column names for this question
func (q *Question) CSVCols() []string {
	cols := make([]string, 0)
	suffixes := []string{}
	if q.qType == PickGroupRank {
		for _, c := range q.choices {
			suffixes = append(suffixes, fmt.Sprintf("_%s_GROUP", c.ID))
			suffixes = append(suffixes, fmt.Sprintf("_%s_RANK", c.ID))
			if c.HasText {
				suffixes = append(suffixes, "_"+c.ID+"_TEXT")
			}
		}
	} else {
		suffixes = q.qType.suffixes(q)
	}
	for _, s := range suffixes {
		cols = append(cols, q.ID+s)
	}

	return cols
}

// RColType returns the R type of columns associated with this question
func (q *Question) RColType() string {
	switch q.qType {
	case MatrixSingleResponse:
		fallthrough
	case MultipleChoiceSingleResponse:
		fallthrough
	case NPS:
		fallthrough
	case PickGroupRank:
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
		suffixes := q.qType.suffixes(q)
		for _, s := range suffixes {
			cols = append(cols, r.answers[q.ID+s])
		}
	}

	return cols
}

// groupsAndRanks returns the groups and ranks for each choice in this Question.
// PickGroupRank responses are structured differently than other question types,
// so it needs its own logic.
func (q *Question) groupsAndRanks(r *Response) []string {
	cols := make([]string, 0)
	for _, c := range q.choices {
		var group, rank string
		for gi, g := range q.groups {
			gk := fmt.Sprintf("%s_%d_GROUP_%s", q.ID, gi, c.ID)
			if v, ok := r.answers[gk]; ok && len(v) > 0 {
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
	return cols
}

func newQuestion(p *qsfPayload) (*Question, error) {
	q := new(Question)
	q.ID = p.QuestionID
	q.Wording = p.QuestionText
	q.qType = newQTypeFromString(p.QuestionType, p.Selector, p.SubSelector)
	var err error
	if q.qType.choicesAreQuestions() {
		q.subQuestions, err = p.OrderedChoices()
		if err != nil {
			return nil, fmt.Errorf("could not parse choices: %s", err)
		}
		q.choices, err = p.OrderedAnswers()
		if err != nil {
			return nil, fmt.Errorf("could not parse answers: %s", err)
		}
	} else {
		q.choices, err = p.OrderedChoices()
		if err != nil {
			return nil, fmt.Errorf("could not parse choices: %s", err)
		}
	}
	q.groups = p.Groups

	return q, nil
}

// QType represents the type of a survey question
type QType int

// Types of supported survey questions
const (
	Unknown QType = iota
	Description
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
)

// newQTypeFromString returns the corresponding QType value for the given string
func newQTypeFromString(t, s, ss string) QType {
	switch t {
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
	}

	return Unknown
}

func (qt QType) String() string {
	s := []string{
		"Unknown",
		"Description",
		"Form",
		"MultipleChoiceSingleResponse",
		"MultipleChoiceMultiResponse",
		"MatrixSingleResponse",
		"MatrixMultiResponse",
		"MaxDiff",
		"NPS",
		"PickGroupRank",
		"RankOrder",
		"TextEntry",
	}
	return s[qt]
}

// choicesAreQuestions returns true if the survey definition for this question uses the Choices field to hold question wording
func (qt QType) choicesAreQuestions() bool {
	retval := false
	switch qt {
	case MatrixSingleResponse:
		fallthrough
	case MatrixMultiResponse:
		retval = true
	}
	return retval
}

func (qt QType) suffixes(q *Question) []string {
	// These are the formats for the keys to lookup responses for each question type
	// Form: [question id]_[choice id]
	// MultipleChoiceSingleResponse: [question id]
	// MultipleChoiceMultiResponse: [question id]_[choice id]
	// MatrixSingleResponse: [question id]_[subquestion id]
	// MatrixMultiResponse: [question id]_[subquestion id]_[choice id]
	// MaxDiff: [question id]_[choice id]
	// PickGroupRank groupings: [question id]_[group index]_[choice id]
	// PickGroupRank rankings: [question id]_G[group index]_choice id]_RANK
	// RankOrder: [question id]_[choice id]
	// TextEntry: [question id]_TEXT
	// NPS: [question id] and [question id]_NPS_GROUP
	suffixes := []string{}
	switch qt {
	case Form, MaxDiff, MultipleChoiceMultiResponse, RankOrder:
		for _, c := range q.choices {
			suffixes = append(suffixes, "_"+c.ID)
			if c.HasText {
				suffixes = append(suffixes, "_"+c.ID+"_TEXT")
			}
		}
	case MatrixMultiResponse:
		for _, sq := range q.subQuestions {
			for _, c := range q.choices {
				suffixes = append(suffixes, "_"+sq.ID+"_"+c.ID)
			}
			if sq.HasText {
				suffixes = append(suffixes, "_"+sq.ID+"_TEXT")
			}
		}
	case MatrixSingleResponse:
		for _, sq := range q.subQuestions {
			suffixes = append(suffixes, "_"+sq.ID)
			if sq.HasText {
				suffixes = append(suffixes, "_"+sq.ID+"_TEXT")
			}
		}
	case MultipleChoiceSingleResponse:
		suffixes = append(suffixes, "")
		for _, c := range q.choices {
			if c.HasText {
				suffixes = append(suffixes, "_"+c.ID+"_TEXT")
			}
		}
	case NPS:
		suffixes = append(suffixes, "")
		suffixes = append(suffixes, "_NPS_GROUP")
	case PickGroupRank:
		for gi := range q.groups {
			for _, c := range q.choices {
				suffixes = append(suffixes, fmt.Sprintf("_%d_%s", gi, c.ID))
				suffixes = append(suffixes, fmt.Sprintf("_G%d_%s_RANK", gi, c.ID))
				if c.HasText {
					suffixes = append(suffixes, "_"+c.ID+"_TEXT")
				}
			}
		}
	case TextEntry:
		suffixes = append(suffixes, "_TEXT")
	}

	return suffixes
}
