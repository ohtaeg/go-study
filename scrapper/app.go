package main

import (
	"fmt"
	"net/http"
	"time"
)

type scrapResult struct {
	url    string
	status string
}

func main() {
	goroutineExample()
	urlScrap()
}

func goroutineExample() {
	channel := make(chan string)
	people := [2]string{"nico", "ohtaeg"}
	for _, person := range people {
		go isGood(person, channel)
	}

	for i := 0; i < len(people); i++ {
		fmt.Println(<-channel)
	}
}

func isGood(person string, channel chan string) {
	time.Sleep(time.Second * 5)
	channel <- person + " is good"
}

func urlScrap() {
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	c := make(chan scrapResult)
	for _, url := range urls {
		go hitUrl(url, c)
	}

	for range urls {
		fmt.Println(<-c)
	}
}

// 이 함수는 채널에 데이터를 보낼 수만 있고 받기만 할 수 있어라는 설정
// chan<- : Send only
func hitUrl(url string, c chan<- scrapResult) {
	status := "OK"
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- scrapResult{
		url:    url,
		status: status,
	}
}
