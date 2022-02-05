package main

func main() {
	/* Transforming file */
	dictMap := ParseDictionaryFile("")
	tranMap := TransformDictionaryMap(dictMap)
	WriteMapToFile(tranMap)
}
