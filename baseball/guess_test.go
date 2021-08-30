package baseball

import (
	"testing"
)

func TestGuessValidNums(t *testing.T) {
	nums := "123"
	game, _ := CreateGame(nums)

	_, err := game.guess(nums)
	if err != nil {
		t.Fatalf("invalid numbers")
	}
}

func TestGuessInvalidNums(t *testing.T) {
	game, _ := CreateGame("123")

	assertGuessReturnError(t, game, "1")
	assertGuessReturnError(t, game, "12")
	assertGuessReturnError(t, game, "12345")
	assertGuessReturnError(t, game, "1a")
	assertGuessReturnError(t, game, "12a")
	assertGuessReturnError(t, game, "112")
}

func assertGuessReturnError(t *testing.T, game Game, nums string) {
	_, err := game.guess(nums)
	if err == nil {
		t.Fatalf("error must be returned: %s", nums)
	}
}
