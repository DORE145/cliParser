package data

type WordSet struct {
	words map[string]bool
}

func NewWordSet() WordSet {
	return WordSet{make(map[string]bool)}
}

func (set WordSet) Add(word string) {
	set.words[word] = true
}

func (set WordSet) Remove(word string) {
	set.words[word] = false
}

func (set WordSet) Size() int {
	return len(set.words)
}

func (set WordSet) Contains(word string) bool {
	return set.words[word]
}
