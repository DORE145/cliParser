package main

import (
	"flag"
	"fmt"
)

func main() {
	symbols := *flag.Bool("-symbols", false, "Count the number of sybols in the text file")
	lines := *flag.Bool("-lines", false, "Counts the number of lines in the text file")
	words := *flag.Bool("-words", false, "Counts the number of words in the text file")
	uniqueWords := *flag.Bool("-uniqueWords", false, "Prints out all unique words from the text file")
	help := *flag.Bool("-help", true, "Prints usage information of the application")
	flag.Parse()

	if help {
		printUsage()
		return
	}

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
	fmt.Println("Avaliable flags:\n" +
		"--help\n\tPrints usage information of the application (default true)\n" +
		"--lines\n\tCounts the number of lines in the text file\n" +
		"--symbols\n\tCount the number of sybols in the text file\n" +
		"--uniqueWords\n\tPrints out all unique words from the text file\n" +
		"--words\n\tCounts the number of words in the text file\n")
}
