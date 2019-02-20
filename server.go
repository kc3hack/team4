package main

import (
    "./api"
    "github.com/gin-gonic/gin"
)
func main() {
        r := gin.Default()
        r.GET("/", func(c *gin.Context) {
                c.String(200, "今日は世界!")
        })
        r.GET("/api", func(c *gin.Context){
            c.String(200, api.GetApi())
        })
        r.GET("/animes", func(c *gin.Context){
            c.JSON(200, api.GetAnimes())
        })
        r.GET("/sanctuaries", func(c *gin.Context){
            c.JSON(200, api.GetSanctuaries())
        })
        r.Run()
}