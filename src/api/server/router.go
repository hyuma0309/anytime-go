package server

import (
    "github.com/gin-gonic/gin"
    "api/controller"
)

func GetRouter() *gin.Engine {    // *gin.Engineの表記は返り値の型
    router := gin.Default()
    router.LoadHTMLGlob("view/*.html")

    router.GET("/", controller.IndexDisplayAction)
    router.GET("/anytime", controller.AnytimeDisplayAction)

    return router
}

