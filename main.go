package main

import (
	"os"
	"time"

	"github.com/gaetschwartz/devcleaner-go/internal/utils/apps"
	"github.com/gaetschwartz/devcleaner-go/internal/utils/io"
	"github.com/gaetschwartz/devcleaner-go/internal/utils/log"
	"github.com/gaetschwartz/devcleaner-go/internal/utils/path"
)

func main() {
	l := log.NewFromEnv()

	l.Debug("Getting manifest...")
	timer := time.Now()
	manifest, err := apps.GetManifest(l)
	took := time.Since(timer)
	if err != nil {
		l.Error("Error fetching manifest: %s", err)
		os.Exit(1)
	}
	l.Debug("Manifest fetched (took %s)", took)
	ctx := path.NewPathContext()
	var total int64
	for _, app := range manifest.Apps {
		l.Debug("  Evaluating app %s", app.Name)
		path, err := app.Path.Eval(ctx)
		if err != nil {
			l.Debug("  Skipping app %s: %s", app.Name, err)
			continue
		}
		l.Info("  Found %s at %s", app.Name, path)
		for _, cache := range app.Caches {
			l.Debug("    Evaluating cache %s", cache)
			path, err := cache.Eval(ctx)
			if err != nil {
				l.Debug("    Skipping cache %s: %s", cache, err)
				continue
			}
			l.Info("    Found cache path %s", path)
			size, err := io.DiskUsage(path)
			total += size
			if err != nil {
				l.Error("Error calculating disk usage: %s", err)
				os.Exit(1)
			}
			l.Debug("    Cache %s takes %d bytes", path, size)
			l.Info("    Cache %s takes %s", path, io.HumanizeBytes(size))
		}
	}

	l.Info("Total disk usage: %s", io.HumanizeBytes(total))
}
