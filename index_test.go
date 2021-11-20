package fts

import (
	"testing"
)

func TestInvertedIndex(t *testing.T) {
	t.Run("TestIndexConstruction", func(t *testing.T) {
		idx := NewInvIndex()
		idx.Add([]Document{{id: 1, content: "A donut on a glass plate. Only the donuts."}})
		idx.Add([]Document{{id: 2, content: "donut is a donut"}})
		expectedIndex := InvIndex{
			docs: make(map[int64]Document),
			index: map[string][]int64{
				"donut": {1, 2}, "glass": {1}, "onli": {1}, "plate": {1},
			},
		}

		for k, v := range idx.index {
			expected := expectedIndex.Get(k)
			got := v
			for idx, id := range expected {
				if got[idx] != id {
					t.Fatalf("failed to assert index consistency expected %d got %d", got[idx], v)
				}
			}
		}
	})
	t.Run("TestIndexSearch", func(t *testing.T) {
		path := "corpus/enwiki-snapshot.xml"
		docs, err := loadXMLDocuments(path)
		if err != nil {
			t.Fatal("test failed with error : ", err)
		}
		if len(docs) == 0 {
			t.Fatal("test failed got 0 docs")
		}
		collection := NewCollection(docs)
		index := NewInvIndex()
		index.Add(collection)

		idMatches := index.Search("Chulek")
		t.Log(idMatches)
		if len(idMatches) != 1 {
			t.Fatal("test failed expected 1 match got :", len(idMatches))
		}
	})
}

func BenchmarkIndexSearch(b *testing.B) {
	b.Run("BenchmarkIndexSearch", func(b *testing.B) {
		path := "corpus/enwiki-latest-abstract10.xml"
		docs, err := loadXMLDocuments(path)
		if err != nil {
			b.Fatal("test failed with error : ", err)
		}
		if len(docs) == 0 {
			b.Fatal("test failed got 0 docs")
		}
		collection := NewCollection(docs)
		index := NewInvIndex()
		index.Add(collection)

		idMatches := index.Search("Small wild cat")
		b.Log(idMatches)

	})
}
