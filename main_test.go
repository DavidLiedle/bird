package main

import (
	"strings"
	"testing"
)

// TestBehavior300 runs the selector 300 times and verifies every category
// is exercised, no panics occur, and no output is empty. With three
// categories chosen uniformly, the probability any one is missed in 300
// runs is effectively zero ((2/3)^300).
func TestBehavior300(t *testing.T) {
	counts := map[string]int{"bird": 0, "sentence": 0, "fact": 0}
	for i := 0; i < 300; i++ {
		category, output := behavior()
		if _, ok := counts[category]; !ok {
			t.Fatalf("iteration %d: unknown category %q", i, category)
		}
		if strings.TrimSpace(output) == "" {
			t.Fatalf("iteration %d (category %s): empty output", i, category)
		}
		counts[category]++
	}
	for cat, n := range counts {
		if n == 0 {
			t.Errorf("category %q never exercised across 300 runs", cat)
		}
	}
}

// TestSourcesLoaded verifies the embedded files were parsed into non-empty
// slices. Regression guard against an accidental //go:embed rename or an
// overzealous splitLines.
func TestSourcesLoaded(t *testing.T) {
	if len(sentences) == 0 {
		t.Error("sentences slice is empty; check //go:embed and sentences.txt")
	}
	if len(facts) == 0 {
		t.Error("facts slice is empty; check //go:embed and facts.txt")
	}
	if len(birds) == 0 {
		t.Error("birds slice is empty; check birds.go")
	}
}
