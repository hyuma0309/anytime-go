package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/sclevine/agouti"
)

func IndexDisplayAction(c *gin.Context){
    c.HTML(200, "index.html", gin.H{"message": "hello gin"})
}

func AnytimeDisplayAction(c *gin.Context){
agoutiDriver := agouti.ChromeDriver()
    agoutiDriver.Start()
    defer agoutiDriver.Stop()
    page, _ := agoutiDriver.NewPage()
    
    
    // 六町店
   page.Navigate("https://www.anytimefitness.co.jp/rokucho/"); 
   title1, _ := page.Title()
   info1, _ := page.FindByID("info-box").Text()

   // 新柴又店
   page.Navigate("https://www.anytimefitness.co.jp/n-shibamata/"); 
   title2, _ := page.Title()
   info2, _ := page.FindByID("info-box").Text()


    c.HTML(200, "anytime.html", gin.H{
        "message": "エニタイム",
        "title1": title1,
        "info1": info1,
        "title2": title2,
        "info2": info2,
    })
}
