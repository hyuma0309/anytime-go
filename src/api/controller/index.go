package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sclevine/agouti"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func IndexDisplayAction(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{"message": "hello gin"})
}

func chrome() *agouti.Page {
	agoutiDriver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			"--headless", // headlessモードの指定 サイト読み込み時の起動をなくす
			"--no-sandbox",
			"--disable-gpu",           // アクセラレーションの排除　ちらつき解消
			"--disable-dev-shm-usage", //chromeが破損しないようDockerのメモリスペースを大きいものに変更
		}),
	)
	agoutiDriver.Start()
	// defer agoutiDriver.Stop()
	page, _ := agoutiDriver.NewPage()
	return page
}

type anytime struct {
	Title string
	Price int
	Info  string
	Url   *url.URL
}

type Lists []anytime

func (l Lists) Len() int {
	// Len is the number of elements in the collection.
	return len(l)
}

func (l Lists) Less(i, j int) bool {
	// Less reports whether the element with
	// index i should sort before the element with index j.
	return l[i].Price < l[j].Price
	// return i < j ← これだとうまく動かない
}

func (l Lists) Swap(i, j int) {
	// Swap swaps the elements with indexes i and j.
	l[i], l[j] = l[j], l[i]
}

func NewAnytime(title string, price int, info string, url *url.URL) anytime {
	return anytime{Title: title, Price: price, Info: info, Url: url}
}

func AnytimeDisplayAction(c *gin.Context) {
	ch := make(chan []anytime)

	go func() {
		page := chrome()
		link := "https://www.anytimefitness.co.jp/ebisu/"

		// 恵比寿店
		page.Navigate(link)
		title, _ := page.Title()
		priceClass, _ := page.FindByClass("price").Text()
		priceString := string([]rune(priceClass)[:23])

		// 文字列から数字を抜き出す
		re := regexp.MustCompile("(..[0-9]+)[^0-9]*$")
		result := re.FindStringSubmatch(priceString)
		if result == nil {
			fmt.Println("数字なし")
		}
		// 数字ののコンマを削除
		result[1] = strings.Replace(result[1], ",", "", 1)
		// 文字列をint型に変更
		price, _ := strconv.Atoi(result[1])

		url, _ := url.Parse(link)
		info, _ := page.Find("#campaign > div > div > p.fl").Text()
		info = info[:strings.Index(info, "※")]

		var y []anytime
		y = append(y, NewAnytime(title, price, info, url))

		// log.Println(y)
		ch <- y
	}()

	go func() {
		page := chrome()
		link := "https://www.anytimefitness.co.jp/keiosasazuka/"

		// 京王笹塚店
		page.Navigate(link)

		title, _ := page.Title()
		priceClass, _ := page.FindByClass("price").Text()
		priceString := string([]rune(priceClass)[:23])

		re := regexp.MustCompile("(..[0-9]+)[^0-9]*$")
		result := re.FindStringSubmatch(priceString)
		if result == nil {
			fmt.Println("数字なし")
		}
		result[1] = strings.Replace(result[1], ",", "", 1)
		price, _ := strconv.Atoi(result[1])

		url, _ := url.Parse(link)
		info, _ := page.Find("#campaign > div > div > p.fl").Text()
		info = info[:strings.Index(info, "※")]

		var y []anytime
		y = append(y, NewAnytime(title, price, info, url))

		ch <- y
	}()

	go func() {
		page := chrome()
		link := "https://www.anytimefitness.co.jp/hatsudai/"

		// 渋谷初台店
		page.Navigate(link)

		title, _ := page.Title()
		priceClass, _ := page.FindByClass("price").Text()
		priceString := string([]rune(priceClass)[:23])

		re := regexp.MustCompile("(..[0-9]+)[^0-9]*$")
		result := re.FindStringSubmatch(priceString)
		if result == nil {
			fmt.Println("数字なし")
		}
		result[1] = strings.Replace(result[1], ",", "", 1)
		price, _ := strconv.Atoi(result[1])

		url, _ := url.Parse(link)
		info := "キャンペーン情報はありません"

		var y []anytime
		y = append(y, NewAnytime(title, price, info, url))

		ch <- y
	}()

	go func() {
		page := chrome()
		link := "https://www.anytimefitness.co.jp/hiroohsf/"

		// 広尾高校前店
		page.Navigate(link)

		title, _ := page.Title()
		priceClass, _ := page.FindByClass("price").Text()
		priceString := string([]rune(priceClass)[:23])

		re := regexp.MustCompile("(..[0-9]+)[^0-9]*$")
		result := re.FindStringSubmatch(priceString)
		if result == nil {
			fmt.Println("数字なし")
		}
		result[1] = strings.Replace(result[1], ",", "", 1)
		price, _ := strconv.Atoi(result[1])

		url, _ := url.Parse(link)
		info, _ := page.Find("#campaign > div > div > p.fl").Text()
		info = info[:strings.Index(info, "※")]

		var y []anytime
		y = append(y, NewAnytime(title, price, info, url))

		ch <- y
	}()

	go func() {
		page := chrome()
		link := "https://www.anytimefitness.co.jp/yoyogi/"

		// 元代々木店
		page.Navigate(link)

		title, _ := page.Title()
		priceClass, _ := page.FindByClass("price").Text()
		priceString := string([]rune(priceClass)[:23])

		re := regexp.MustCompile("(..[0-9]+)[^0-9]*$")
		result := re.FindStringSubmatch(priceString)
		if result == nil {
			fmt.Println("数字なし")
		}
		result[1] = strings.Replace(result[1], ",", "", 1)
		price, _ := strconv.Atoi(result[1])

		url, _ := url.Parse(link)
		info, _ := page.Find("#campaign > div > div > p.fl").Text()
		info = info[:strings.Index(info, "※")]

		var y []anytime
		y = append(y, NewAnytime(title, price, info, url))

		ch <- y
	}()

	go func() {
		page := chrome()
		link := "https://www.anytimefitness.co.jp/yoyogi-sendagaya/"

		// 代々木店
		page.Navigate(link)

		title, _ := page.Title()
		priceClass, _ := page.FindByClass("price").Text()
		priceString := string([]rune(priceClass)[:23])

		re := regexp.MustCompile("(..[0-9]+)[^0-9]*$")
		result := re.FindStringSubmatch(priceString)
		if result == nil {
			fmt.Println("数字なし")
		}
		result[1] = strings.Replace(result[1], ",", "", 1)
		price, _ := strconv.Atoi(result[1])

		url, _ := url.Parse(link)
		info := "キャンペーン情報はありません"

		var y []anytime
		y = append(y, NewAnytime(title, price, info, url))

		ch <- y
	}()

	lists := Lists{}
	lists = append(lists, <-ch...)
	lists = append(lists, <-ch...)
	lists = append(lists, <-ch...)
	lists = append(lists, <-ch...)
	lists = append(lists, <-ch...)
	lists = append(lists, <-ch...)

	sort.Sort(lists)

	c.HTML(200, "anytime.html", gin.H{
		"message": "エニタイム",
		"titles":  lists,
	})
}
