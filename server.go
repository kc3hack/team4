package main

import (
    "./api"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/kr/pretty"
    "io/ioutil"
    "bytes"
    "net/http"
    "googlemaps.github.io/maps"
)

type CreateParams struct{
    OriginPoint OriginPoint `json:"originpoint"`
    DestinationPoint DestinationPoint `json:"destinationpoint"`
    WayPoint WayPoint `json:"waypoint"`
    Ways maps.Mode `json:"ways"`
}

type Data struct {
	Sanc_id int64 `json:"sanc_id"`
    Anime_id int64 `json:"anime_id"`
	//Area_id int64 `json:"area_id"`
	Place_id string `json:"place_id"`
}

type OriginPoint struct{
    OriginData Data `json:"origindata"`
}

type DestinationPoint struct{
    DestinationData Data `json:"destinatindata"`
}

type WayPoint struct{
    WayData WayData `json:"waydataarr"`
}

type WayData []*Data


type Data_arr []*Data

func main() {
    r := gin.Default()
    var way_data_arr WayData
    data_arr := []string{}
    //place_arr := []string{}
    //var waypoint_data_arr = WayData
    var params CreateParams
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
        if params.WayPoint.WayData != nil {
            c.IndentedJSON(200, api.MapsWithWayPoint(params.OriginPoint.OriginData.Place_id, params.DestinationPoint.DestinationData.Place_id, params.Ways, data_arr))
            pretty.Println(api.MapsWithWayPoint(params.OriginPoint.OriginData.Place_id, params.DestinationPoint.DestinationData.Place_id, params.Ways, data_arr))
        }else{
            c.IndentedJSON(200, api.Maps(params.OriginPoint.OriginData.Place_id, params.DestinationPoint.DestinationData.Place_id, params.Ways))
            pretty.Println(api.Maps(params.OriginPoint.OriginData.Place_id, params.DestinationPoint.DestinationData.Place_id, params.Ways))
        }
    })
    r.POST("/json", func(c *gin.Context){

        buf := make([]byte, 2048)
        n, _ := c.Request.Body.Read(buf)
        b := string(buf[0:n])
        c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(b)))
        fmt.Println(b)

        if err := c.BindJSON(&params); err != nil {
            //c.Warningf("Failed binding request parameters: %s", err.Error())
            c.Status(http.StatusBadRequest)
            return
        }

        c.JSON(200, params)
        //data_arr = append(data_arr, &params)
        //for _, i, j, k, l := range params {
            //fmt.Println(i.Place_id)
        //}

        fmt.Println(params)

        if params.WayPoint.WayData != nil {
            way_data_arr = params.WayPoint.WayData
            for _, i := range way_data_arr {
                fmt.Println(i.Place_id)
                data_arr = append(data_arr, i.Place_id)
            }
        }

        //place_arr = append(place_arr, data_arr[len(data_arr)-1].Place_id)
        //c.JSON(200, gin.H{"1": foo.Anime_id, "2": foo.Area_id})
    })
    r.Run()
}