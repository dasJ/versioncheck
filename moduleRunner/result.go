package moduleRunner

import "github.com/dasJ/versioncheck/config"

// RunnerResult is the result of a runner run after all modules were called.
type RunnerResult struct {
	Failed  []string
	Changed []config.UpstreamConfig
}
