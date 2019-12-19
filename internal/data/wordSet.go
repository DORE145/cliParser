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

func (set WordSet) AddAll(words []string) {
	for _, word := range words {
		set.Add(word)
	}
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

func (set WordSet) Words() []string {
	words := make([]string, len(set.words))
	i := 0
	for word := range set.words {
		words[i] = word
		i++
	}

	return words
}
