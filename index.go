package fts

// InvIndex implemens an in-memory inverted index.
type InvIndex struct {
	docs  map[int64]Document
	index map[string][]int64
}

func NewInvIndex() InvIndex {
	invIndex := InvIndex{
		docs:  make(map[int64]Document),
		index: make(map[string][]int64),
	}
	return invIndex
}

// Add does index construction on a set of documents
func (index *InvIndex) Add(docs []Document) {
	// Index construction (naive)
	for _, doc := range docs {
		// TODO
		// doc.Text must be tokenized to extract singular words
		// the inverted index maps words to posting lists
		for _, token := range Preprocess(doc.content) {
			ids := index.index[token]
			if ids != nil && ids[len(ids)-1] == doc.id {
				// Don't add same ID twice.
				continue
			}
			index.index[token] = append(index.index[token], doc.id)
		}

	}
}

// Get returns the postings lists for a term.
func (index *InvIndex) Get(term string) []int64 {
	return index.index[term]
}

// Search runs returns a list of all postings lists with the term.
func (index *InvIndex) Search(text string) [][]int64 {
	var r [][]int64
	for _, token := range Preprocess(text) {
		if ids, ok := index.index[token]; ok {
			r = append(r, ids)
		}
	}
	return r
}
