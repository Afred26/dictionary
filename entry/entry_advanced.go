package entry

import "strings"

// FromWordPairCsv erzeugt ein neues Entry-Objekt aus einem String,
// der ein Wortpaar enthält. Das deutsche und das englische Wort sind
// durch ein Komma getrennt.
// Gibt es keines oder mehrere Kommas im String, wird ein leerer Eintrag zurückgegeben.
func FromWordPairCsv(pair string) Entry {
	var word Entry
	word.de = strings.Split(pair, ",")[0]
	word.en = strings.Split(pair, ",")[1]

	return word
}
