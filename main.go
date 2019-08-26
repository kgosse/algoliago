package main

import "flag"

func main() {
	var (
		port    string
		logfile string
	)
	flag.StringVar(&port, "p", "8080", "server port")
	flag.StringVar(&logfile, "l", "hn_logs.tsv", "tsv logfile path")
	flag.Parse()
	populateDB(logfile)
	serve(port)
}
