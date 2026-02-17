package main

import (
	"testing"
)

func TestInsertAndSearch(t *testing.T) {
	trie := NewTrie()
	wordsToInsert := []string{"apple", "app", "banana"}

	for _, word := range wordsToInsert {
		trie.Insert(word)
	}

	testCases := []struct {
		name     string
		word     string
		expected bool
	}{
		{"ExistingWord", "apple", true},
		{"ExistingWord", "app", true},
		{"NonExistentWord", "orange", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := trie.Search(tc.word)
			if result != tc.expected {
				t.Errorf("%s: got %v, expected %v", tc.name, result, tc.expected)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	trie := NewTrie()
	wordsToInsert := []string{"cat", "car", "dog"}

	for _, word := range wordsToInsert {
		trie.Insert(word)
	}

	wordsToRemove := []string{"card", "cat"}

	deleteTests := []struct {
		name     string
		word     string
		expected bool
	}{
		{"Deleted", "card", false},
		{"Deleted", "cat", false},
		{"StillExists", "car", true},
	}

	for _, word := range wordsToRemove {
		trie.Delete(word)
	}

	for _, tc := range deleteTests {
		t.Run(tc.name, func(t *testing.T) {
			result := trie.Search(tc.word)
			if result != tc.expected {
				t.Errorf("%s: got %v, expected %v", tc.name, result, tc.expected)
			}
		})
	}
}

func TestListAll(t *testing.T) {
	trie := NewTrie()
	wordsToInsert := []string{"hello", "world", "test", "trie"}

	for _, word := range wordsToInsert {
		trie.Insert(word)
	}

	result := trie.ListAll()

	if len(result) != len(wordsToInsert) {
		t.Errorf("ListAll length: got %d, expected %d", len(result), len(wordsToInsert))
	}

	for i, word := range result {
		if word != wordsToInsert[i] {
			t.Errorf("ListAll at index %d: got %v, expected %v", i, word, wordsToInsert[i])
		}
	}
}
