package libsp

import (
	"bufio"
	"errors"
	"fmt"
	"html"
	"io"
	"strings"
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
		s := html.UnescapeString(string(line))
		// And fix spaces, too
		s = strings.ReplaceAll(s, "&nbsp;", " ")

		line = []byte(s)
		bytes = append(bytes, line...)
	}

	var s = new(Survey)
	if err := s.UnmarshalJSON(bytes); err != nil {
		e := fmt.Errorf("could not parse: %s", err)
		return nil, e
	}

	return s, nil
}
