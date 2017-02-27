package moduleRunner

import (
	"github.com/dasJ/versioncheck/modules"
)

// initModules creates a map of all modules
func initModules() map[string]versionModule {
	ret := make(map[string]versionModule)

	mods := [...]versionModule{
		modules.Github{},
	}
	for _, mod := range mods {
		mod.Init()
		ret[mod.Name()] = mod
	}

	return ret
}