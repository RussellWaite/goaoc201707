package main

import (
	"fmt"
	"log"
)

func main() {
	p1result, err := SolvePart1("./input")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("PART 1 : %s\n", p1result)

	p2result, err := SolvePart2("./input", p1result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("PART 2 : %d\n", p2result)
}
