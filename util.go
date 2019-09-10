package main

import (
	"sort"

	log "github.com/sirupsen/logrus"
)

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func sortURL(u []int, m map[string]int) {
	sort.Slice(u, func(i, j int) bool {
		iv := m[iurls[u[i]]]
		jv := m[iurls[u[j]]]
		return iv > jv
	})
}

func updateDisctinctCount(dc *int, m map[string]int, su *[]int, u string) {
	v, ok := m[u]
	m[u] = v + 1
	if !ok {
		*dc++
		*su = append(*su, urls[u])
	}
}
