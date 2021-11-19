package fts

import "testing"

func TestSearch(t *testing.T) {
	t.Run("TestNaiveSearch", func(t *testing.T) {
		path := "corpus/enwiki-latest-abstract10.xml"
		docs, err := loadXMLDocuments(path)
		if err != nil {
			t.Fatal("test failed with error : ", err)
		}
		if len(docs) == 0 {
			t.Fatal("test failed got 0 docs")
		}
		collection := NewCollection(docs)
		idMatches := NaiveSearch(collection, "cat")
		t.Log(idMatches)

	})
}
