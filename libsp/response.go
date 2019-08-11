package libsp

// Response models a Qualtrics participant response
type Response struct {
	ID       string
	Progress int
	Duration int
	Finished bool
}
