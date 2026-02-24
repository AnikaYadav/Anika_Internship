package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TrieNode struct {
	children map[rune]*TrieNode
	isLeaf   bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
	}
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

func (t *Trie) Insert(word string) {
	curr := t.root
	for _, ch := range word {
		if curr.children[ch] == nil {
			curr.children[ch] = NewTrieNode()
		}
		curr = curr.children[ch]
	}
	curr.isLeaf = true
}

func (t *Trie) Search(word string) bool {
	curr := t.root
	for _, ch := range word {
		if curr.children[ch] == nil {
			return false
		}
		curr = curr.children[ch]
	}
	return curr.isLeaf
}

func deleteHelper(node *TrieNode, word []rune, depth int) bool {
	if node == nil {
		return false
	}

	if depth == len(word) {
		if !node.isLeaf {
			return false
		}
		node.isLeaf = false
		return true
	}

	ch := word[depth]
	child := node.children[ch]
	if child == nil {
		return false
	}

	deleted := deleteHelper(child, word, depth+1)

	if deleted && len(child.children) == 0 && !child.isLeaf {
		delete(node.children, ch)
	}

	return deleted
}

func (t *Trie) Delete(word string) bool {
	runes := []rune(word)
	return deleteHelper(t.root, runes, 0)
}

func (t *Trie) dfs(node *TrieNode, prefix []rune, result *[]string) {
	if node == nil {
		return
	}

	if node.isLeaf {
		*result = append(*result, string(prefix))
	}

	for ch, child := range node.children {
		prefix = append(prefix, ch)
		t.dfs(child, prefix, result)
		prefix = prefix[:len(prefix)-1]
	}
}

func (t *Trie) ListAll() []string {
	var result []string
	t.dfs(t.root, []rune{}, &result)
	return result
}

func readInput(reader *bufio.Reader) (string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func main() {
	trie := NewTrie()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Command (add, list, remove, check, quit):")
		fmt.Print("> ")

		command, err := readInput(reader)
		if err != nil {
			fmt.Println("Error reading command:", err)
			continue
		}
		command = strings.ToLower(command)

		switch command {

		case "add":
			fmt.Print("Input: ")
			input, err := readInput(reader)
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue
			}
			trie.Insert(input)
			fmt.Println("Added!")

		case "check":
			fmt.Print("Input: ")
			input, err := readInput(reader)
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue
			}
			if trie.Search(input) {
				fmt.Println("Exists!")
			} else {
				fmt.Println("Does not exist.")
			}

		case "remove":
			fmt.Print("Input: ")
			input, err := readInput(reader)
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue
			}
			if trie.Delete(input) {
				fmt.Println("Removed!")
			} else {
				fmt.Println("Item not found.")
			}

		case "list":
			list := trie.ListAll()
			if len(list) == 0 {
				fmt.Println("No items found.")
				continue
			}
			fmt.Println("Items:")
			for _, item := range list {
				fmt.Println("-", item)
			}

		case "quit":
			return

		default:
			fmt.Println("Invalid command.")
		}
	}
}
