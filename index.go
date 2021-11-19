package fts

// InvIndex implemens an in-memory inverted index.
type InvIndex struct {
	docs  map[int64]XMLDocument
	index map[string][]int64
}

func NewInvIndex(docs []XMLDocument) InvIndex {
	invIndex := InvIndex{
		docs:  make(map[int64]XMLDocument),
		index: make(map[string][]int64),
	}
	// Index construction (naive)
	for _, doc := range docs {
		// TODO
		// doc.Text must be tokenized to extract singular words
		// the inverted index maps words to posting lists
		invIndex.docs[doc.ID] = doc
		invIndex.index[doc.Text] = append(invIndex.index[doc.Text], doc.ID)
	}
	return invIndex
}
