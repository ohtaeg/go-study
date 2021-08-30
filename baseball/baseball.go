package baseball

import "fmt"

type Game interface {
	guess(nums string) (Score, error)
}

type gameNumbers struct {
	numbers string
}

type Score struct {
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

	return Score{}, nil
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
