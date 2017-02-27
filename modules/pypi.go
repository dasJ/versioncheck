package modules

import (
	"encoding/json"
	"errors"
	"net/http"
)

type PyPi struct{}

func (pp PyPi) Init() {}

func (pp PyPi) Name() string {
	return "pypi"
}

type pyPiInfo struct {
	Version string `json:"version"`
}

type pyPiResponse struct {
	Info pyPiInfo `json:"info"`
}

func (pp PyPi) Run(upstream string, params interface{}) (version string, err error) {
	conf := params.(map[string]interface{})
	pkg, ok := conf["package"]
	if !ok {
		err = errors.New("Package name is missing")
		return
	}
	resp, err := http.Get("https://pypi.python.org/pypi/" + pkg.(string) + "/json")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	// Decode
	jsonResp := new(pyPiResponse)
	err = json.NewDecoder(resp.Body).Decode(jsonResp)
	if err != nil {
		return
	}
	// Get package version
	if jsonResp.Info.Version == "" {
		err = errors.New("No Version")
		return
	}
	version = jsonResp.Info.Version
	return
}
