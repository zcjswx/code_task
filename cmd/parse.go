package main

import (
	"log"
	"os"
	"tucows_test/app"
)

func main() {
	url := os.Args[1]
	err := app.ProcessGraph(url)
	if err != nil {
		log.Fatal(err)
	}
}
