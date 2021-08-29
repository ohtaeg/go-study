package baseball

import "fmt"

type Game interface {
}

type gameNumbers struct {
	numbers string
}

func CreateGame(numbers string) (Game, error) {
	if len(numbers) < 3 || len(numbers) > 4 {
		return nil, fmt.Errorf("invalid nums : %s", numbers)
	}

	for i := 0; i < len(numbers); i++ {
		if !(numbers[i] >= '0' && numbers[i] <= '9') {
			return nil, fmt.Errorf("invalid nums : %s", numbers)
		}
	}

	used := make(map[uint8]bool)
	for i := 0; i < len(numbers); i++ {
		found := used[numbers[i]]
		if found {
			return nil, fmt.Errorf("invalid nums : %s", numbers)
		}
		used[numbers[i]] = true
	}
	return &gameNumbers{numbers: numbers}, nil
}
