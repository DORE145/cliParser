package internal

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/DORE145/cliParser/internal/data"
)

func Test_scanLine(t *testing.T) {
	type args struct {
		line      string
		getUnique bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"LineScan1", args{"He ran out of money, so he had to stop playing poker.", false},
			[]string{"He", "ran", "out", "of", "money,", "so", "he", "had", "to", "stop", "playing", "poker."}},
		{"LineScan2", args{"I currently have 4 windows open up… and I don’t know why.", true},
			[]string{"i", "currently", "have", "4", "windows", "open", "up", "and", "i", "don’t", "know", "why"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scanLine(tt.args.line, tt.args.getUnique); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("scanLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	file1, err1 := os.Open("../testdata/textfile1")
	file2, err2 := os.Open("../testdata/textfile2")
	file3, err3 := os.Open("../testdata/textfile3")

	if err1 != nil || err2 != nil || err3 != nil {
		t.Error("Could not read test data")
		t.FailNow()
	}
	words := data.NewWordSet()
	words.AddAll([]string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"})
	type args struct {
		file  *os.File
		flags int
	}
	tests := []struct {
		name string
		args args
		want *ParsingResult
	}{
		{"Parse1", args{file1, 7}, &ParsingResult{LinesCounter: 3, WordsCounter: 128, SymbolsCounter: 657}},
		{"Parse2", args{file2, 5}, &ParsingResult{WordsCounter: 93, SymbolsCounter: 562}},
		{"Parse3", args{file3, 8}, &ParsingResult{UniqueWords: words}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.args.file, tt.args.flags); !checkEquality(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func checkEquality(got, want *ParsingResult) bool {
	if got.LinesCounter != want.LinesCounter {
		return false
	}
	if got.SymbolsCounter != want.SymbolsCounter {

		return false
	}
	if got.WordsCounter != want.WordsCounter {
		return false
	}
	fmt.Println(got.UniqueWords)
	fmt.Println(want.UniqueWords)
	if (got.UniqueWords.Size() == 0) && (want.UniqueWords.Size() == 0) {
		return true
	}
	return reflect.DeepEqual(got.UniqueWords, want.UniqueWords)
}
