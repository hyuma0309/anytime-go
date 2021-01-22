package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sclevine/agouti"
	"log"
	"net/url"
)

func IndexDisplayAction(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{"message": "hello gin"})
}

func chrome() *agouti.Page {
	agoutiDriver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			"--headless", // headlessモードの指定 サイト読み込み時の起動をなくす
			"--no-sandbox", 
			"--disable-gpu", // アクセラレーションの排除　ちらつき解消
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
  Price string
  Url *url.URL
}

func NewAnytime(title string, price string, url *url.URL) anytime {
    return anytime{Title: title, Price: price, Url: url}
}

func AnytimeDisplayAction(c *gin.Context) {
	var messages []interface{}
	ch := make(chan []interface{})

	go func(){
		page := chrome()
		link := "https://www.anytimefitness.co.jp/rokucho/"

		// 六町店
		page.Navigate(link)
		title, _ := page.Title()
		priceClass, _ := page.FindByClass("price").Text()
		price := string([]rune(priceClass)[:23])
		url, _ := url.Parse(link)

		
		var y []interface{}
		y = append(y, NewAnytime(title,price,url))

		log.Println(y)
		ch <- y
	}()

	go func() {
		page := chrome()
		link := "https://www.anytimefitness.co.jp/n-shibamata/"

		// 新柴又店
		page.Navigate(link)

		title, _ := page.Title()
		priceClass, _ := page.FindByClass("price").Text()
		price := string([]rune(priceClass)[:23])
		url, _ := url.Parse(link)
		
		var y []interface{}
		y = append(y, NewAnytime(title,price,url))

		ch <- y
	}()

	go func() {
		page := chrome()
		link := "https://www.anytimefitness.co.jp/akebonobashi/"

		// 曙橋駅店
		page.Navigate(link)

		title, _ := page.Title()
		priceClass, _ := page.FindByClass("price").Text()
		price := string([]rune(priceClass)[:23])
		url, _ := url.Parse(link)
		
		var y []interface{}
		y = append(y, NewAnytime(title,price,url))

		ch <- y
	}()


	go func() {
		page := chrome()
		link := "https://www.anytimefitness.co.jp/shinnakano/"

		// 新中野店
		page.Navigate(link)

		title, _ := page.Title()
		priceClass, _ := page.FindByClass("price").Text()
		price := string([]rune(priceClass)[:23])
		url, _ := url.Parse(link)
		
		var y []interface{}
		y = append(y, NewAnytime(title,price,url))

		ch <- y
	}()

	go func() {
		page := chrome()
		link := "https://www.anytimefitness.co.jp/ochiai/"

		// 落合店
		page.Navigate(link)

		title, _ := page.Title()
		priceClass, _ := page.FindByClass("price").Text()
		price := string([]rune(priceClass)[:23])
		url, _ := url.Parse(link)
		
		var y []interface{}
		y = append(y, NewAnytime(title,price,url))

		ch <- y
	}()

	messages = append(messages, <-ch...)
	messages = append(messages, <-ch...)
	messages = append(messages, <-ch...)
	messages = append(messages, <-ch...)
	messages = append(messages, <-ch...)

	c.HTML(200, "anytime.html", gin.H{
		"message": "エニタイム",
		"titles":  messages,
	})
}
