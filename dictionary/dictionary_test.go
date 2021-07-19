package dictionary

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDictionary_Add(t *testing.T) {
	// given
	dictionary := Dictionary{}
	expect := "world"

	// when
	err := dictionary.Add("hello", "world")

	// then
	assert.Nil(t, err)
	assert.Equal(t, expect, dictionary["hello"])
}

func TestDictionary_Add_ErrWordExists(t *testing.T) {
	// given
	dictionary := Dictionary{}
	dictionary.Add("hello", "world")
	// expect := errors.New("error")

	// when
	err := dictionary.Add("hello", "world")

	// then
	assert.Error(t, err)
}

func TestDictionary_Search(t *testing.T) {
	// given
	dictionary := Dictionary{}
	expect := "world"
	dictionary.Add("hello", expect)

	// when
	actual, _ := dictionary.Search("hello")

	// then
	assert.Equal(t, expect, actual)
}

func TestDictionary_Search_ErrorNotFound(t *testing.T) {
	// given
	dictionary := Dictionary{}
	expect := "world"
	dictionary.Add("hello", expect)

	// when
	_, error := dictionary.Search("hello2")

	// then
	assert.Error(t, error)
}

func TestDictionary_Update(t *testing.T) {
	// given
	expect := "world2"
	dictionary := Dictionary{}
	dictionary.Add("hello", "world")

	// when
	error := dictionary.Update("hello", expect)

	// then
	assert.Nil(t, error)
	assert.Equal(t, expect, dictionary["hello"])
}
