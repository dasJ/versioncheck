package modules

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Rubygems struct{}

func (rg Rubygems) Init() {}

func (rg Rubygems) Name() string {
	return "rubygems"
}

type rubygemsResponse struct {
	Version string `json:"version"`
}

func (rg Rubygems) Run(upstream string, params interface{}) (version string, err error) {
	conf := params.(map[string]interface{})
	gem, ok := conf["gem"]
	if !ok {
		err = errors.New("Gem name is missing")
		return
	}
	resp, err := http.Get("https://rubygems.org/api/v1/versions/" + gem.(string) + "/latest.json")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	// Decode
	jsonResp := new(rubygemsResponse)
	err = json.NewDecoder(resp.Body).Decode(jsonResp)
	if err != nil {
		return
	}
	version = jsonResp.Version
	return
}
