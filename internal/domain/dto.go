package domain

type MovieDTO struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	Type             string `json:"type"`
	TypeNumber       int    `json:"typeNumber"`
	Year             int    `json:"year"`
	Description      string `json:"description"`
	ShortDescription string `json:"shortDescription"`
	Slogan           string `json:"slogan"`
	Status           string `json:"status"`
	Rating           struct {
		Kp                 float64 `json:"kp"`
		Imdb               float64 `json:"imdb"`
		Tmdb               float64 `json:"tmdb"`
		FilmCritics        float64 `json:"filmCritics"`
		RussianFilmCritics float64 `json:"russianFilmCritics"`
		Await              float64 `json:"await"`
	} `json:"rating"`
	MovieLength int `json:"movieLength"`
	Persons     []struct {
		Id           int    `json:"id"`
		Photo        string `json:"photo"`
		Name         string `json:"name"`
		EnName       string `json:"enName"`
		Description  string `json:"description"`
		Profession   string `json:"profession"`
		EnProfession string `json:"enProfession"`
	} `json:"persons"`
	Budget struct {
		Value    int    `json:"value"`
		Currency string `json:"currency"`
	} `json:"budget"`
	Fees struct {
		World struct {
			Value    int    `json:"value"`
			Currency string `json:"currency"`
		} `json:"world"`
	} `json:"fees"`
}

type PersonDTO struct {
	Id     int    `json:"id"`
	Name   string `json:"name,omitempty"`
	EnName string `json:"enName,omitempty"`
	Growth int    `json:"growth,omitempty"`
	Age    int    `json:"age,omitempty"`
	Movies []struct {
		Id int `json:"id"`
	} `json:"movies,omitempty"`
}
