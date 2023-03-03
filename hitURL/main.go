package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

//var errRequestFailed = errors.New("Request failed")

func main() {
	//var results map[string]string //이렇게 map을 정의하고 값을 넣으면 panic이 발생한다. panic은 컴파일러가 모르는 에러
	results := map[string]string{} //make(map[string]string) 이렇게 map을 만들어도 된다.
	c := make(chan requestResult)
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
	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- requestResult) { // 이 함수에 있는 channel은 인자에 "chan<-"를 사용해서 송신만 가능하도록 지정할 수 있다.
	//fmt.Println("Checking:", url)
	//fmt.Println(<-c) // channel을 send only로 지정해서 메시지를 수신할 수 없다.

	resp, err := http.Get(url) //http와 관련된 method는 Go 표준 라이브러리에 있다.
	status := "OK"
	if err != nil || resp.StatusCode >= 400 { //http 응답코드가 400이상부터는 문제가 있으므로 예외처리를 해야한다.
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
}
