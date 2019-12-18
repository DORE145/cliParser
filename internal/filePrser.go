package internal

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/DORE145/cliParser/internal/data"
)

func Parse(file *os.File, flags int) {
	scaner := bufio.NewScanner(file)
	var linesCounter, wordsCounter, symbolsCounter, printableSymbolsCounter int
	uniqueWords := data.NewWordSet()
	regExp, regExpErr := regexp.Compile("[[:punct:]]") //Regexp to detect punctuation

	if regExpErr != nil {
		fmt.Printf("Regular expression creation failed: %s", regExpErr)
		return
	}

	for scaner.Scan() {
		linesCounter++
		if (flags & 13) > 0 { // flag mask 1101
			line := scaner.Text()
			symbolsCounter += len(line)

			lineWords := strings.Fields(line)
			wordsCounter += len(lineWords)
			for _, word := range lineWords {
				symbolsCounter += len(word)
				normalizedWord := strings.ToLower(word) //Changing to lower case and removing puncuation
				normalizedWord = regExp.ReplaceAllString(normalizedWord, "")

				printableSymbolsCounter += len(normalizedWord)
				uniqueWords.Add(normalizedWord)
			}
		}
	}

	//TODO: Results printer
	//TODO: Split parse function
	//TODO: Write tests

	scannerErr := scaner.Err()
	if scannerErr != nil {
		fmt.Printf("Error parsing the file: %s", scannerErr)
	}
}
