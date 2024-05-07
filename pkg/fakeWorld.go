package pkg

import (
	"github.com/nwtgck/go-fakelish"
	"strings"
)

func GenerateFakeWord() string {
	minLength := 8
	maxLength := 16
	fakeWord := fakelish.GenerateFakeWord(minLength, maxLength)
	fakeWord = strings.Title(fakeWord)
	return fakeWord
}
