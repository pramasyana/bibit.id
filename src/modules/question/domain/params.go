package domain

type ResponseNumber1 struct {
	Database string `json:"database"`
	Query    string `json:"query"`
}

type ResponseNumber3 struct {
	Text   string `json:"text"`
	Result string `json:"result"`
}

type ParamsNumber3 struct {
	Text string `json:"text"`
}

type ResponseNumber4 struct {
	Text   []string    `json:"text"`
	Result interface{} `json:"result"`
}

type ParamsNumber4 struct {
	Text []string `json:"text"`
}

type ResponseNumber2List struct {
	Search   []ResponseNumber2ListSerach `json:"Search,omitempty"`
	Response string                      `json:"Response,omitempty"`
	Error    string                      `json:"Error,omitempty"`
}

type ResponseNumber2ListSerach struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type ResponseNumber2Detail struct {
	Title      string    `json:"Title,omitempty"`
	Year       string    `json:"Year,omitempty"`
	Rated      string    `json:"Rated,omitempty"`
	Released   string    `json:"Released,omitempty"`
	Runtime    string    `json:"Runtime,omitempty"`
	Genre      string    `json:"Genre,omitempty"`
	Director   string    `json:"Director,omitempty"`
	Writer     string    `json:"Writer,omitempty"`
	Actors     string    `json:"Actors,omitempty"`
	Plot       string    `json:"Plot,omitempty"`
	Language   string    `json:"Language,omitempty"`
	Country    string    `json:"Country,omitempty"`
	Awards     string    `json:"Awards,omitempty"`
	Poster     string    `json:"Poster,omitempty"`
	Ratings    []Ratings `json:"Ratings,omitempty"`
	Metascore  string    `json:"Metascore,omitempty"`
	ImdbRating string    `json:"imdbRating,omitempty"`
	ImdbVotes  string    `json:"imdbVotes,omitempty"`
	ImdbID     string    `json:"imdbID,omitempty"`
	Type       string    `json:"Type,omitempty"`
	DVD        string    `json:"DVD,omitempty"`
	BoxOffice  string    `json:"BoxOffice,omitempty"`
	Production string    `json:"Production,omitempty"`
	Website    string    `json:"Website,omitempty"`
	Response   string    `json:"Response,omitempty"`
	Error      string    `json:"Error,omitempty"`
}

type Ratings struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}

type ParamsNumber2 struct {
	ID     string `json:"i,omitempty"`
	Search string `json:"s,omitempty"`
	Year   string `json:"y,omitempty"`
	Plot   string `json:"p,omitempty"`
	Title  string `json:"t,omitempty"`
	Type   string `json:"type,omitempty"`
	Page   string `json:"page,omitempty"`
}
