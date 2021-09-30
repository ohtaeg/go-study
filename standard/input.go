package standard

import (
	"bufio"
	"fmt"
	"os"
)

func Input() {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	// bufio.Scanner를 리턴한다. 이 스캐너는 파라미터로 넘긴 표준 입력으로부터 입력된 내용을 한 줄씩 읽는다.
	// go run main.go
	// test
	// test1
	// TODO 유닉스 파이프
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
}
