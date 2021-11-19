package fts

import "testing"

func TestXML(t *testing.T) {
	t.Run("TestLoadDocuments", func(t *testing.T) {
		path := "corpus/enwiki-snapshot.xml"
		docs, err := loadXMLDocuments(path)
		if err != nil {
			t.Fatal("test failed with error : ", err)
		}
		if len(docs) != 14 {
			t.Fatal("expected to read 14 files got ", len(docs))
		}

	})
}
