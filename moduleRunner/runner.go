package moduleRunner

import (
	"fmt"
	"github.com/dasJ/versioncheck/config"
	"github.com/dasJ/versioncheck/verdb"
	"os"
)

// RunModules runs all modules
// Could make this threaded, but I'm too lazy.
func RunModules(cfg config.VersioncheckConfig, verDb *verdb.Verdb) (res RunnerResult) {
	modules := initModules()

	for _, us := range cfg.Upstreams {
		mod, ok := modules[us.Module]
		if !ok {
			fmt.Fprintf(os.Stderr, "[%s] Module '%s' is not available\n", us.Name, us.Module)
			res.Failed = append(res.Failed, us.Name)
			continue
		}
		newVersion, err := mod.Run(us.Name, us.Parameters)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[%s] Module failed\n", us.Name)
			res.Failed = append(res.Failed, us.Name)
			continue
		}
		oldVersion := verDb.VersionOf(us.Name)
		if oldVersion == "" {
			fmt.Printf("[%s] -> %s\n", us.Name, newVersion)
			verDb.UpdateVersion(us.Name, newVersion)
			continue
		}

		if newVersion != oldVersion {
			fmt.Printf("[%s] %s -> %s\n", us.Name, oldVersion, newVersion)
			verDb.UpdateVersion(us.Name, newVersion)
			res.Changed = append(res.Changed, us)
		}
	}

	return
}
