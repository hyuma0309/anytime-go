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
  Info string
  Url *url.URL
}

func NewAnytime(title string, info string, url *url.URL) anytime {
    return anytime{Title: title, Info: info, Url: url}
}

func AnytimeDisplayAction(c *gin.Context) {


	var messages []interface{}
	test := make(chan []interface{})

	go func(){
		page := chrome()

		// 六町店
		page.Navigate("https://www.anytimefitness.co.jp/rokucho/")
		title, _ := page.Title()
		info, _ := page.FindByClass("price").Text()
		url, _ := url.Parse("https://www.anytimefitness.co.jp/rokucho/")
		
		var y []interface{}
		y = append(y, NewAnytime(title,info,url))

		log.Println(y)
		test <- y
	}()

	go func() {
		page := chrome()

		// 新柴又店
		page.Navigate("https://www.anytimefitness.co.jp/n-shibamata/")

		title, _ := page.Title()
		info, _ := page.FindByID("info-box").Text()
		url, _ := url.Parse("https://www.anytimefitness.co.jp/n-shibamata/")
		
		var y []interface{}
		y = append(y, NewAnytime(title,info,url))

		log.Println(y)
		test <- y
	}()

	go func() {
		page := chrome()

		// 曙橋駅店
		page.Navigate("https://www.anytimefitness.co.jp/akebonobashi/")

		title, _ := page.Title()
		info, _ := page.FindByID("info-box").Text()
		url, _ := url.Parse("https://www.anytimefitness.co.jp/akebonobashi/")
		
		var y []interface{}
		y = append(y, NewAnytime(title,info,url))

		log.Println(y)
		test <- y
	}()


	go func() {
		page := chrome()

		// 新中野店
		page.Navigate("https://www.anytimefitness.co.jp/shinnakano/")

		title, _ := page.Title()
		info, _ := page.FindByID("info-box").Text()
		url, _ := url.Parse("https://www.anytimefitness.co.jp/shinnakano/")
		
		var y []interface{}
		y = append(y, NewAnytime(title,info,url))

		log.Println(y)
		test <- y
	}()

	go func() {
		page := chrome()

		// 落合店
		page.Navigate("https://www.anytimefitness.co.jp/ochiai/")

		title, _ := page.Title()
		info, _ := page.FindByID("info-box").Text()
		url, _ := url.Parse("https://www.anytimefitness.co.jp/ochiai/")
		
		var y []interface{}
		y = append(y, NewAnytime(title,info,url))

		log.Println(y)
		test <- y
	}()

	messages = append(messages, <-test...)
	messages = append(messages, <-test...)
	messages = append(messages, <-test...)
	messages = append(messages, <-test...)
	messages = append(messages, <-test...)

	c.HTML(200, "anytime.html", gin.H{
		"message": "エニタイム",
		"titles":  messages,
	})
}
