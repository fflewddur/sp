package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fflewddur/sp/libsp"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sp [qsf file]",
	Short: "sp is a survey parser for Qualtrics data",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		qsfPath := args[0]

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
		err = s.WriteR(w)
		log.Println("Completed successfully!")
		if err != nil {
			log.Fatalf("Error writing '%s': %s", rPath, err)
		}
	},
	Version: "0.0.1",
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

// Execute runs the command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
