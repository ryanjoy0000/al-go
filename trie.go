package main

import (
	"fmt"
)

type Trie struct {
	root *TrieNode
}

func NewTrie(root *TrieNode) *Trie {
	return &Trie{
		root: root,
	}
}

func (t *Trie) insert(word string) {
	currentNode := t.root
	for pos, char := range word {
		var next *TrieNode
		// check if char present
		if currentNode.isCharPresent(char) {
			next = currentNode.children[char]
		} else {
			// char not present, create new node
			next = NewTrieNode()
		}
		currentNode.insertChar(char, next)
		currentNode = next

		if pos == len(word)-1 {
			currentNode.isWord = true
		}
	}
}

func (t *Trie) show() {
	currentNode := t.root
	currentNode.show()
}

func (t *Trie) hasWord(word string) bool {
	fmt.Println("searching for: ", word)
	result := false
	currentNode := t.root
	for _, char := range word {
		if currentNode.isCharPresent(char) {
			fmt.Println(string(char), " is present")
			currentNode = currentNode.children[char]
			if currentNode.isWord {
				result = true
			}
		} else {
			result = false
			fmt.Println(string(char), " is not present")
			break
		}
	}
	return result
}

// ------------------------------------------------
type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isWord:   false,
	}
}

func (node *TrieNode) isCharPresent(r rune) bool {
	_, ok := node.children[r]
	return ok
}

func (node *TrieNode) insertChar(r rune, next *TrieNode) {
	node.children[r] = next
}

func (node *TrieNode) show() {
	if node == nil {
		fmt.Println("nil...")
		return
	}
	fmt.Println()
	for rune, next := range node.children {
		fmt.Print(string(rune), "_", rune, " -> ", next)
		next.show()
	}
}

// ------------------------------------------------
