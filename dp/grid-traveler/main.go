package main

import (
	"fmt"

	"github.com/sosalejandro/algo-practice/dp/grid-traveler/domain"
)

func main() {

	printResult(
		domain.GridTraveler(1, 1),
		domain.GridTraveler(2, 3),
		domain.GridTraveler(3, 2),
		domain.GridTraveler(3, 3),
		domain.GridTraveler(18, 18))
}

func printResult(results ...int) {
	for _, r := range results {
		fmt.Println(r)
	}
}
