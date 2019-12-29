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
func Parse(file *os.File, flags parserParams) *ParsingResult {
	scaner := bufio.NewScanner(file)
	result := ParsingResult{UniqueWords: data.NewWordSet()}

	for scaner.Scan() {
		if flags.isLines() {
			result.LinesCounter++
		}
		if flags.isSymbols() || flags.isUniqueWords() || flags.isWords() {
			line := scaner.Text()
			if flags.isSymbols() {
				result.SymbolsCounter += len(line)
			}
			if flags.isWords() || flags.isUniqueWords() { //Flag mask 1100
				lineWords := scanLine(line, flags.isUniqueWords())
				if flags.isWords() {
					result.WordsCounter += len(lineWords)
				}
				if flags.isUniqueWords() {
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
