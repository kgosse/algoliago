package main

import (
	log "github.com/sirupsen/logrus"
)

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
