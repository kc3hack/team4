package api

type Animes struct {
	Anime_ID int `json:"anime_id"`
	Title string `json:"anime_name"`
	Anime_details Anime_details `json:"animes_details"`
}

type Anime_details struct{
	Area string `json:"area_name"`
	Area_ID int `json:"area_id"`
	Url string `json:"url"`
}

type Position struct{
	Sanctuaries string `json:"sanctuaries"`
	Lat int `json:"lat"`
	Lon int `json:"lon"`
}

type Animes_Lists struct{
	Anime_List []Animes `json:"anime_list"`
}

func GetAnimes() Animes_Lists{
	return Animes_Lists{
		Anime_List: []Animes{
			Animes{
				Anime_ID: 1,
				Title: "けいおん！",
				Anime_details: Anime_details{
					Area: "滋賀県犬上群豊郷町",
					Area_ID: 1,
					Url: "http://www.tbs.co.jp/anime/k-on/",
				},
			},
		},
	}
}	