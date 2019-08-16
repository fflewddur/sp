package libsp

// Question represents a survey question
type Question struct {
	ID           string
	Wording      string
	qType        QType
	choices      []Choice
	subQuestions []Choice
}

// Choice represents one possible response to a survey question
type Choice struct {
	ID    string
	Label string
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
	suffixes := q.qType.suffixes(q)
	cols := make([]string, 0)
	for _, s := range suffixes {
		cols = append(cols, q.ID+s)
	}
	return cols
}

// ResponseCols returns a slice of string holding the ordered responses in r for this question
func (q *Question) ResponseCols(r *Response) []string {
	return nil
}

func newQuestion(p *qsfPayload) *Question {
	q := new(Question)
	q.ID = p.QuestionID
	q.Wording = p.QuestionText
	q.qType = newQTypeFromString(p.QuestionType, p.Selector)
	if q.qType.choicesAreQuestions() {
		q.subQuestions = p.OrderedChoices()
		q.choices = p.OrderedAnswers()
	} else {
		q.choices = p.OrderedChoices()
	}

	return q
}

// QType represents the type of a survey question
type QType int

// Types of supported survey questions
const (
	Unknown QType = iota
	Description
	Form
	MultipleChoiceSingleResponse
	MultipleChoiceMultiResponse
	MatrixSingleResponse
	MatrixMultiResponse
	MaxDiff
	RankOrder
	TextEntry
)

// newQTypeFromString returns the corresponding QType value for the given string
func newQTypeFromString(t, s string) QType {
	switch t {
	case "DB":
		return Description
	case "Matrix":
		if s == "MaxDiff" {
			return MaxDiff
		} else if s == "MultipleAnswer" {
			return MatrixMultiResponse
		}
		return MatrixSingleResponse
	case "MC":
		if s == "MAVR" || s == "MAHR" || s == "MACOL" || s == "MSB" {
			return MultipleChoiceMultiResponse
		}
		return MultipleChoiceSingleResponse
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
		fallthrough
	case MaxDiff:
		retval = true
	}
	return retval
}

func (qt QType) suffixes(q *Question) []string {
	// TODO support each type below:
	// Description: ?
	// Form: [question id]_[choice id]
	// MultipleChoiceSingleResponse: [question id]
	// MultipleChoiceMultiResponse: [question id]_[choice id]
	// MatrixSingleResponse: [question id]_[subquestion id]
	// MatrixMultiResponse: ?
	// MaxDiff: [question id]_[choice id]
	// RankOrder: [question id]_[choice id]
	// TextEntry: [question id]_TEXT
	suffixes := []string{}
	switch qt {
	case MultipleChoiceSingleResponse:
		suffixes = append(suffixes, "")
	case MultipleChoiceMultiResponse:
		fallthrough
	case Form:
		for _, c := range q.choices {
			suffixes = append(suffixes, "_"+c.ID)
		}
	}

	return suffixes
}
