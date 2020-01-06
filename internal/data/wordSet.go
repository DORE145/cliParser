package data

import "fmt"

//WordSet is a collection that stores unique words
//Based on binary tree
type WordSet struct {
	root *node
}

type node struct {
	value string
	left  *node
	right *node
}

//NewWordSet function initializes empty set
func NewWordSet() WordSet {
	return WordSet{}
}

//Add function adds word to the collection. If word is already present, nothing will happen
func (set *WordSet) Add(word string) {
	n := &node{word, nil, nil}
	if set.root == nil {
		set.root = n
	} else {
		add(set.root, n)
	}
}

//Internal recursive add function
func add(parrentNode, newNode *node) {
	if newNode.value < parrentNode.value {
		if parrentNode.left == nil {
			parrentNode.left = newNode
		} else {
			add(parrentNode.left, newNode)
		}
	}
	if newNode.value > parrentNode.value {
		if parrentNode.right == nil {
			parrentNode.right = newNode
		} else {
			add(parrentNode.right, newNode)
		}
	}
	// Case when words are equal intentianally ignored
}

//AddAll function adds all words from a slice to the collection. If word is already present, nothing will happen
func (set *WordSet) AddAll(words []string) {
	for _, word := range words {
		set.Add(word)
	}
}

//Remove function deletes word from the collection
func (set *WordSet) Remove(word string) {
	set.root = remove(set.root, word)
}

//Internal recursive delete function
func remove(parrentNode *node, value string) *node {
	if parrentNode == nil {
		return nil
	}

	if value < parrentNode.value {
		parrentNode.left = remove(parrentNode.left, value)
		return parrentNode
	}

	if value > parrentNode.value {
		parrentNode.right = remove(parrentNode.right, value)
		return parrentNode
	}

	if parrentNode.left == nil && parrentNode.right == nil {
		parrentNode = nil
		return nil
	}

	if parrentNode.left == nil {
		parrentNode = parrentNode.right
		return parrentNode
	}

	if parrentNode.right == nil {
		parrentNode = parrentNode.left
		return parrentNode
	}

	//Findong smallest right child
	smallestChild := parrentNode.right
	for {
		if smallestChild != nil && smallestChild.left != nil {
			smallestChild = smallestChild.left
		} else {
			break
		}
	}
	parrentNode.value = smallestChild.value
	parrentNode.right = remove(parrentNode.right, smallestChild.value)
	return parrentNode
}

//Size function returns the size of the internal collection
func (set *WordSet) Size() int {
	return size(set.root)
}

//Internal recursive size function
func size(root *node) int {
	if root == nil {
		return 0
	}
	return size(root.left) + size(root.right) + 1
}

//Contains function check is the provided word is present in the collection
func (set *WordSet) Contains(word string) bool {
	return contains(set.root, word)
}

//Internal recursive contains function
func contains(root *node, word string) bool {
	if root == nil {
		return false
	}
	if word < root.value {
		return contains(root.left, word)
	}
	if word > root.value {
		return contains(root.right, word)
	}
	return true
}

//Words function returns all present words in the set
func (set *WordSet) Words() []string {
	container := make([]string, 0)
	container = words(set.root, container)
	return container
}

//Internal recursive words function
func words(root *node, container []string) []string {
	if root != nil {
		fmt.Println(root.value)
		container = words(root.left, container)
		container = append(container, root.value)
		container = words(root.right, container)
	}
	return container
}
