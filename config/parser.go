package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Reads the configuration file from Argv[1] and parses it.
func Read() (ret VersioncheckConfig, err error) {
	// Default configuration
	ret = VersioncheckConfig{
		DbLocation: "/var/db/versioncheck/db.json",
	}
	// Parse configuration
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s [configfile]\n", os.Args[0])
		os.Exit(2)
	}
	configFile, err := os.Open(os.Args[1])
	if err != nil {
		return
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&ret)
	if err != nil {
		return
	}

	return
}
