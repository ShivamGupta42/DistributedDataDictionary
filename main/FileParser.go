package main

import (
	"encoding/json"
	os "os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ParseDictionaryFile(fName string) map[string]string {
	data, err := os.ReadFile(fName)
	check(err)
	dict := make(map[string]string, 0)
	err = json.Unmarshal(data, &dict)
	check(err)
	return dict
}

func TransformDictionaryMap(m map[string]string) map[string]string {
	for key := range m {
		m[key] = strings.Repeat(key, 3)
	}
	return m
}

func WriteMapToFile(m map[string]string) {
	f, err := os.Create("./words_dictionary_transformed.json")
	check(err)
	defer f.Close()

	jsonStr, err := json.Marshal(m)
	check(err)
	f.Write(jsonStr)
}
