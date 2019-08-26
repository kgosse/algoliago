package main

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func serve(p string) {
	if p == "" {
		p = "8080"
	}
	srv := &http.Server{
		Handler:      newHandler(),
		Addr:         "127.0.0.1:" + p,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Infof("server running on 127.0.0.1:%s\n", p)
	log.Fatal(srv.ListenAndServe())
}
