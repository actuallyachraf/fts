package fts

import (
	"strings"
)

// NaiveSearch does naive search on documents.
func NaiveSearch(docs []Document, term string) []Document {
	r := make([]Document, 0)

	for _, doc := range docs {
		if strings.Contains(doc.content, term) {
			r = append(r, doc)
		}
	}
	return r
}
