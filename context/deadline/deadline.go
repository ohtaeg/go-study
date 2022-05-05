package main

import (
	"context"
	"fmt"
	"time"
)

var format = "2006-01-02 15:04:05"

func main() {
	ctx := context.Background()
	fmt.Println("시작 전 : " + time.Now().Format(format))
	timeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	deadline, ok := timeout.Deadline()
	fmt.Println("데드라인 : " + deadline.Format(format))
	fmt.Println(ok)

	cancel()

	time.Sleep(6 * time.Second)
	_, after := timeout.Deadline()
	fmt.Println(after)
}
