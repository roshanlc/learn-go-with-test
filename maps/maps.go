package maps

type Dictionary map[string]string

func (d Dictionary) SearchDictionary(word string) string {
	return d[word]
}
