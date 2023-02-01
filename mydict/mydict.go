package mydict

import (
	"errors"
)

// Dictionary type
type Dictionary map[string]string //map[string]string은 dict다.

var (
	errNotFound   = errors.New("Not Found") //error 메세지를 미리 정의해서 사용할 수 있다.
	errCantUpdate = errors.New("Can't update non-existing word")
	errWordExists = errors.New("That word already exists")
) // 변수들이 같은 type이면 ()를 사용해서 한번에 type 지정을 할 수 있다.

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word] //Go에서 map에서 요소를 찾으면 2가지 값을 반환한다. value와 ok(값이 있는 여부, boolean)
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errWordExists
	}
	return nil
}

// Update a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
