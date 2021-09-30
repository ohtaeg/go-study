package standard

import (
	"fmt"
	"os"
	"strconv" // string 데이터를 산술 데이터 타입으로 변환
)

func Convert() {
	argumentSize := len(os.Args)
	if argumentSize == 1 {
		fmt.Print("plz input command line arguments")
		os.Exit(1)
	}

	arguments := os.Args
	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[1], 64)

	for i := 2; i < argumentSize; i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)
		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	fmt.Println("min : ", min)
	fmt.Println("max : ", max)
}
