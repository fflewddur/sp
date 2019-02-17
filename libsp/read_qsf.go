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

	log.Printf("Title = '%s', # of questions = %d, description = '%s'\n", s.Title, len(s.Questions), s.Description)
	for _, q := range s.Questions {
		log.Printf("ID: %s Wording: %s\n", q.ID, q.Wording)
	}
	survey = s
	return
}
