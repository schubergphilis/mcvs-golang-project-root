package main

import (
	"log/slog"
	"os"

	"github.com/schubergphilis/mcvs-golang-project-root/pkg/projectroot"
)

func main() {
	projectRoot, err := projectroot.FindProjectRoot()
	if err != nil {
		slog.Error("unable to find project root", "error", err)
		os.Exit(1)
	}

	slog.Info("project root found at: " + projectRoot)
}
