//TEAM 2
//NAME: S TEJASHRI
//SRN: PES1UG20ME087

func sortedSquares(nums []int) []int {
	sol := []int{}

	for _, num := range nums {
		sol = append(sol, num*num)
	}

	sort.Ints(sol)

	return sol
}

