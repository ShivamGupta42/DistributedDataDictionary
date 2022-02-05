package main

import "testing"

func acceptAndValidate(t *testing.T, got interface{}, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("%s -> Got %s, Wanted %s", t.Name(), got, want)
	}
}

func TestParseDictionaryFile(t *testing.T) {
	fName := "./words_dictionary.json"
	resMap := ParseDictionaryFile(fName)

	t.Run("Dictionary is fully loaded ", func(t *testing.T) {
		acceptAndValidate(t, len(resMap), 370101)
	})

	t.Run("Dictionary contains a random word ", func(t *testing.T) {
		_, ok := resMap["insuitable"]
		if ok != true {
			t.Fail()
		}
	})
}

func TestTransformDictionaryMap(t *testing.T) {
	resMap := TransformDictionaryMap(map[string]string{"a": ""})

	if resMap["a"] != "aaa" {
		t.Fail()
	}
}
