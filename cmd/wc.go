package main

import (
	"flag"
	"fmt"
)

func main() {
	var symbols, lines, words, uniqueWords bool
	flag.BoolVar(&symbols, "symbols", false, "Count the number of sybols in the text file")
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
}

func printUsage() {
	fmt.Println("Wc is a tool for basic analysis of text files\n\n" +
		"Usage:\n\t wc [flags] path/to/file \n\n" +
		"Avaliable flags:\n" +
		"--help\n\tPrints usage information of the application\n" +
		"--lines\n\tCounts the number of lines in the text file\n" +
		"--symbols\n\tCount the number of sybols in the text file\n" +
		"--uniqueWords\n\tPrints out all unique words from the text file\n" +
		"--words\n\tCounts the number of words in the text file\n")
}
