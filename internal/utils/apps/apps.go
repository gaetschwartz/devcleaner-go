package apps

import "github.com/gaetschwartz/devcleaner-go/internal/utils/path"

type App struct {
	Name   string             `json:"name"`
	Path   path.PathPattern   `json:"path"`
	Caches []path.PathPattern `json:"caches"`
}

var knownApps = []App{
	// homebrew
	{
		Name:   "homebrew",
		Path:   "{os==darwin:/opt/homebrew/bin/brew,linux:/home/linuxbrew/.linuxbrew/bin/brew}",
		Caches: []path.PathPattern{"{os==darwin:/opt/homebrew/Cellar,linux:/home/linuxbrew/.linuxbrew/Cellar}"},
	},
	// npm
	{
		Name:   "npm",
		Path:   "{os==darwin:/usr/local/bin/npm,linux:/usr/bin/npm}",
		Caches: []path.PathPattern{"{os==darwin:/usr/local/lib/node_modules,linux:/usr/lib/node_modules}"},
	},
	// yarn
	{
		Name:   "yarn",
		Path:   "{os==darwin:/usr/local/bin/yarn,linux:/usr/bin/yarn}",
		Caches: []path.PathPattern{"{os==darwin:/usr/local/lib/node_modules,linux:/usr/lib/node_modules}"},
	}}
