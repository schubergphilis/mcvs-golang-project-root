package main

import (
	"github.com/schubergphilis/mcvs-golang-project-root/pkg/projectroot"
	log "github.com/sirupsen/logrus"
)

func main() {
	projectRoot, err := projectroot.FindProjectRoot()
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Fatal("unable to find project root")
	}
	log.Infof("project root found at: %s", projectRoot)
}
