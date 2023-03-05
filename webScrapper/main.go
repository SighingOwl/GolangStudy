package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"goquery"
)

var baseURL string = "http://pptbizcam.co.kr/?cat=2" // 조땡 템플릿 공유 페이지

func main() {
	totalpages := getPages()

	for i := 0; i < totalpages; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	pageURL := baseURL + "&paged=" + strconv.Itoa(page)
	fmt.Println("Requesting", pageURL)
}

func getPages() int {
	var pages string = "0"
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body) //res.Body는 기본적으로 IO을 수행하므로 사용하지 않을 때 close 해서 메모리 누수를 방지한다.
	checkErr(err)

	doc.Find(".pgntn-page-pagination-block").Each(func(i int, s *goquery.Selection) {
		pages = s.Find(".page-numbers").Text()
		fmt.Println(pages)
	})

	totalPages := pageParser(pages)

	return totalPages
}

func pageParser(pages string) int {
	startIndex := 0
	lastIndex := 0

	for i := 0; i < len(pages); i++ {
		dotCom := pages[i : i+3]
		nCom := pages[i : i+4]

		if dotCom == "…" {
			fmt.Println("dot", i)
			startIndex = i + 3
		}
		if nCom == "Next" {
			fmt.Println("n", i)
			lastIndex = i
			break
		}
	}
	lastPage, _ := strconv.Atoi(pages[startIndex:lastIndex])

	return lastPage
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}
