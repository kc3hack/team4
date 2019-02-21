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
	Anime_ID int `json:"anime_id"`
	Anime_Name string `json:"anime_name"`
	Area_ID int `json:"area_id"`
	Image string `json:"image"`
	Url string `json:"url"`
}