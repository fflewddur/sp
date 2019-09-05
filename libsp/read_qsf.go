package libsp

import (
	"bufio"
	"errors"
	"fmt"
	"io"
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
	// TODO investigate using json.NewDecoder here (https://eli.thegreenplace.net/2019/go-json-cookbook/)
	// Would let us move all QSF logic from survey.go
	if err := s.UnmarshalJSON(bytes); err != nil {
		e := fmt.Errorf("could not parse: %s", err)
		return nil, e
	}

	return s, nil
}
