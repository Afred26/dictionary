package dict

import "dictionary/entry"

type Dict struct {
	entries []entry.Entry
}

// New erzeugt ein neues leeres Dict-Objekt.
func New() Dict {
	return Dict{}
}

// Add fügt einen Eintrag zum Wörterbuch hinzu.
func (d *Dict) Add(e entry.Entry) {
	// Hinweis: Fügen Sie den Eintrag e mit append() zu d.entries hinzu.
	d.entries = append(d.entries, e)
}

// Size gibt die Anzahl der Einträge im Wörterbuch zurück.
func (d Dict) Size() int {
	// Hinweis: Geben Sie die Anzahl der Einträge von d.entries zurück.
	return len(d.entries)
}

// GetDe gibt den Eintrag mit dem deutschen Wort de zurück.
// Wenn kein Eintrag gefunden wird, wird ein leerer Eintrag zurückgegeben.
func (d Dict) GetDe(de string) entry.Entry {
	// Hinweis: Iterieren Sie über alle Einträge in d.entries
	// und geben Sie den Eintrag zurück, der das deutsche Wort de enthält.
	for _, e := range d.entries {
		if e.De() == de {
			return e
		}
	}
	return entry.Empty()
}

// Lookup sucht nach dem ersten Eintrag mit dem deutschen Wort de.
// Wenn ein Eintrag gefunden wird, wird der entsprechende englische string geliefert.
// Wenn kein Eintrag gefunden wird, wird ein leerer string zurückgegeben.
func (d Dict) Lookup(de string) string {
	// Hinweis: Verwenden Sie die Methode GetDe().
	return d.GetDe(de).En()
}

// GetAllDe gibt alle Einträge zurück, die das deutsche Wort de enthalten.
func (d Dict) GetAllDe(de string) []entry.Entry {
	// Hinweis: Iterieren Sie über alle Einträge in d.entries
	// und fügen Sie alle Einträge, die das deutsche Wort de enthalten, zu result hinzu.
	var result []entry.Entry
	for _, e := range d.entries {
		if e.De() == de {
			result = append(result, e)
		}
	}
	return result
}

// LookupAll sucht nach allen Einträgen mit dem deutschen Wort de.
// Gibt eine Liste mit den entsprechenden englischen Wörtern zurück.
func (d Dict) LookupAll(de string) []string {
	// Hinweis: Verwenden Sie die Methode GetAllDe().
	// Iterieren Sie über die Ergebnisse und fügen Sie die englischen Wörter zu result hinzu.
	var result []string
	for _, e := range d.GetAllDe(de) {
		result = append(result, e.En())
	}
	return result
}
