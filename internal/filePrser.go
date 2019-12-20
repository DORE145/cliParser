package internal

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/DORE145/cliParser/internal/data"
)

//ParsingResult stores resuls of the file parsing
//Field names corresponds the flags used to configure parser
type ParsingResult struct {
	LinesCounter   int
	WordsCounter   int
	SymbolsCounter int
	UniqueWords    data.WordSet
}

//Parse function parses incoming file according the flags
//Flags are represernted as a bit mask
func Parse(file *os.File, flags int) *ParsingResult {
	scaner := bufio.NewScanner(file)
	result := ParsingResult{UniqueWords: data.NewWordSet()}

	for scaner.Scan() {
		if (flags & 2) > 0 {
			result.LinesCounter++
		}
		if (flags & 13) > 0 { //Flag mask 1101
			line := scaner.Text()
			if (flags & 1) > 0 {
				result.SymbolsCounter += len(line)
			}
			if (flags & 12) > 0 { //Flag mask 1100
				getUnique := (flags & 8) > 0
				lineWords := scanLine(line, getUnique)
				if (flags & 4) > 0 {
					result.WordsCounter += len(lineWords)
				}
				if getUnique {
					result.UniqueWords.AddAll(lineWords)
				}
			}
		}
	}

	scannerErr := scaner.Err()
	if scannerErr != nil {
		fmt.Printf("Error parsing the file: %s", scannerErr)
	}

	return &result
}

func scanLine(line string, getUnique bool) []string {
	lineWords := strings.Fields(line)

	if !getUnique {
		return lineWords
	}

	result := make([]string, 0, len(lineWords))

	regExp, regExpErr := regexp.Compile("[[:punct:]“”…]") //Regexp to detect punctuation and non standart quotation marks
	if regExpErr != nil {
		fmt.Printf("Regular expression creation failed: %s", regExpErr)
		return nil
	}

	for _, word := range lineWords {
		normalizedWord := strings.ToLower(word) //Changing to lower case and removing puncuation
		normalizedWord = regExp.ReplaceAllString(normalizedWord, "")
		if len(normalizedWord) != 0 { //Skiping empty words
			result = append(result, normalizedWord)
		}
	}

	return result
}
