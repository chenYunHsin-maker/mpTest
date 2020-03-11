package main

import (
	"flag"
	"log"

	injecttag "github.com/jesseai/protoc-go-inject-tag"
)

func main() {
	var inputFile string

	flag.StringVar(&inputFile, "input", "", "path to input file")

	flag.Parse()

	if len(inputFile) == 0 {
		log.Fatal("input file is mandatory")
	}

	areas, err := injecttag.ParseFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	if err = injecttag.WriteFile(inputFile, areas); err != nil {
		log.Fatal(err)
	}
}
