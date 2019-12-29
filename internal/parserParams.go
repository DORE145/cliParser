package internal

type parserParams struct {
	symbols     bool
	lines       bool
	words       bool
	uniqueWords bool
}

//NewParams function constructs new parser parameters structure
func NewParams(symbols bool, lines bool, words bool, uniqueWords bool) parserParams {
	return parserParams{symbols, lines, words, uniqueWords}
}

func (params parserParams) isSymbols() bool {
	return params.symbols
}

func (params parserParams) isLines() bool {
	return params.lines
}

func (params parserParams) isWords() bool {
	return params.words
}

func (params parserParams) isUniqueWords() bool {
	return params.uniqueWords
}
