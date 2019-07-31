package libsp

// Question represents a survey question
type Question struct {
	ID      string
	Wording string
	Type    string
	Choices []string
}
