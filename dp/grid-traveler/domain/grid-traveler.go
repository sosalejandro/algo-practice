package domain

var memo = make(map[int]map[int]int)

func GridTraveler(x, y int) int {

	if val, ok := memo[x][y]; ok {
		return val
	}

	if x == 0 || y == 0 {
		return 0
	}

	if x == 1 && y == 1 {
		return 1
	}

	if _, ok := memo[x]; !ok {
		memo[x] = make(map[int]int)
	}

	memo[x][y] = GridTraveler(x-1, y) + GridTraveler(x, y-1)
	return memo[x][y]
}
