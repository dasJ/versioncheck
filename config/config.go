package config

// UpstreamConfig is the configuration of a single upstream project.
type UpstreamConfig struct {
	Name       string      `json:"name"`
	Module     string      `json:"module"`
	Parameters interface{} `json:"params"`
	Tags       []string    `json:"tags"`
	OldVersion string      `json:"-"`
	NewVersion string      `json:"-"`
}

// VersioncheckConfig is the main configuration of this application.
// It contains paths and the list of upstreams.
type VersioncheckConfig struct {
	DbLocation  string           `json:"dbLocation"`
	Notificator string           `json:"notificator"`
	Upstreams   []UpstreamConfig `josn:"upstreams"`
}
