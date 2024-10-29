package main

import (
	"fmt"

	"github.com/sosalejandro/algo-practice/dp/how-sum/domain"
)

func main() {

	printResults(
		domain.HowSum(100, []int{14, 5}),
	)
}

func printResults(results ...[]int) {
	for _, r := range results {
		for _, v := range r {
			fmt.Printf("%d ", v)
		}
		fmt.Printf("\n")
	}
}
