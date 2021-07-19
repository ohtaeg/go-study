package main

import (
	"fmt"
	"github.com/ohtaeg/job-scrapper/dictionary"
)

func main() {

	dictionaryTest()
}

func dictionaryTest() {
	fmt.Println("## dictionary test ##")
	dictionary := dictionary.Dictionary{}

	// search
	result, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// add
	word := "hello"
	error := dictionary.Add(word, "Greeting")
	if error != nil {
		fmt.Println(error)
	}
	definition, err := dictionary.Search(word)
	fmt.Println(definition)

	// add exist
	existError := dictionary.Add(word, "Greeting")
	if existError != nil {
		fmt.Println(existError)
	}

	// update
	updateError := dictionary.Update(word, "Second")
	if updateError != nil {
		fmt.Println(updateError)
	}
	updatedWord, _ := dictionary.Search(word)
	fmt.Println(updatedWord)

	// delete
	dictionary.Delete(word)
	fmt.Println(dictionary.Search(word))
}
