package modules

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Github struct{}

func (gh Github) Init() {}

func (gh Github) Name() string {
	return "github"
}

type githubResponse struct {
	TagName    string `json:"tag_name"`
	Draft      bool   `json:"draft"`
	Prerelease bool   `json:"prerelease"`
}

func (gh Github) Run(upstream string, params interface{}) (version string, err error) {
	config := params.(map[string]interface{})
	namespace, ok := config["namespace"]
	if !ok {
		err = errors.New("Namespace is missing")
		return
	}
	project, ok := config["project"]
	if !ok {
		err = errors.New("Project name is missing")
		return
	}
	prereleaseOkay, ok := config["prerelease"]
	if !ok {
		prereleaseOkay = false
	}
	draftOkay, ok := config["draft"]
	if !ok {
		draftOkay = false
	}

	resp, err := http.Get("https://api.github.com/repos/" + namespace.(string) + "/" + project.(string) + "/releases")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	// Decode
	jsonResp := new([]githubResponse)
	err = json.NewDecoder(resp.Body).Decode(jsonResp)
	if err != nil {
		return
	}
	// TODO Support paging
	// Search for newest release
	for _, ver := range *jsonResp {
		if (ver.Prerelease && prereleaseOkay.(bool)) || (ver.Draft && draftOkay.(bool)) || (!ver.Prerelease && !ver.Draft) {
			version = ver.TagName
			break
		}
	}
	return

}
