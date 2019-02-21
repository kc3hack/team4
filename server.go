package main

import (
    _ "github.com/go-sql-driver/mysql"
    "./api"
    "github.com/gin-gonic/gin"
    "database/sql"
    "fmt"
    "log"
    "net/http"
)
func main() {
        r := gin.Default()
        db, err := sql.Open("mysql","root:password@tcp(docker.for.mac.localhost:3306)/team4_db")
            log.Println("Connected to mysql.")
        //接続でエラーが発生した場合の処理
        if err != nil {
            log.Fatal(err)
        }
        r.GET("/", func(c *gin.Context) {
                c.String(200, "今日は世界!")
        })
        r.GET("/api", func(c *gin.Context){
            c.String(200, api.GetApi())
        })
        r.GET("/animes", func(c *gin.Context){
            var (
                animes_lists api.Animes_Lists
                animes []api.Animes_Lists
            )
            rows, err := db.Query("SELECT * FROM anime")
            if err != nil {
                panic(err.Error())
            }

            //レコード一件一件をあらかじめ用意しておいた構造体に当てはめていく。
            for rows.Next() { 
                err := rows.Scan(&animes_lists.Anime_ID, &animes_lists.Anime_Name, &animes_lists.Area_ID,&animes_lists.Image, &animes_lists.Url)
                animes = append(animes, animes_lists)
                if err != nil {
                    fmt.Print(err.Error())
                }
            }
            c.JSON(http.StatusOK, animes)
        })
        r.GET("/sanctuaries", func(c *gin.Context){
            var (
                sanctuary_lists api.Sanctuary_detals
                sanctuaries []api.Sanctuary_detals
            )
            rows, err := db.Query("SELECT * FROM sanctuary")
            if err != nil {
                panic(err.Error())
            }

            //レコード一件一件をあらかじめ用意しておいた構造体に当てはめていく。
            for rows.Next() { 
                err := rows.Scan(&sanctuary_lists.Sanc_ID, &sanctuary_lists.Sanc_Name, &sanctuary_lists.Anime_ID,&sanctuary_lists.Place_ID, &sanctuary_lists.Image)
                sanctuaries = append(sanctuaries, sanctuary_lists)
                if err != nil {
                    fmt.Print(err.Error())
                }
            }
            c.JSON(http.StatusOK, sanctuaries)
        })
        r.Run()
}