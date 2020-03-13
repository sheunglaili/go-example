package maps

import "testing"

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("Known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("Unknow word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrWordNotFound

		assertErr(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dict := Dictionary{word: definition}
		err := dict.Add(word, definition)

		assertErr(t, err, ErrWordExists)
		assertDefinition(t, dict, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dict := Dictionary{word: definition}
		newDefinition := "updated"
		err := dict.Update(word, newDefinition)

		if err != nil {
			t.Fatal("not expecting error ")
		}

		assertDefinition(t, dict, word, newDefinition)
	})

	t.Run("update not exist word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dict := Dictionary{word: definition}
		notExistWord := "notExist"
		newDefinition := "updated"
		err := dict.Update(notExistWord, newDefinition)

		assertErr(t, err, ErrWordDoesNotExist)
		// assertDefinition(t, dict, word, newDefinition)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete", func(t *testing.T) {
		word := "test"
		dict := Dictionary{word: "test definition"}
		dict.Delete(word)
		_, err := dict.Search(word)
		if err != ErrWordNotFound {
			t.Errorf("Expected %q to be deleted", word)
		}
	})
}

func assertErr(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}
