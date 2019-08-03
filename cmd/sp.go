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
		file, err := os.Open(path) // For read access.
		if err != nil {
			log.Fatalf("Error reading '%s': %s", path, err)
		}
		reader := bufio.NewReader(file)
		libsp.ReadQsf(reader)
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
