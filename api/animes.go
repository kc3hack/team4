package api

type Animes struct {
	Anime_ID int `json:"anime_id"`
	Anime_Name string `json:"anime_name"`
	Image string `json:"image"`
	Anime_details Anime_details `json:"animes_details"`
}

type Anime_details struct{
	Area_ID int `json:"area_id"`
	Url string `json:"url"`
}

type ResponseData struct {
	Datas        []Animes_Lists   `json:"datas"`
}

type Animes_Lists struct{
	// Anime_List []Animes `json:"anime_list"`
	Anime_ID int `json:"anime_id"`
	Anime_Name string `json:"anime_name"`
	Area_ID int `json:"area_id"`
	Image string `json:"image"`
	Url string `json:"url"`
}

// func GetAnimes() Animes_Lists{
// 	return Animes_Lists{
// 		Anime_List: []Animes{
// 			Animes{
// 				Anime_ID: 1,
// 				Anime_Name: "けいおん！",
// 				Image: "",
// 				Anime_details: Anime_details{
// 					Area_ID: 1,
// 					Url: "http://www.tbs.co.jp/anime/k-on/",
// 				},
// 			},
// 		},
// 	}
// }	