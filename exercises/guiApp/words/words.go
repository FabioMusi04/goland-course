package words

import (
	"math/rand"
)

var wordsList = []string{
	"apple",
	"banana",
	"cherry",
	"date",
	"elderberry",
	"fig",
}

func GetWords() []string {
	return wordsList
}

func GetRandomWord() string {
	return wordsList[rand.Intn(len(wordsList))]
}

func GetRandomWords(n int) []string {
	var words []string
	for i := 0; i < n; i++ {
		words = append(words, GetRandomWord())
	}
	return words
}

func GetWord(i int) string {
	return wordsList[i]
}

func MixLetters(s string) string {
	letters := []rune(s)
	rand.Shuffle(len(letters), func(i, j int) {
		letters[i], letters[j] = letters[j], letters[i]
	})
	return string(letters)
}
