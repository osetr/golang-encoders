package main

import "strings"

var EncoderDictionary = map[string]string{
	"a": "1",
	"e": "2",
	"i": "3",
	"o": "4",
	"u": "5",
}

var DecoderDictionary = map[string]string{
	"1": "a",
	"2": "e",
	"3": "i",
	"4": "o",
	"5": "u",
}

func TranslateCodec(word string, Dictionary map[string]string) string { // please, don't use upper-camel-case for parameters, as they are local for a func and thus can't be public
	for key, value := range Dictionary {
		word = strings.ReplaceAll(word, key, value)
		word = strings.ReplaceAll(word, strings.ToUpper(key), value) // it's more efficient to define upper case letters in the Dict as well to avoid additinal operations
	}
	return word
}
