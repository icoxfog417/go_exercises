package main

import (
	"flag"
	"fmt"
)

func main() {
	// maximum count
	const maxCount = 10

	// define argument
	var count = flag.Int("count", 3, "number of iterations")
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Println("YOU NEED TO SET COUNT!")
	} else if *count <= maxCount && *count > 0 {
		for i := 0; i < *count; i++ {
			fmt.Println("GOGO!")
		}
	} else if *count <= 0 {
		fmt.Println("COUNT SHOULD BE GREATER THAN 0.")
	} else {
		// please see below to check the format
		// https://golang.org/pkg/fmt/
		fmt.Println(fmt.Sprintf("%d is TOO MATCH!", *count))
	}

}
