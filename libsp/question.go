package libsp

import "log"

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

// Question represents a survey question
type Question struct {
	ID      string
	Wording string
	Type    QType
	Choices []string
}

// QTypeFromString returns the corresponding QType value for the given string
func QTypeFromString(t, s string) QType {
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
