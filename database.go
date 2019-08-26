package main

import "time"

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

func init() {
	years = make(map[int]*year)
	urls = make(map[string]int)
	iurls = make(map[int]string)
}
