package fts

import "testing"

func TestXML(t *testing.T) {
	t.Run("TestLoadDocuments", func(t *testing.T) {
		path := "corpus/enwiki-latest-abstract10.xml"
		docs, err := loadXMLDocuments(path)
		if err != nil {
			t.Fatal("test failed with error : ", err)
		}
		for _, doc := range docs {
			t.Log(doc.Title)
		}
		t.Log(len(docs))

	})
}
