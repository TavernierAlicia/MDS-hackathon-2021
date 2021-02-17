package main

// try/catch fix structs
type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

type Exception interface{}


// Create data struct for houses previews
type Preview struct {
	Records []struct{
		Id string `json:"recordid"`
		House struct {
			Pic string `json:"photo"`
			Title string `json:"title"`
			Zone string `json:"zone"`
		} `json:"fields"`
	} `json:"records"`
}

// Create data struct for houses details
type FullInfos struct {
	Records []struct{
		Id string `json:"recordid"`
		House struct {
			Pic string `json:"photo"`
			Title string `json:"title"`
			Zone string `json:"zone"`
			Org string `json:"short_desc"`
			Addr string `json:"address"`
			Url string `json:"appointmenturl"`
			Loc [2]string `json:"geolocalisation"`
			Num string `json:"phone"`
			Info string `json:"contact"`
			Details string `json:"services"`
			Desc string `json:"infos"`
		} `json:"fields"`
	} `json:"records"`
}


// Article structs
type ArticleInfos struct {
	Id int `db:"id"`
	Title string `db:"title"`
	MainPicture string `db:"main_picture"`
	Prev string `db:"text"`
	Relevance float64 `db:"relevance"`
}


type ArticleContent struct {
	Id int `db:"id"`
	LowTitle string `db:"low_title"`
	Text string `db:"text"`
	Picture string `db:"picture"`
}

// Create FAQ struct
type FAQ struct {
	Id int `db:"id"`
	Question string `db:"question"`
	Answer string `db:"answer"`
}
