package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/DORE145/cliParser/internal"
)

func main() {
	var symbols, lines, words, uniqueWords bool
	flag.BoolVar(&symbols, "symbols", false, "Count the number of symbols in the text file")
	flag.BoolVar(&lines, "lines", false, "Counts the number of lines in the text file")
	flag.BoolVar(&words, "words", false, "Counts the number of words in the text file")
	flag.BoolVar(&uniqueWords, "uniqueWords", false, "Prints out all unique words from the text file")
	flag.Parse()

	args := flag.Args()
	if len(args) > 1 {
		fmt.Println("Too much arguments")
		printUsage()
		return
	} else if len(args) == 0 {
		fmt.Println("Text file name not found")
		printUsage()
		return
	}

	//Storing enabled flags as separate data structure
	flags := internal.NewParams(symbols, lines, words, uniqueWords)

	file, err := os.Open(args[0])
	if err != nil {
		fmt.Printf("Could not read the file: %s", err)
		return
	}
	result := internal.Parse(file, flags)

	file.Close()

	if lines {
		fmt.Printf("Number of lines in the file: %d\n", result.LinesCounter)
	}
	if words {
		fmt.Printf("Number of words in the file: %d\n", result.WordsCounter)
	}
	if symbols {
		fmt.Printf("Number of symbols in the file: %d\n", result.SymbolsCounter)
	}
	if uniqueWords {
		printUniqueWords(result.UniqueWords.Words())
	}
}

func printUsage() {
	fmt.Println("Wc is a tool for basic analysis of text files\n\n" +
		"Usage:\n\t wc [flags] path/to/file \n\n" +
		"Avaliable flags:\n" +
		"--help\n\tPrints usage information of the application\n" +
		"--lines\n\tCounts the number of lines in the text file\n" +
		"--symbols\n\tCount the number of symbols in the text file\n" +
		"--uniqueWords\n\tPrints out all unique words from the text file\n" +
		"--words\n\tCounts the number of words in the text file\n")
}

func printUniqueWords(words []string) {
	fmt.Printf("Number of unique words: %d\nList of the unique words:\n", len(words))
	for index, word := range words {
		fmt.Printf("%13s\t", word)
		if (index % 6) == 0 {
			fmt.Println()
		}
	}
}
