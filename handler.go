package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func newHandler() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/1/queries/count/{date}", countHandler)
	r.HandleFunc("/1/queries/popular/{date}", popularHandler)

	return r
}

func popularHandler(w http.ResponseWriter, r *http.Request) {}

func countHandler(w http.ResponseWriter, r *http.Request) {}
