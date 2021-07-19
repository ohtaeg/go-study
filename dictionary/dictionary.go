package dictionary

import "errors"

// Dictionary : type에 대한 alias
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not found")
	errCantUpdate = errors.New("Can't update non-existing word")
	errWordExists = errors.New("Word already exist")
)

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add : Don't need the `*` on the receiver because maps on Go are automatically using `*`
func (d Dictionary) Add(word string, definition string) error {
	_, error := d.Search(word)
	if error == errNotFound {
		d[word] = definition
		return nil
	}
	return errWordExists
}

func (d Dictionary) Update(word string, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
