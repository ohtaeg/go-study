package baseball

import "fmt"

type Game interface {
	guess(nums string) (Score, error)
}

type gameNumbers struct {
	numbers string
}

type Score struct {
	strike int
	ball   int
}

func CreateGame(numbers string) (Game, error) {
	if len(numbers) < 3 || len(numbers) > 4 {
		return nil, fmt.Errorf("invalid nums : %s", numbers)
	}

	if containsNonDigit(numbers) {
		return nil, fmt.Errorf("invalid numbers : %s", numbers)
	}

	if hasDuplicate(numbers) {
		return nil, fmt.Errorf("invalid nums : %s", numbers)
	}

	return &gameNumbers{numbers: numbers}, nil
}

func (g gameNumbers) guess(numbers string) (Score, error) {
	if len(g.numbers) != len(numbers) {
		return Score{}, fmt.Errorf("length mismatch")
	}

	if containsNonDigit(numbers) {
		return Score{}, fmt.Errorf("invalid numbers : %s", numbers)
	}

	if hasDuplicate(numbers) {
		return Score{}, fmt.Errorf("invalid nums : %s", numbers)
	}

	strike := calculateStrike(g.numbers, numbers)
	ball := calculateBall(g.numbers, numbers)

	return Score{
		strike: strike,
		ball:   ball,
	}, nil
}

func calculateStrike(numbers string, guessNumbers string) int {
	strike := 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == guessNumbers[i] {
			strike++
		}
	}
	return strike
}

func calculateBall(numbers string, guessNumbers string) int {
	ball := 0
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if numbers[i] == guessNumbers[j] && i != j {
				ball++
			}
		}
	}
	return ball
}

func containsNonDigit(numbers string) bool {
	for i := 0; i < len(numbers); i++ {
		if !(numbers[i] >= '0' && numbers[i] <= '9') {
			return true
		}
	}
	return false
}

func hasDuplicate(numbers string) bool {
	used := make(map[uint8]bool)
	for i := 0; i < len(numbers); i++ {
		found := used[numbers[i]]
		if found {
			return true
		}
		used[numbers[i]] = true
	}
	return false
}
