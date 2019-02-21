package main

import (
    "./api"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/kr/pretty"
    "io/ioutil"
    "bytes"
    "net/http"
)


type Data struct {
    Anime_id int64 `json:"anime_id"`
	Area_id int64 `json:"area_id"`
	Sanc_id int64 `json:"sanc_id"`
	Place_id string `json:"place_id"`
}


type Data_arr []*Data

func main() {
    r := gin.Default()
    var data_arr Data_arr
    place_arr := []string{}
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
        c.IndentedJSON(200, api.Maps(place_arr[0], place_arr[len(place_arr)-1]))
        pretty.Println(api.Maps("Rokujizo Station", "Port Island Kita Park"))
    })
    r.POST("/json", func(c *gin.Context){

        buf := make([]byte, 2048)
        n, _ := c.Request.Body.Read(buf)
        b := string(buf[0:n])
        c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(b)))
        fmt.Println(b)

        var params Data

        if err := c.BindJSON(&params); err != nil {
            //c.Warningf("Failed binding request parameters: %s", err.Error())
            c.Status(http.StatusBadRequest)
            return
        }

        c.JSON(200, params)
        data_arr = append(data_arr, &params)
        for _, i := range data_arr {
            fmt.Println(i.Place_id)
        }

        place_arr = append(place_arr, data_arr[len(data_arr)-1].Place_id)
        //c.JSON(200, gin.H{"1": foo.Anime_id, "2": foo.Area_id})
    })
    r.Run()
}