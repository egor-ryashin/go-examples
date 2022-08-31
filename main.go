package main

import "os"

func main() {
	o := os.Args[1]
	if o == "parallel" {
		parallel()
	} else if o == "walk" {
		walk()
	}
}
