package baseball

import (
	"github.com/stretchr/testify/assert"
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

func TestZeroStrike(t *testing.T) {
	game, _ := CreateGame("123")
	guessNumber := "567"
	expected := 0

	score, _ := game.guess(guessNumber)

	assert.Equal(t, expected, score.strike)
}

func TestOneStrike(t *testing.T) {
	game, _ := CreateGame("123")
	guessNumber := "167"
	expected := 1

	score, _ := game.guess(guessNumber)

	assert.Equal(t, expected, score.strike)
}

func TestTwoStrike(t *testing.T) {
	game, _ := CreateGame("123")
	guessNumber := "124"
	expected := 2

	score, _ := game.guess(guessNumber)

	assert.Equal(t, expected, score.strike)
}

func TestThreeStrike(t *testing.T) {
	game, _ := CreateGame("123")
	guessNumber := "123"
	expected := 3

	score, _ := game.guess(guessNumber)

	assert.Equal(t, expected, score.strike)
}

func TestZeroBall(t *testing.T) {
	game, _ := CreateGame("123")
	guessNumber := "567"
	expected := 0

	score, _ := game.guess(guessNumber)

	assert.Equal(t, expected, score.ball)
}

func TestOneBall(t *testing.T) {
	game, _ := CreateGame("123")
	guessNumber := "517"
	expected := 1

	score, _ := game.guess(guessNumber)

	assert.Equal(t, expected, score.ball)
}

func TestTwoBall(t *testing.T) {
	game, _ := CreateGame("123")
	guessNumber := "512"
	expected := 2

	score, _ := game.guess(guessNumber)

	assert.Equal(t, expected, score.ball)
}

func TestThreeBall(t *testing.T) {
	game, _ := CreateGame("123")
	guessNumber := "312"
	expected := 3

	score, _ := game.guess(guessNumber)

	assert.Equal(t, expected, score.ball)
}
