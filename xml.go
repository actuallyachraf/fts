package fts

import (
	"encoding/xml"
	"os"
)

// XMLDocument is for representing in memory xml docs.
type XMLDocument struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int64
}

// XMLDocuments represents a set of documents.
type XMLDocuments struct {
	Documents []XMLDocument `xml:"doc"`
}

func loadXMLDocuments(path string) ([]XMLDocument, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dec := xml.NewDecoder(f)
	dump := struct {
		Documents []XMLDocument `xml:"doc"`
	}{}
	if err := dec.Decode(&dump); err != nil {
		return nil, err
	}

	docs := dump.Documents
	for i := range docs {
		docs[i].ID = int64(i)
	}

	return docs, nil
}
