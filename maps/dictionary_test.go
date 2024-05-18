package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		d := Dictionary{}
		word := "test"
		def := "this is just a test"

		d.Add(word, def)

		assertDefinition(t, d, word, def)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		d := Dictionary{word: def}
		err := d.Add(word, "replaced")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, d, word, def)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "def"

		d := Dictionary{word: def}
		newDef := "def updated"
		d.Update(word, newDef)

		assertDefinition(t, d, word, newDef)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"

		d := Dictionary{}
		newDef := "def updated"
		err := d.Update(word, newDef)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	d := Dictionary{word: "def"}

	d.Delete(word)

	_, err := d.Search(word)
	if err != ErrNotFound {
		t.Errorf("expected %q to be deleted", word)
	}
}

func assertDefinition(t testing.TB, d Dictionary, word, def string) {
	t.Helper()

	got, err := d.Search(word)
	if err != nil {
		t.Fatal("should find added word: ", err)
	}

	assertStrings(t, got, def)
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
