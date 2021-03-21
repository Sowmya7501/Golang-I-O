func sortedSquares(nums []int) []int {
	sol := []int{}

	for _, num := range nums {
		sol = append(sol, num*num)
	}

	sort.Ints(sol)

	return sol
}

