package controller

import (
	"fmt"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sclevine/agouti"
)

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

// Anytime is struct
type Anytime struct {
	Title string
	Price int
	Info  string
	Url   *url.URL
}

// Lists is []Anytime
type Lists []Anytime

func (l Lists) Len() int {
	return len(l)
}

func (l Lists) Less(i, j int) bool {
	return l[i].Price < l[j].Price
}

func (l Lists) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// NewAnytime is make Anytime
func NewAnytime(title string, price int, info string, url *url.URL) Anytime {
	return Anytime{Title: title, Price: price, Info: info, Url: url}
}

// AnytimeDisplayAction is AnytimeDisplayAction
func AnytimeDisplayAction(c *gin.Context) {
	ch := make(chan []Anytime)

	//渋谷区
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

		var y []Anytime
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

		var y []Anytime
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

		var y []Anytime
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

		var y []Anytime
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

		var y []Anytime
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

		var y []Anytime
		y = append(y, NewAnytime(title, price, info, url))

		ch <- y
	}()

	// 港区
	go func() {
		page := chrome()
		link := "https://www.anytimefitness.co.jp/shibaura/"

		// 芝浦店
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

		var y []Anytime
		y = append(y, NewAnytime(title, price, info, url))

		ch <- y
	}()

	go func() {
		page := chrome()
		link := "https://www.anytimefitness.co.jp/mita/"

		// 三田店
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

		var y []Anytime
		y = append(y, NewAnytime(title, price, info, url))

		ch <- y
	}()

	go func() {
		page := chrome()
		link := "https://www.anytimefitness.co.jp/shimbashi/"

		// 新橋店
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

		var y []Anytime
		y = append(y, NewAnytime(title, price, info, url))

		ch <- y
	}()

	go func() {
		page := chrome()
		link := "https://www.anytimefitness.co.jp/hamamatsuchoshiodome/"

		// 浜松町汐留店
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

		var y []Anytime
		y = append(y, NewAnytime(title, price, info, url))

		ch <- y
	}()

	go func() {
		page := chrome()
		link := "https://www.anytimefitness.co.jp/bentencho/"

		// 弁天町店
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

		var y []Anytime
		y = append(y, NewAnytime(title, price, info, url))

		ch <- y
	}()

	go func() {
		page := chrome()
		link := "https://www.anytimefitness.co.jp/shibah/"

		// 芝浜松町店
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

		var y []Anytime
		y = append(y, NewAnytime(title, price, info, url))

		ch <- y
	}()

	go func() {
		page := chrome()
		link := "https://www.anytimefitness.co.jp/hiroo/"

		// 広尾店
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

		var y []Anytime
		y = append(y, NewAnytime(title, price, info, url))

		ch <- y
	}()

	go func() {
		page := chrome()
		link := "https://www.anytimefitness.co.jp/akasaka/"

		// 赤坂店
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

		var y []Anytime
		y = append(y, NewAnytime(title, price, info, url))

		ch <- y
	}()

	go func() {
		page := chrome()
		link := "https://www.anytimefitness.co.jp/s-azabu3/"

		// 南麻布3丁目店
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

		var y []Anytime
		y = append(y, NewAnytime(title, price, info, url))

		ch <- y
	}()

	go func() {
		page := chrome()
		link := "https://www.anytimefitness.co.jp/s-azabu/"

		// 南麻布2丁目店
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

		var y []Anytime
		y = append(y, NewAnytime(title, price, info, url))

		ch <- y
	}()

	go func() {
		page := chrome()
		link := "https://www.anytimefitness.co.jp/sengakuji/"

		// 泉岳寺駅前店
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

		var y []Anytime
		y = append(y, NewAnytime(title, price, info, url))

		ch <- y
	}()

	go func() {
		page := chrome()
		link := "https://www.anytimefitness.co.jp/asashiobashi/"

		// 朝潮橋店
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

		var y []Anytime
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
	lists = append(lists, <-ch...)
	lists = append(lists, <-ch...)
	lists = append(lists, <-ch...)
	lists = append(lists, <-ch...)
	lists = append(lists, <-ch...)
	lists = append(lists, <-ch...)
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
