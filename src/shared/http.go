package shared

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func RequestHTTP(method, url string, body io.Reader, target interface{}, headers ...map[string]string) (int, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return 0, err
	}

	// iterate optional data of headers
	for _, header := range headers {
		for key, value := range header {
			req.Header.Set(key, value)
		}
	}

	client := &http.Client{}
	r, err := client.Do(req)
	if err != nil {
		return r.StatusCode, err
	}

	defer r.Body.Close()

	resp, _ := ioutil.ReadAll(r.Body)
	if e := json.Unmarshal(resp, &target); e != nil {
		return r.StatusCode, e
	}

	return r.StatusCode, nil
}
