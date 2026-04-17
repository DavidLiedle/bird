// bird — a command-line tool for tired brains.
//
// Type `bird`, receive something. No flags, no args, no config, no network.
// Whatever you type after `bird` is ignored. The tired brain is allowed to
// type nonsense.
package main

import (
	"crypto/rand"
	_ "embed"
	"fmt"
	"math/big"
	"strings"
)

//go:embed sentences.txt
var sentencesRaw string

//go:embed facts.txt
var factsRaw string

var (
	sentences = splitLines(sentencesRaw)
	facts     = splitLines(factsRaw)
)

// splitLines turns an embedded file into a slice of non-empty lines.
func splitLines(s string) []string {
	out := []string{}
	for _, line := range strings.Split(s, "\n") {
		line = strings.TrimRight(line, "\r ")
		if line != "" {
			out = append(out, line)
		}
	}
	return out
}

// pick returns a uniformly random element of xs using crypto/rand.
// Seeding from crypto/rand avoids same-second repetition on rapid invocations.
func pick(xs []string) string {
	if len(xs) == 0 {
		return ""
	}
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(xs))))
	if err != nil {
		return xs[0]
	}
	return xs[n.Int64()]
}

// behavior chooses one of three behaviors uniformly at random and returns
// both a category tag (for tests) and the output to print.
func behavior() (category, output string) {
	n, err := rand.Int(rand.Reader, big.NewInt(3))
	if err != nil {
		return "sentence", pick(sentences)
	}
	switch n.Int64() {
	case 0:
		return "bird", pick(birds)
	case 1:
		return "sentence", pick(sentences)
	default:
		return "fact", pick(facts)
	}
}

func main() {
	_, output := behavior()
	fmt.Println(output)
}
