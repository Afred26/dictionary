package dict

import (
	"dictionary/entry"
	"os"
	"strings"
)

// FromWordPairsCsv erzeugt ein neues Wörterbuch aus einer Liste von Wortpaaren.
// Ein Wortpaar ist dabei jeweils gegeben als ein String mit zwei Wörtern,
// getrennt durch ein Komma.
// Gibt es in einem Wortpaar kein oder mehrere Kommas, wird das Wortpaar ignoriert.
func FromWordPairsCsv(words []string) Dict {
	d := New()
	for _, w := range words {
		e := entry.FromWordPairCsv(w)
		if e.IsValid() {
			d.Add(e)
		}
	}
	return d
}

// ReadFileLines liest den Inhalt einer Datei und liefert die Zeilen als String-Slice.
// Wenn die Datei nicht gelesen werden kann, wird eine leere Liste zurückgegeben.
func ReadFileLines(filename string) []string {
	file, err := os.ReadFile(filename)
	if err != nil {
		return []string{}
	}
	return strings.Split(string(file), "\n")
}

// FromFile erzeugt ein neues Wörterbuch aus dem Inhalt einer Datei.
// Die Datei enthält eine Liste von Wortpaaren, wobei jedes Wortpaar in einer eigenen Zeile steht.
// Ein Wortpaar ist dabei jeweils gegeben als ein String mit zwei Wörtern,
// getrennt durch ein Komma.
// Gibt es in einem Wortpaar kein oder mehrere Kommas, wird das Wortpaar ignoriert.
// Wenn die Datei nicht gelesen werden kann, wird ein leeres Wörterbuch zurückgegeben.
func FromFile(filename string) Dict {
	//d := New()
	Word := ReadFileLines(filename)

	return FromWordPairsCsv(Word)
}
