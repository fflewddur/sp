package libsp

import (
	"io/ioutil"
	"log"
)

// ReadQsf reads a Qualtrics survey definition file (.qsf) from disk
func ReadQsf(path string) (survey *Survey, err error) {
	log.Printf("Reading '%s'", path)

	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading '%s': %s", path, err)
	}

	var s = new(Survey)
	if err := s.UnmarshalJSON(b); err != nil {
		log.Fatalf("Error parsing '%s': %s", path, err)
	}

	log.Printf("Title = %s, description = %s\n", s.Title, s.Description)
	survey = s
	return
}
