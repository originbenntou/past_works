package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

type error interface {
	Error() string
}

// 後学のためスタックトレースのやり方を残す
func getStackTrace() {
	pt, file, line, ok := runtime.Caller(0)
	if !ok {
		panic("failed stack trace")
	}

	fmt.Printf("file=%s, line=%d, func=%v\n", file, line, runtime.FuncForPC(pt).Name())
}

type shop struct {
	name     string
	minPrice int
	maxPrice int
	star     float64
	holiday  string
	photo    string
	url      string
	location string
}

func createShop(document *goquery.Document, shops []shop) []shop {
	var err error

	document.Find("li.list-rst").Each(func(i int, selection *goquery.Selection) {
		shops[i].name = selection.Find("div.list-rst__rst-name > a").Text()

		prices := make([]string, 2)
		split := strings.Split(selection.Find("span.cpy-lunch-budget-val").Text(), "～")
		if len(split) != 2 {
			prices[0] = "0"
			prices[1] = "0"
		} else {
			prices = split
		}

		min := regexp.MustCompile(`[￥|,]`).ReplaceAllString(prices[0], "")
		max := regexp.MustCompile(`[￥|,]`).ReplaceAllString(prices[1], "")

		// 記載なしの場合、"0"を代わりにセット
		if min == " " {
			min = "0"
		}
		if max == " " {
			max = "0"
		}

		shops[i].minPrice, err = strconv.Atoi(min)
		if err != nil {
			log.Fatal(err)
		}

		shops[i].maxPrice, err = strconv.Atoi(max)
		if err != nil {
			log.Fatal(err)
		}

		shops[i].star, err = strconv.ParseFloat(selection.Find("span.list-rst__rating-val").Text(), 64)
		if err != nil {
			log.Fatal(err)
		}

		shops[i].holiday = selection.Find("span.list-rst__holiday-datatxt").Text()

		photo, exist := selection.Find("img.cpy-main-image").Attr("src")
		if exist == false {
			log.Fatal("photo is not found")
		}
		shops[i].photo = photo

		url, exist := selection.Find("a.list-rst__image-target").Attr("href")
		if exist == false {
			log.Fatal("url is not found")
		}
		shops[i].url = url

		lowerDoc, err := goquery.NewDocument(url)
		if err != nil {
			log.Fatal(err)
		}
		shops[i].location = lowerDoc.Find("p.rstinfo-table__address").Text()
	})

	return shops
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	log.Print("Start")

	document, err := goquery.NewDocument("https://tabelog.com/fukuoka/A4004/A400401/R3954/rstLst/1/?Srt=D&SrtT=rtl")
	if err != nil {
		log.Fatal(err)
	}

	count, err := strconv.ParseInt(document.Find("span.list-condition__count").Text(), 10, 0)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Page is %v", count)

	shops := make([]shop, 0, count)

	// 並列処理は最大3つまで
	limit := make(chan struct{}, 3)

	var wg sync.WaitGroup
	// 1ページ目から
	for i := 1; i < 5; i++ {
		page := i

		// メモリリーク
		runtime.GC()

		wg.Add(1)
		go func() {
			defer wg.Done()

			limit <- struct{}{}

			document, err := goquery.NewDocument("https://tabelog.com/fukuoka/A4004/A400401/R3954/rstLst/" + strconv.Itoa(page) + "/?Srt=D&SrtT=rtl")
			if err != nil {
				log.Fatal(err)
			}

			s := make([]shop, 20)
			s = createShop(document, s)

			shops = append(shops, s...)

			<-limit

			log.Printf("SET %d'sParams: OK, %v is complete", page, len(shops))
		}()
	}

	wg.Wait()

	log.Printf("All Complete, Number of shops => %d", len(shops))
}
