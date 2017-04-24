package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	// open the file
	file, err := os.Open("README.md")
	if err != nil {
		log.Fatal("Can not open the file")
	}
	defer file.Close()

	// make reader (for line by line)
	reader := bufio.NewScanner(file)

	// pattern matcher
	r := regexp.MustCompile("^##")

	for reader.Scan() {
		line := reader.Text()
		if r.MatchString(line) {
			raw := strings.Replace(line, "## ", "", 1)
			fmt.Println(raw)
		}
	}

}
