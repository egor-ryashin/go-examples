package main

import "os"

func main() {
	o := os.Args[1]
	if o == "parallel" {
		parallel()
	} else if o == "walk" {
		walk()
	} else if o == "limit" {
		limit()
	} else if o == "geo" {
		geo()
	} else if o == "json" {
		json_test()
	}

}
