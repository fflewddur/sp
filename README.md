# SP (survey parser)

sp is a tool for parsing Qualtrics data into a format suitable for R.

[![Build Status](https://travis-ci.org/fflewddur/survey_parser.svg?branch=master)](https://travis-ci.org/fflewddur/survey_parser) [![codecov](https://codecov.io/gh/fflewddur/survey_parser/branch/master/graph/badge.svg)](https://codecov.io/gh/fflewddur/survey_parser) [![Go Report Card](https://goreportcard.com/badge/github.com/fflewddur/survey_parser)](https://goreportcard.com/report/github.com/fflewddur/survey_parser)

## Getting started

sp takes Qualtrics data and creates both a CSV and R import script.

1. Install sp with the command `go install github.com/fflewddur/sp` (you can [download Go here](https://golang.org/) if you don't already have it).

1. Export your Qualtrics survey as a QSF file (in Qualtrics: Survey &rarr; Tools &rarr; Import/Export &rarr; Export survey).

1. Export your Qualtrics responses as an XML file (in Qualtrics: Data & Analysis &rarr; Export & Import &rarr; Export data, select XML).

1. Rename both the QSF and XML files to have the same base name (e.g., _survey.qsf_ and _survey.xml_). Both files need to be in the same folder.

1. Run sp on your survey data with the command `sp <PATH_TO_QSF_FILE>`. For example, if you saved your QSF and XML files to ~/Downloads with the names _survey.qsf_ and _survey.xml_, you would run the command `sp ~/Downloads/survey.qsf`. sp will read the survey structure from the QSF file and participants' responses from the XML file. It will create two files: a CSV containing participants' responses, and an R script for importing the CSV into R. These files will be created in the same folder as the QSF file and share the same base name (e.g., running `sp ~/Downloads/survey.qsf` will create _survey.csv_ and _survey.r_ in your Downloads folder).

1. Import the data into R by running import script sp generated. Continuing the above example, we'd start R and run the command `source("survey.r")`.

1. (Optional) You can edit the generated R script as appropriate. By default it will define a type for each CSV column (logical, factor, integer, etc.) and include factor levels. For questions that allow multiple responses, logical columns for each response will be generated.

## Building

To build sp, run `go build ./cmd/sp` from the root of the repository.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
