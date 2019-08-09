package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/fflewddur/sp/libsp"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sp [file to parse]",
	Short: "sp is a survey parser for Qualtrics data",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		log.Printf("Reading '%s'", path)
		qsf, err := os.Open(path) // For read access.
		if err != nil {
			log.Fatalf("Error reading '%s': %s", path, err)
		}
		defer qsf.Close()

		qsfReader := bufio.NewReader(qsf)

		s, err := libsp.ReadQsf(qsfReader)
		if err != nil {
			log.Fatalf("Error parsing '%s': %s", path, err)
		}

		log.Printf("Title = '%s', # of questions = %d, description = '%s'\n", s.Title, len(s.Questions), s.Description)
		for _, q := range s.Questions {
			log.Printf("ID: %s Wording: %s\n\tChoices: %v", q.ID, q.Wording, q.ResponseChoices())
		}
	},
	Version: "0.0.1",
}
var inputFile string
var cfgFile string

// Execute runs the command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
