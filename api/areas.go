package api

type Animes struct{
	Place string `json:"place"`
	Title string `json:"title"`
}

func Areas_list() Animes{
	return Animes{
		Place: "aaa",
		Title: "bbb",
	}
}