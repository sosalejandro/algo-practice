package domain

var memo = make(map[int]bool)

func CanSum(target int, arr []int) bool {
	if target == 0 {
		return true
	}
	if target < 0 {
		return false
	}

	for i := range arr {
		r := target - arr[i]
		memo[target] = CanSum(r, arr)

		if memo[target] {
			return memo[target]
		}
	}

	return false
}
