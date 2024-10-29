package main

import (
	"fmt"

	"github.com/sosalejandro/algo-practice/dp/can-sum/domain"
)

func main() {

	printResults(
		domain.CanSum(6, []int{5, 3, 4, 7}),
		domain.CanSum(7, []int{2, 4}),
	)
}

func printResults(results ...bool) {
	for _, r := range results {
		fmt.Println(r)
	}
}
