package main

import (
	"fmt"
)

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) populateSuffixTrie(str string) {
	for i := 0; i < len(str); i++ {
		t.insertSubstring(i, str)
	}
}

func (t *Trie) insertSubstring(index int, str string) {
	node := t.root
	for i := index; i < len(str); i++ {
		letter := str[i]
		if node.children == nil {
			node.children = make(map[rune]*TrieNode)
		}
		if _, ok := node.children[rune(letter)]; !ok {
			newNode := &TrieNode{}
			node.children[rune(letter)] = newNode
		}
		node = node.children[rune(letter)]

	}
	node.isEnd = true
}

func (t *Trie) insertion(str string) {

	node := t.root
	for i := 0; i < len(str); i++ {
		letter := rune(str[i])
		if node.children == nil {
			node.children = map[rune]*TrieNode{}
		}
		if _, ok := node.children[letter]; !ok {
			newNode := &TrieNode{}
			node.children[letter] = newNode
		}
		node = node.children[letter]

	}
	node.isEnd = true
}

func (t *Trie) contains(str string) bool {

	node := t.root
	for i := 0; i < len(str); i++ {
		letter := str[i]
		if _, ok := node.children[rune(letter)]; !ok {

			return node.isEnd
		}

		node = node.children[rune(letter)]
	}
	return node.isEnd
}
func (t *Trie) Display() {
	t.displayHelper(t.root, []rune{})
}

func (t *Trie) displayHelper(node *TrieNode, prefix []rune) {
	if node.isEnd {
		fmt.Println(string(prefix))
	}
	for char, child := range node.children {
		t.displayHelper(child, append(prefix, char))
	}
}

func main() {
	trie := Trie{root: &TrieNode{}}
	trie.populateSuffixTrie("athunl")

	trie.Display()
	fmt.Println(trie.contains("athunlal"))
}
