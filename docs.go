package fts

// DocType represents document types could be JSON, XML or plain text.
type DocType int

const (
	Plain DocType = iota
	XML
	JSON
)

// Document represents a document in a corpus
type Document struct {
	t       DocType
	content string
	length  int
	id      int64
}

// NewDocument creates a new instance of Document.
func NewDocument(doc XMLDocument) Document {
	return Document{
		t:       XML,
		content: doc.Text,
		length:  len(doc.Text),
		id:      doc.ID,
	}
}

// NewCollection creates a collection of documents from a base
// document representation.
func NewCollection(docs []XMLDocument) []Document {
	documents := make([]Document, 0, len(docs))
	for _, doc := range docs {
		documents = append(documents, NewDocument(doc))
	}
	return documents
}
