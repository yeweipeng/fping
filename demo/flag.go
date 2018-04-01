package main

import (
	// "fmt"
	"flag"
)

func initFlag () {
	var file_name string
	var exist bool
	flag.StringVar(&file_name, "f", "", "file what include ip")
	flag.BoolVar(&exist, "x", true, "file what include ip")
	flag.Parse()
}

func main () {
	initFlag()
}