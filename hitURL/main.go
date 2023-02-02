package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	//var results map[string]string //이렇게 map을 정의하고 값을 넣으면 panic이 발생한다. panic은 컴파일러가 모르는 에러
	var results = map[string]string{} //make(map[string]string) 이렇게 map을 만들어도 된다.
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
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)                //http와 관련된 method는 Go 표준 라이브러리에 있다.
	if err != nil || resp.StatusCode >= 400 { //http 응답코드가 400이상부터는 문제가 있으므로 예외처리를 해야한다.
		return errRequestFailed
	}
	return nil
}
