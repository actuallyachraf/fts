package fts

import (
	"strings"
	"unicode"
)

// token preprocessing
// Reuters RCV! stopwords
var stopwords = map[string]struct{}{
	"a":    {},
	"has":  {},
	"to":   {},
	"an":   {},
	"he":   {},
	"was":  {},
	"and":  {},
	"in":   {},
	"were": {},
	"are":  {},
	"is":   {},
	"will": {},
	"as":   {},
	"on":   {},
	"it":   {},
	"with": {},
	"the":  {},
}

// Map on strings.
func Map(input []string, f func(s string) string) []string {
	r := make([]string, len(input))
	for i, token := range input {
		r[i] = f(token)
	}
	return r
}

// Filter on strings.
func Filter(input []string, f func(s string) bool) []string {
	r := make([]string, 0, len(input))
	for _, token := range input {
		if !f(token) {
			r = append(r, token)
		}
	}
	return r
}

// IsStopWord checks if a word is a stop word,
func IsStopWord(s string) bool {
	_, ok := stopwords[strings.ToLower(s)]
	return ok
}

// Tokenize a character sequence using whitespace.
func Tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

// Normalize a list of tokens
func Normalize(tokens []string) []string {
	lower := Map(tokens, strings.ToLower)
	stops := Filter(lower, IsStopWord)
	stemmed := Map(stops, Stem)
	return stemmed
}

// Preprocess tokenize and normalize tokens.
func Preprocess(text string) []string {
	tokens := Tokenize(text)
	normalized := Normalize(tokens)
	return normalized
}
