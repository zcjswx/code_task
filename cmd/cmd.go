package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"tucows_test/app"
)

func main() {
	fmt.Println("paste your JSON queries")
	jsonStr := ""
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		inputLine := strings.TrimSpace(scanner.Text())
		jsonStr += inputLine
		// Check if the line is empty or contains only whitespace
		if strings.TrimSpace(inputLine) == "" {
			fmt.Println("Detected an empty line. Processing...")
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading standard input: %s\n", err)
	}

	fmt.Println(jsonStr, "received")
	json, err := app.Process(jsonStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(json)
}
