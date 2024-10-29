package domain

var memo = make(map[int][]int)

func HowSum(target int, arr []int) []int {
	if target == 0 {
		return []int{}
	}

	if target < 0 {
		return nil
	}

	result := make([]int, 0)
	for _, n := range arr {
		r := target - n
		memo[r] = HowSum(r, arr)

		if memo[r] != nil {
			result = append(memo[r], n)
			return result
		}
	}

	return nil
}
