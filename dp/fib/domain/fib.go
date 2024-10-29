package fib

var memo = make(map[int]int)

func Fib(n int) int {
	if val, ok := memo[n]; ok {
		return val
	}

	if n <= 1 {
		return n
	}

	memo[n] = Fib(n-1) + Fib(n-2)

	return memo[n]
}
