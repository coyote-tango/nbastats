package datamodel

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Resource struct {
	Resource   string         `json:"resource"`
	Parameters map[string]int `json:"parameters"`
	ResultSets []ResultSet    `json:"resultSets"`
}

type ResultSet struct {
	Name    string          `json:"name"`
	Headers []string        `json:"headers"`
	RowSet  [][]interface{} `json:"rowSet"`
}

func getResource(resourceName string, params map[string]string) (*Resource, error) {
	req, err := NewRequest("GET", resourceName, params)
	client := NewHTTPClient()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	var resource *Resource

	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &resource)
	if err != nil {
		return nil, err
	}

	return resource, err
}

func NewHTTPClient() *http.Client {
	return &http.Client{}
}

func NewRequest(method, resourceName string, params map[string]string) (*http.Request, error) {
	url := "https://stats.nba.com/stats/" + resourceName
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:72.0) Gecko/20100101 Firefox/72.0")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Referer", "https://stats.nba.com/")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")

	q := req.URL.Query()
	for key := range params {
		q.Add(key, params[key])
	}
	req.URL.RawQuery = q.Encode()

	return req, nil
}
