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
