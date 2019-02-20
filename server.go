package main

import (
    "./api"
    "github.com/gin-gonic/gin"
    "github.com/kr/pretty"
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
            c.String(200, api.Animes_list())
        })
	    r.GET("/areas", func(c *gin.Context){
            c.JSON(200, api.Areas_list())
        })
        r.GET("/direction", func(c *gin.Context){
            c.IndentedJSON(200, api.Maps())
            pretty.Println(api.Maps())
        })
        r.Run()
}