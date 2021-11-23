package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fflewddur/sp/libsp"
)

type options struct {
	CompatNames bool
}

func main() {
	showVer := flag.Bool("v", false, "display version and exit")
	compatNames := flag.Bool("c", false, "use variable names compatible with releases < 0.2.2")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage:\n  sp [flags] <qsf file>\n\nFlags:\n")

		flag.PrintDefaults()
	}
	flag.Parse()

	if *showVer {
		fmt.Printf("sp version %s\n", libsp.Version)
		return
	}
	opt := options{}
	opt.CompatNames = *compatNames

	if opt.CompatNames {
		log.Println("Using compatible names")
	}

	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	qsfPath := flag.Args()[0]
	parseSurvey(qsfPath, opt)
}

func parseSurvey(qsfPath string, opt options) {
	log.Printf("Reading '%s'", qsfPath)
	qsf, err := os.Open(qsfPath)
	if err != nil {
		log.Fatalf("Error reading '%s': %s", qsfPath, err)
	}
	defer qsf.Close()

	qsfReader := bufio.NewReader(qsf)
	s, err := libsp.ReadQsf(qsfReader)
	if err != nil {
		log.Fatalf("Error parsing '%s': %s", qsfPath, err)
	}
	s.UseCompatNames = opt.CompatNames

	xmlPath := buildXMLPath(qsfPath)
	log.Printf("Reading '%s'", xmlPath)
	xml, err := os.Open(xmlPath)
	if err != nil {
		log.Fatalf("Error reading '%s': %s", xmlPath, err)
	}
	defer xml.Close()
	xmlReader := bufio.NewReader(xml)
	err = s.ReadXML(xmlReader)
	if err != nil {
		log.Fatalf("Error parsing '%s': %s", xmlPath, err)
	}

	// log.Printf("Title = '%s', # of questions = %d, description = '%s'\n", s.Title, len(s.Questions), s.Description)
	// for _, q := range s.Questions {
	// log.Printf("ID: %s Wording: %s\n\tChoices: %v", q.ID, q.Wording, q.ResponseChoices())
	// }

	csvPath := buildCSVPath(qsfPath)
	log.Printf("Writing '%s'", csvPath)
	csv, err := os.Create(csvPath)
	if err != nil {
		log.Fatalf("Error opening '%s': %s", csvPath, err)
	}
	defer csv.Close()
	w := bufio.NewWriter(csv)
	err = s.WriteCSV(w)
	if err != nil {
		log.Fatalf("Error writing '%s': %s", csvPath, err)
	}

	rPath := buildRPath(qsfPath)
	log.Printf("Writing '%s'", rPath)
	r, err := os.Create(rPath)
	if err != nil {
		log.Fatalf("Error opening '%s': %s", csvPath, err)
	}
	defer r.Close()
	w = bufio.NewWriter(r)
	err = s.WriteR(w, filepath.Base(csvPath))
	log.Println("Completed successfully!")
	if err != nil {
		log.Fatalf("Error writing '%s': %s", rPath, err)
	}
}

func buildXMLPath(qsfPath string) string {
	i := strings.LastIndex(qsfPath, ".")
	if i > 0 {
		qsfPath = qsfPath[:i]
	}
	return qsfPath + ".xml"
}

func buildCSVPath(qsfPath string) string {
	i := strings.LastIndex(qsfPath, ".")
	if i > 0 {
		qsfPath = qsfPath[:i]
	}
	return qsfPath + ".csv"
}

func buildRPath(qsfPath string) string {
	i := strings.LastIndex(qsfPath, ".")
	if i > 0 {
		qsfPath = qsfPath[:i]
	}
	return qsfPath + ".r"
}
