package libsp

// QType represents the type of a survey question
type QType int

// Types of supported survey questions
const (
	_                    = iota
	MultipleChoice QType = iota
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
func QTypeFromString(s string) QType {
	return MultipleChoice
}
