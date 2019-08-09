package libsp

import "log"

// Question represents a survey question
type Question struct {
	ID           string
	Wording      string
	Type         QType
	Choices      []string
	subQuestions []string
}

func newQuestion(p *qsfPayload) *Question {
	q := new(Question)
	q.ID = p.QuestionID
	q.Wording = p.QuestionText
	q.Type = newQTypeFromString(p.QuestionType, p.Selector)
	if q.Type.choicesAreQuestions() {
		q.subQuestions = p.OrderedChoices()
		q.Choices = p.OrderedAnswers()
	} else {
		q.Choices = p.OrderedChoices()
	}

	return q
}

// QType represents the type of a survey question
type QType int

// Types of supported survey questions
const (
	Unknown QType = iota
	MultipleChoice
	TextEntry
	Matrix
	MaxDiff
	RankOrder
)

// newQTypeFromString returns the corresponding QType value for the given string
func newQTypeFromString(t, s string) QType {
	switch t {
	case "Matrix":
		if s == "MaxDiff" {
			return MaxDiff
		}
		return Matrix
	case "MC":
		return MultipleChoice
	case "RO":
		return RankOrder
	case "TE":
		return TextEntry
	default:
		log.Fatalf("'%s' is not a valid question type", t)
	}

	return Unknown
}

// choicesAreQuestions returns true if the survey definition for this question uses the Choices field to hold question wording
func (qt QType) choicesAreQuestions() bool {
	retval := false
	switch qt {
	case Matrix:
		retval = true
	case MaxDiff:
		retval = true
	}
	return retval
}
