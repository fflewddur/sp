package libsp

// Response models a Qualtrics participant response
type Response struct {
	ID       string
	Progress int
	Duration int
	Finished bool
	answers  map[string]string
}

// NewResponse creates and initializes a Response
func NewResponse() *Response {
	var r Response
	r.answers = make(map[string]string)
	return &r
}

// AddAnswer adds a question answer to the response
func (r *Response) AddAnswer(id string, answer string) {
	r.answers[id] = answer
}
