package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("word not found")
var ErrKeyPresent = errors.New("key already present")

func (d Dictionary) Search(word string) (string, error) {
	w, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return w, nil
}

func (d Dictionary) addWord(word, definition string) error {
	_, nok := d[word]
	if nok {
		return ErrKeyPresent
	}
	d[word] = definition
	return nil
}
