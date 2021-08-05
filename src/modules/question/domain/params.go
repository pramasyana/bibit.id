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
	Search   []ResponseNumber2ListSerach `json:"Search"`
	Response string                      `json:"Response"`
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
	Title      string    `json:"Title"`
	Year       string    `json:"Year"`
	Rated      string    `json:"Rated"`
	Released   string    `json:"Released"`
	Runtime    string    `json:"Runtime"`
	Genre      string    `json:"Genre"`
	Director   string    `json:"Director"`
	Writer     string    `json:"Writer"`
	Actors     string    `json:"Actors"`
	Plot       string    `json:"Plot"`
	Language   string    `json:"Language"`
	Country    string    `json:"Country"`
	Awards     string    `json:"Awards"`
	Poster     string    `json:"Poster"`
	Ratings    []Ratings `json:"Ratings"`
	Metascore  string    `json:"Metascore"`
	ImdbRating string    `json:"imdbRating"`
	ImdbVotes  string    `json:"imdbVotes"`
	ImdbID     string    `json:"imdbID"`
	Type       string    `json:"Type"`
	DVD        string    `json:"DVD"`
	BoxOffice  string    `json:"BoxOffice"`
	Production string    `json:"Production"`
	Website    string    `json:"Website"`
	Response   string    `json:"Response"`
	Error      string    `json:"Error,omitempty"`
}

type Ratings struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}
