package main

import "strings"

const (
	PigLatinSuffix             string = "ay"
	FirstLetterExceptions      string = "aeiou"
	FirstLetterExceptionSuffix string = "d" + PigLatinSuffix
)

func TranslatePigLatin(sentence string) string {
	var PigLatinWords []string
	EnglishWords := strings.Split(sentence, " ")

	for _, word := range EnglishWords {
		first := word[0:1]
		if strings.Contains(FirstLetterExceptions, first) {
			PigLatinWords = append(
				PigLatinWords,
				word + FirstLetterExceptionSuffix,
			)
		} else {
			PigLatinWords = append(
				PigLatinWords,
				word[1:] + first + PigLatinSuffix,
			)
		}
	}
	return strings.Join(PigLatinWords, " ")
}