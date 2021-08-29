package baseball

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateGame(t *testing.T) {
	game, _ := CreateGame("012")
	if game == nil {
		t.Fatalf("Game create fail")
	}
}

func TestValidNumbers(t *testing.T) {
	game, err := CreateGame("345")
	if game == nil || err != nil {
		t.Fatalf("game must be returned")
	}
}

func TestInvalidNumbers(t *testing.T) {
	assertInvalidNumbersError(t, "01")
	assertInvalidNumbersError(t, "01234")
	assertInvalidNumbersError(t, "abc")
	assertInvalidNumbersError(t, "1bc")
}

func assertInvalidNumbersError(t *testing.T, numbers string) {
	// when
	_, err := CreateGame(numbers)

	// then
	assert.Error(t, err)
}
