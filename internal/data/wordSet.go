package data

//WordSet is a collection that stores unique words
//Based on hashmap and exploits default values of bool variables
type WordSet struct {
	words map[string]bool
}

//NewWordSet function initializes internal map and returns new WordSet
func NewWordSet() WordSet {
	return WordSet{make(map[string]bool)}
}

//Add funtion adds word to the collection. If word is already present, nothing will happen
func (set WordSet) Add(word string) {
	set.words[word] = true
}

//AddAll funtion adds all words from a slice to the collection. If word is already present, nothing will happen
func (set WordSet) AddAll(words []string) {
	for _, word := range words {
		set.Add(word)
	}
}

//Remove function marks word in the collection as it is not present in the collection.
//As a side effect it leaves a tombstone in the collection and not actualy removes it from map.
//As a second side efffect if you will try to remove word that is not present in the collection it will create a tombstone and will increase size of the collection
func (set WordSet) Remove(word string) {
	set.words[word] = false
}

//Size function returns the size of the internal collection including tombstones
func (set WordSet) Size() int {
	return len(set.words)
}

//Contains function check is the provided word is present in the collection
func (set WordSet) Contains(word string) bool {
	return set.words[word]
}

//Words function returns all present words in the set excluding tombstones
func (set WordSet) Words() []string {
	words := make([]string, len(set.words))
	i := 0
	for word, present := range set.words {
		if (present) {
			words[i] = word
			i++
		}
	}

	return words
}
