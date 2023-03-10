package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"goquery"
)

type extractedItem struct {
	itemNo    string
	itemBrand string
	itemName  string
	itemPrice string
	itemScore string
	itemShop  string
}

// var baseURL string = "http://pptbizcam.co.kr/?cat=2" // 조땡 템플릿 공유 페이지
// var baseURL string = "http://www.yes24.com/Product/Search?domain=ALL&query=%ED%81%B4%EB%9D%BC%EC%9A%B0%EB%93%9C" //yes24 클라우드 관련 서적 검색 페이지
// var baseURL string = "https://product.kyobobook.co.kr/category/KOR/26#?page=1&type=all&per=20&sort=new" //교보문고 기술/공학 관련 페이지
// var baseURL string = "http://browse.auction.co.kr/search?keyword=sony&itemno=&nickname=&frm=hometab&dom=auction&isSuggestion=No&retry=&Fwk=sony&acode=SRP_SU_0100&arraycategory=&encKeyword=sony&k=9" // Auction sony 검색결과
var baseURL string = "https://browse.auction.co.kr/search?keyword=LG&itemno=&nickname=&encKeyword=LG&arraycategory=&frm=&dom=auction&isSuggestion=No&retry=&k=29"

func main() {
	var items []extractedItem
	c := make(chan []extractedItem)
	totalpages := getPages()

	for i := 1; i <= totalpages; i++ {
		go getPage(i, c)
	}

	for i := 0; i < totalpages; i++ {
		extractItems := <-c
		items = append(items, extractItems...)
	}

	writeItems(items)
	fmt.Println("Done, extracted", len(items))
}

func writeItems(items []extractedItem) { // 수집된 item 정보를 csv파일에 작성
	var mu = new(sync.Mutex) // write 작업시 slice 보호를 위해 mutex를 사용

	file, err := os.Create("LGitems.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"Itemno", "Itembrand", "Itemname", "Itemprice", "Itemscore", "Itemshop"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, item := range items {
		itemSlice := []string{item.itemNo, item.itemBrand, item.itemName, item.itemPrice, item.itemScore, item.itemShop}
		go itemWrite(mu, w, itemSlice)
	}
}

func itemWrite(mu *sync.Mutex, w *csv.Writer, itemSlice []string) { // csv에 slice 작성
	mu.Lock()
	iwErr := w.Write(itemSlice)
	checkErr(iwErr)
	mu.Unlock()
}

func getPage(page int, mainC chan<- []extractedItem) { //page에서 item을 추출해서 channel을 사용해 main으로 전달
	var items []extractedItem
	c := make(chan extractedItem)
	pageURL := baseURL + "&p=" + strconv.Itoa(page)
	fmt.Println("Requesting", pageURL)

	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchitem := doc.Find(".section--itemcard")
	searchitem.Each(func(i int, card *goquery.Selection) { //doc.Find(".pgntn-page-pagination-block").Each(func(i int, s *goquery.Selection) 이런 형식으로 붙여써도 된다.
		go extractItem(card, c)
	})

	for i := 0; i < searchitem.Length(); i++ {
		item := <-c
		items = append(items, item)
	}
	mainC <- items
}

func extractItem(card *goquery.Selection, c chan<- extractedItem) { // getPage에서 추출한 itemcard에서 item의 정보를 추출
	item, _ := card.Find("a").Attr("href")
	id := cleanString(strings.Join(strings.Split(item, "http://itempage3.auction.co.kr/DetailView.aspx?itemno="), " ")) //href에 포함되어 있는 itemno를 추출
	itemBrand := cleanString(card.Find(".text--brand").Text())
	itemName := cleanString(card.Find(".text--title").Text())
	itemPrice := cleanString(card.Find(".text--price_seller").Text() + "원")
	itemScore := cleanString(card.Find(".awards").Text())
	shop, _ := card.Find(".link--shop").Attr("href")
	itemShop := cleanString(strings.Join(strings.Split(shop, "http://stores.auction.co.kr/"), " ")) //href에 포함되어 있는 판매자 정보 추출

	c <- extractedItem{itemNo: id,
		itemBrand: itemBrand,
		itemName:  itemName,
		itemPrice: itemPrice,
		itemScore: itemScore,
		itemShop:  itemShop}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getPages() int { // 스크랩핑을 진행할 사이트의 페이지 수를 확인
	var pages int = 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body) //res.Body는 기본적으로 IO을 수행하므로 사용하지 않을 때 close 해서 메모리 누수를 방지한다.
	checkErr(err)

	doc.Find(".component--pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func pageParser(pages string) int { // 조뗑 템플릿 공유 사이트 페이지 파인
	startIndex := 0
	lastIndex := 0

	for i := 0; i < len(pages); i++ {
		dotCom := pages[i : i+3]
		nCom := pages[i : i+4]

		if dotCom == "…" {
			startIndex = i + 3
		}
		if nCom == "Next" {
			lastIndex = i
			break
		}
	}
	lastPage, _ := strconv.Atoi(pages[startIndex:lastIndex])

	return lastPage
}

func checkErr(err error) { // err 확인
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) { // html 응답 코드 확인 및 예외 처리
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}
