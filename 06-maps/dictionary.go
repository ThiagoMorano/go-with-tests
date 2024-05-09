package main

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	if err == nil {
		// expects word to not exist
		return ErrWordExists
	} else if err != ErrNotFound {
		// a different error was returned
		return err
	}

	d[word] = definition
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	if err == ErrNotFound {
		// word was not found
		return ErrWordDoesNotExist
	} else if err != nil {
		// a different error was returned
		return err
	}

	d[word] = definition
	return nil
}
