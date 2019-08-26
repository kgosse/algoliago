package main

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/valyala/tsvreader"
)

type year struct {
	id            int
	distinctCount int
	urls          map[string]int
	sortedurls    []int
	months        [12]*month
}

type month struct {
	id            int
	distinctCount int
	urls          map[string]int
	sortedurls    []int
	days          [31]*day
}

type day struct {
	id            int
	distinctCount int
	urls          map[string]int
	sortedurls    []int
	hours         [24]*hour
}

type hour struct {
	id            int
	distinctCount int
	urls          map[string]int
	sortedurls    []int
	minutes       [60]*minute
}

type minute struct {
	id            int
	distinctCount int
	urls          map[string]int
	sortedurls    []int
	seconds       [60]*second
}

type second struct {
	id            int
	distinctCount int
	urls          map[string]int
	sortedurls    []int
	records       []int
}

type record struct {
	id   int
	time time.Time
	url  string
}

var years map[int]*year
var records []*record
var urls map[string]int
var iurls map[int]string
var recordID int
var urlID int
var formats map[int]string

func init() {
	years = make(map[int]*year)
	urls = make(map[string]int)
	iurls = make(map[int]string)
	formats = make(map[int]string)
	formats[4] = "2006"
	formats[7] = "2006-01"
	formats[10] = "2006-01-02"
	formats[16] = "2006-01-02 15:04"
	formats[19] = "2006-01-02 15:04:05"
}

func populateDB(p string) {
	file, err := os.Open(p)
	handleError(err)
	defer file.Close()

	r := tsvreader.New(file)
	log.Info("populating the db...")
	for r.Next() {
		recordID++
		d := r.String()
		u := r.String()
		t, err := time.Parse(formats[19], d)
		handleError(err)
		rc := &record{
			recordID,
			t,
			u,
		}

		records = append(records, rc)
	}
	if err := r.Error(); err != nil {
		log.Fatalf("Failed to parse tsv: %s", err)
	}
	log.Infof("Done. %d new records.\n", len(records))
}
