package main

import (
	"flag"
	"fmt"
)

func main() {
	var nameAges = map[string]int{
		"Andy":  20,
		"Bob":   21,
		"Cony":  21,
		"Debit": 20,
		"Elphy": 21,
		"Fox":   23,
	}

	// define argument
	var age = flag.Int("age", -1, "age of person")
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Println("YOU NEED TO SET AGE!")
	} else if *age > 0 {
		var matched []string
		for key, value := range nameAges {
			if value == *age {
				matched = append(matched, key)
			}
		}

		//print names
		fmt.Println(fmt.Sprintf("Matched Names: %q", matched))
	} else {
		fmt.Println("YOU HAVE TO SPECIFY AGE OR SET SUITABLE VALUE.")
	}

}
