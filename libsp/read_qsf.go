package libsp

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
)

// ReadQsf reads a Qualtrics survey definition file (.qsf) from disk
func ReadQsf(r *bufio.Reader) (survey *Survey, err error) {
	if r == nil {
		e := errors.New("r cannot be nil")
		return nil, e
	}
	bytes := make([]byte, 0)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			e := fmt.Errorf("could not read file: %s", err)
			return nil, e
		}
		bytes = append(bytes, line...)
	}

	var s = new(Survey)
	if err := s.UnmarshalJSON(bytes); err != nil {
		e := fmt.Errorf("could not parse: %s", err)
		return nil, e
	}

	log.Printf("Title = '%s', # of questions = %d, description = '%s'\n", s.Title, len(s.Questions), s.Description)
	for _, q := range s.Questions {
		log.Printf("ID: %s Wording: %s\n", q.ID, q.Wording)
	}
	survey = s
	return
}
