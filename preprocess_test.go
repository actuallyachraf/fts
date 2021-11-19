package fts

import "testing"

func TestPreprocess(t *testing.T) {
	t.Run("TestTokenize", func(t *testing.T) {
		text := "A donut on a glass plate. Only the donuts."
		tokens := Tokenize(text)
		expected := []string{"A", "donut", "on", "a", "glass", "plate", "Only", "the", "donuts"}
		if len(tokens) == 0 {
			t.Fatal("test failed got 0 tokens")
		}
		for idx, token := range tokens {
			if expected[idx] != token {
				t.Fatalf("expected : %s at index %d got : %s", expected[idx], idx, token)
			}
		}
	})
	t.Run("TestNormalize", func(t *testing.T) {
		text := "A donut on a glass plate. Only the donuts."
		tokens := Normalize(Tokenize(text))
		expected := []string{"donut", "glass", "plate", "only", "donuts"}
		if len(tokens) == 0 {
			t.Fatal("test failed got 0 tokens")
		}
		for idx, token := range tokens {
			if expected[idx] != token {
				t.Fatalf("expected : %s at index %d got : %s", expected[idx], idx, token)
			}
		}
	})
}
