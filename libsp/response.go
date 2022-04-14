package libsp

import (
	"log"
	"regexp"
	"strings"
	"time"
)

// Response models a Qualtrics participant response
type Response struct {
	ID         string
	Progress   int
	Duration   int
	Finished   bool
	RecordedOn time.Time
	answers    map[string]string
}

// NewResponse creates and initializes a Response
func NewResponse() *Response {
	var r Response
	r.answers = make(map[string]string)
	return &r
}

var reQIDLoop = regexp.MustCompile(`^_\d*_(QID\d+[^-]*)(-\d+)?$`) //(`^_\d*_(QID\d+.*)(-\d+)?`)
var reTimer = regexp.MustCompile(`_(CLICK|SUBMIT|COUNT)$`)

// AddAnswer adds a question answer to the response
func (r *Response) AddAnswer(id string, answer string) {
	answer = fixSmartQuotes(answer)

	// Remove the extraneous characters in loop+merge response IDs
	// TODO this probably doesn't work for all possible uses of loop+merge
	// FIXME this is broken for the 2018 Go survey
	origID := id
	matches := reQIDLoop.FindStringSubmatch(id)
	if matches != nil {
		// Don't merge all of the timer responses
		timerMatches := reTimer.MatchString(matches[1])
		if !timerMatches {
			id = matches[1]
		}
	}

	v, alreadySet := r.answers[id]
	if alreadySet {
		if answer != "" && v != "" && v != answer {
			log.Fatalf("error adding '%s' response for question '%s': already have '%s' (id=%s)", answer, id, v, origID)
		} else if v == "" && v != answer {
			r.answers[id] = answer
		}
	} else {
		r.answers[id] = answer
	}
}

func fixSmartQuotes(s string) string {
	// Replace smart quotes, R chokes on these
	s = strings.ReplaceAll(s, "“", "\"")
	s = strings.ReplaceAll(s, "”", "\"")
	s = strings.ReplaceAll(s, "‘", "'")
	s = strings.ReplaceAll(s, "’", "'")
	return s
}
