package trie_test

import (
	"dsa/ds/trie"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewTrie(t *testing.T) {
	tr := trie.NewTrie()
	if tr == nil {
		t.Error("NewTrie() returned nil")
	}
}

func TestTrieInsert(t *testing.T) {
	tr := trie.NewTrie()
	
	words := []string{"apple", "app", "application", "banana", "ball", "cat"}
	for _, word := range words {
		tr.Insert(word)
	}
	
	for _, word := range words {
		if !tr.Search(word) {
			t.Errorf("Word %s not found after insertion", word)
		}
	}
	
	notInserted := []string{"ap", "appl", "ban", "batman", "dog"}
	for _, word := range notInserted {
		if tr.Search(word) {
			t.Errorf("Word %s incorrectly found when it wasn't inserted", word)
		}
	}
}

func TestTrieSearch(t *testing.T) {
	tr := trie.NewTrie()
	
	if tr.Search("any") {
		t.Error("Search returned true on empty trie")
	}
	
	tr.Insert("hello")
	tr.Insert("world")
	
	testCases := []struct {
		word     string
		expected bool
	}{
		{"hello", true},
		{"world", true},
		{"hell", false},
		{"worl", false},
		{"worlds", false},
		{"", false},
	}
	
	for _, tc := range testCases {
		result := tr.Search(tc.word)
		if diff := cmp.Diff(tc.expected, result); diff != "" {
			t.Errorf("Search(%s) mismatch (-want +got):\n%s", tc.word, diff)
		}
	}
}

func TestTrieStartsWith(t *testing.T) {
	tr := trie.NewTrie()
	
	if tr.StartsWith("any") {
		t.Error("StartsWith returned true on empty trie")
	}
	
	tr.Insert("hello")
	tr.Insert("helicopter")
	tr.Insert("world")
	
	testCases := []struct {
		prefix   string
		expected bool
	}{
		{"hel", true},
		{"hello", true},
		{"helicopter", true},
		{"helium", false},
		{"wor", true},
		{"world", true},
		{"worlds", false},
		{"", true},
	}
	
	for _, tc := range testCases {
		result := tr.StartsWith(tc.prefix)
		if diff := cmp.Diff(tc.expected, result); diff != "" {
			t.Errorf("StartsWith(%s) mismatch (-want +got):\n%s", tc.prefix, diff)
		}
	}
}

func TestTrieGetAllWord(t *testing.T) {
	tr := trie.NewTrie()
	
	result := tr.GetAllWord()
	if len(result) != 0 {
		t.Errorf("GetAllWord on empty trie should return empty slice, got %v", result)
	}
	
	words := []string{"apple", "app", "application", "banana", "ball", "cat"}
	for _, word := range words {
		tr.Insert(word)
	}
	
	result = tr.GetAllWord()
	
	sort.Strings(result)
	sort.Strings(words)
	
	if diff := cmp.Diff(words, result); diff != "" {
		t.Errorf("GetAllWord mismatch (-want +got):\n%s", diff)
	}
}

func TestTrieDelete(t *testing.T) {
	tr := trie.NewTrie()
	
	if tr.Delete("any") {
		t.Error("Delete returned true on empty trie")
	}
	
	words := []string{"apple", "app", "application", "banana"}
	for _, word := range words {
		tr.Insert(word)
	}
	
	testCases := []struct {
		word           string
		expectedResult bool
		wordsAfter     []string
	}{
		{"appl", false, []string{"apple", "app", "application", "banana"}},
		{"apple", false, []string{"app", "application", "banana"}},
		{"app", false, []string{"application", "banana"}},
		{"banana", false, []string{"application"}},
		{"application", true, []string{}},
	}
	
	for i, tc := range testCases {
		result := tr.Delete(tc.word)
		if diff := cmp.Diff(tc.expectedResult, result); diff != "" {
			t.Errorf("Case %d: Delete(%s) result mismatch (-want +got):\n%s", i, tc.word, diff)
		}
		
		allWords := tr.GetAllWord()
		sort.Strings(allWords)
		sort.Strings(tc.wordsAfter)
		
		if diff := cmp.Diff(tc.wordsAfter, allWords); diff != "" {
			t.Errorf("Case %d: After Delete(%s), words mismatch (-want +got):\n%s", i, tc.word, diff)
		}
	}
}

func TestTrieWithUnicode(t *testing.T) {
	tr := trie.NewTrie()
	
	words := []string{"café", "résumé", "naïve", "piñata", "こんにちは", "你好", "안녕하세요"}
	for _, word := range words {
		tr.Insert(word)
	}
	
	for _, word := range words {
		if !tr.Search(word) {
			t.Errorf("Unicode word %s not found after insertion", word)
		}
	}
	
	prefixes := []struct {
		prefix   string
		expected bool
	}{
		{"caf", true},
		{"rés", true},
		{"こん", true},
		{"안녕", true},
		{"café", true},
		{"cafés", false},
	}
	
	for _, tc := range prefixes {
		result := tr.StartsWith(tc.prefix)
		if diff := cmp.Diff(tc.expected, result); diff != "" {
			t.Errorf("StartsWith(%s) with Unicode mismatch (-want +got):\n%s", tc.prefix, diff)
		}
	}
	
	result := tr.GetAllWord()
	sort.Strings(result)
	sort.Strings(words)
	
	if diff := cmp.Diff(words, result); diff != "" {
		t.Errorf("GetAllWord with Unicode mismatch (-want +got):\n%s", diff)
	}
}