package maps

// never initialize an empty map variable
/*
*
var dictionary = map[string]string{}

// OR

var dictionary = make(map[string]string)
*
*/
type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("word already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	def, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return def, nil
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = def
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = def
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
