package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/vpalmisano/svfs/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
