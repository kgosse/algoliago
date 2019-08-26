package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

type countResponse struct {
	Count int `json:count`
}

type popularResponse struct {
	Queries []queryItem `json:queries`
}

type queryItem struct {
	Query string `json:query`
	Count int    `json:count`
}

func newHandler() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/1/queries/count/{date}", countHandler)
	r.HandleFunc("/1/queries/popular/{date}", popularHandler)

	return r
}

func popularHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	date := params["date"]
	size, err := strconv.Atoi(r.FormValue("size"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	f, _ := formats[len(date)]
	t, err := time.Parse(f, date)
	if err != nil {
		log.Error(err)
		http.Error(w, "bad date", http.StatusInternalServerError)
		return
	}

	var resp popularResponse

	switch len(date) {
	case 4:
		resp = yearPopularCount(t, size)
		break
	case 7:
		resp = monthPopularCount(t, size)
		break
	case 10:
		resp = dayPopularCount(t, size)
		break
	case 16:
		resp = minutePopularCount(t, size)
		break
	case 19:
		resp = secondPopularCount(t, size)
		break
	default:
		http.Error(w, "bad date", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	date := params["date"]
	f, _ := formats[len(date)]
	t, err := time.Parse(f, date)
	if err != nil {
		log.Error(err)
		http.Error(w, "bad date", http.StatusInternalServerError)
		return
	}

	var count int

	switch len(date) {
	case 4:
		count = yearDistinctCount(t)
		break
	case 7:
		count = monthDistinctCount(t)
		break
	case 10:
		count = dayDistinctCount(t)
		break
	case 16:
		count = minuteDistinctCount(t)
		break
	case 19:
		count = secondDistinctCount(t)
		break
	default:
		http.Error(w, "bad date", http.StatusInternalServerError)
		return
	}

	resp := countResponse{Count: count}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
