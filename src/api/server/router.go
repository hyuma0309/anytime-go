package server

import (
    "github.com/gin-gonic/gin"
    "api/controller"
)
// GetRouter is router
func GetRouter() *gin.Engine {   
    router := gin.Default()
    router.LoadHTMLGlob("view/*.html")

    router.GET("/", controller.IndexDisplayAction)
    router.GET("/anytime", controller.AnytimeDisplayAction)

    return router
}

