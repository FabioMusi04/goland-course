package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Random() {
	arrayWords := []string{"Hello", "World", "Golang", "Programming", "Language", "Random", "Words"} //array of words
	rand.New(rand.NewSource(time.Now().UnixNano()))                                                  //random seed
	fmt.Println(rand.Intn(100))                                                                      //random number between 0 and 100
	fmt.Println(arrayWords[rand.Intn(len(arrayWords))])                                              //random word from arrayWords
}

func main() {
	Random() //calling Random function inside main, (entry point)
}
