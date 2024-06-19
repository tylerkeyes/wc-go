package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/tylerkeyes/assert"
)

var (
	cntWords = flag.Bool("w", false, "count words in file")
	cntLines = flag.Bool("l", false, "count lines in file")
	cntChars = flag.Bool("m", false, "count characters in file")
	cntBytes = flag.Bool("c", false, "count bytes in file")
	help     = flag.Bool("h", false, "help information")
)

func main() {
	flag.Parse()

	if *help {
		helpMsg := `ccwc - Coding Challenges WC tool
	Commands:

	-w		count words in the file
	-l		count lines in the file
	-m		count characters in the file
	-c		count bytes in the file
	`
		fmt.Printf(helpMsg)
		os.Exit(0)
	}

	var fileBytes []byte
	var fileData string
	var fileName string
	var err error

	if len(flag.Args()) < 1 {
		// reading file from standard input
		fileBytes, err = io.ReadAll(os.Stdin)
		assert.Assert(err)
	} else {
		fileName = flag.Args()[0]
		fileBytes, err = os.ReadFile(fileName)
		assert.AssertMsg(err, "error reading file")
	}

	var nBytes int
	var nWords int
	var nLines int
	var nChars int

	fileData = string(fileBytes)
	fileLines := strings.Split(fileData, "\n")

	// Process data
	if *cntBytes {
		nBytes = len(fileBytes)
	}
	if *cntLines {
		nLines = len(fileLines) - 1
	}
	if *cntWords {
		nWords = findWords(fileLines)
	}
	if *cntChars {
		nChars = findChars(string(fileBytes))
	}

	// Print results
	if *cntBytes && *cntLines && *cntWords {
		fmt.Printf("   %v   %v   %v %v\n", nLines, nWords, nBytes, fileName)
	} else if *cntBytes {
		fmt.Printf("   %v %v\n", nBytes, fileName)
	} else if *cntLines {
		fmt.Printf("   %v %v\n", nLines, fileName)
	} else if *cntWords {
		fmt.Printf("   %v %v\n", nWords, fileName)
	} else if *cntChars {
		fmt.Printf("   %v %v\n", nChars, fileName)
	}
}

func findWords(data []string) int {
	var regex = regexp.MustCompile(`\S+`)
	words := make([]string, 0)
	for _, line := range data {

		results := regex.FindAllString(line, -1)
		for _, res := range results {
			words = append(words, res)
		}
	}
	return len(words)
}

func findChars(data string) int {
	characters := make([]rune, 0)
	for _, char := range data {
		characters = append(characters, char)
	}
	return len(characters)
}
