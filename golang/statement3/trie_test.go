package main

import (
	"reflect"
	"sort"
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

	sort.Strings(result)
	sort.Strings(wordsToInsert)

	if !reflect.DeepEqual(result, wordsToInsert) {
		t.Errorf("ListAll mismatch: got %v, expected %v", result, wordsToInsert)
	}
}
