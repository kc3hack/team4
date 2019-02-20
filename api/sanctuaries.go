package api


type Sanctuaries struct{
	Sanctuaries_list Sanctuary_detals `json:"sanctuaries_list"`
}

type Sanctuary_detals struct{
	Sanctuary_ID int `json:"sanctuary_id"`
	Sanctuary_Name string `json:"sanctuary_name"`
	Anime_Name string `json:"anime_name"`
	Area_ID int `json:"anime_id"`
	Area_Name string `json:"area_name"`
	Url string `json:"url"`
}

func GetSanctuaries()  Sanctuaries{

	return Sanctuaries{
		Sanctuaries_list: Sanctuary_detals{
			Sanctuary_ID: 1,
			Sanctuary_Name: "豊郷小学校",
			Anime_Name: "けいおん！",
			Area_ID: 1,
			Area_Name: "滋賀県犬上群豊郷町",
			Url: "hoge",
		},
	}
}