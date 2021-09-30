package standard

import (
	"io"
	"os"
)

func Print() {
	message := ""
	arguments := os.Args // os.Args => string 값을 원소로 갖는 슬라이스
	if len(arguments) == 1 {
		// go run main.go
		message = "it's default message"
	} else {
		// argument[0]은 실행 가능한 프로그램 이름(파일경로)이다.
		// go run main.go 123 12 => 123-12
		message = arguments[0] + "-" + arguments[1] + "-" + arguments[2]
	}

	io.WriteString(os.Stdout, message)
	io.WriteString(os.Stdout, "\n")
}
