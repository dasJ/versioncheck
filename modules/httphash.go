package modules

import (
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"io"
	"net/http"
)

type HttpHash struct{}

func (hh HttpHash) Init() {}

func (hh HttpHash) Name() string {
	return "httphash"
}

func (hh HttpHash) Run(upstream string, params interface{}) (version string, err error) {
	config := params.(map[string]interface{})
	url, ok := config["url"]
	if !ok {
		err = errors.New("URL is missing")
		return
	}

	sha := sha512.New()
	resp, err := http.Get(url.(string))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	_, err = io.Copy(sha, resp.Body)
	if err != nil {
		return
	}
	version = base64.URLEncoding.EncodeToString(sha.Sum(nil))
	return

}
