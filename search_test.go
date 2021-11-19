package fts

import "testing"

func TestSearch(t *testing.T) {
	t.Run("TestNaiveSearch", func(t *testing.T) {
		path := "corpus/enwiki-snapshot.xml"
		docs, err := loadXMLDocuments(path)
		if err != nil {
			t.Fatal("test failed with error : ", err)
		}
		if len(docs) == 0 {
			t.Fatal("test failed got 0 docs")
		}
		collection := NewCollection(docs)
		idMatches := NaiveSearch(collection, "Chulek")
		if len(idMatches) != 1 {
			t.Fatal("test failed expected 1 match got :", len(idMatches))
		}

	})
}
