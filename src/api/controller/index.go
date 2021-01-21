package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sclevine/agouti"
	// "log"
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

func AnytimeDisplayAction(c *gin.Context) {

	messages := []string{}
    test := make(chan []string)

	go func(){
		page := chrome()

		// 六町店
		page.Navigate("https://www.anytimefitness.co.jp/rokucho/")
		title, _ := page.Title()
		info, _ := page.FindByID("info-box").Text()
		
		y := []string{}
		y = append(y, title,info)
		
		test <- y
	}()

	go func() {
		page := chrome()

		// 新柴又店
		page.Navigate("https://www.anytimefitness.co.jp/n-shibamata/")

		title, _ := page.Title()
		info, _ := page.FindByID("info-box").Text()

		y := []string{}
		y = append(y, title,info)

		test <- y
	}()

	go func() {
		page := chrome()

		// 曙橋駅店
		page.Navigate("https://www.anytimefitness.co.jp/akebonobashi/")

		title, _ := page.Title()
		info, _ := page.FindByID("info-box").Text()

		y := []string{}
		y = append(y, title,info)

		test <- y
	}()


	go func() {
		page := chrome()

		// 新中野店
		page.Navigate("https://www.anytimefitness.co.jp/shinnakano/")

		title, _ := page.Title()
		info, _ := page.FindByID("info-box").Text()

		y := []string{}
		y = append(y, title,info)

		test <- y
	}()

	go func() {
		page := chrome()

		// 落合店
		page.Navigate("https://www.anytimefitness.co.jp/ochiai/")

		title, _ := page.Title()
		info, _ := page.FindByID("info-box").Text()

		y := []string{}
		y = append(y, title,info)

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
