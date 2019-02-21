package api


type Sanctuaries struct{
	Sanctuaries_list Sanctuary_detals `json:"sanctuaries_list"`
}

type Sanctuary_detals struct{
	Sanc_ID int `json:"sanc_id"`
	Sanc_Name string `json:"sanc_name"`
	Anime_ID int `json:"anime_id"`
	Place_ID string `json:"place_id"`
	Image string `json:"image"`
}

// func GetSanctuaries()  Sanctuaries{

// 	return Sanctuaries{
// 		Sanctuaries_list: Sanctuary_detals{
// 			Sanctuary_ID: 1,
// 			Sanctuary_Name: "豊郷小学校",
// 			Anime_Name: "けいおん！",
// 			Area_ID: 1,
// 			Area_Name: "滋賀県犬上群豊郷町",
// 			Url: "hoge",
// 		},
	// }
// }