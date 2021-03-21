func maxArea(height []int) int {
	res := 0
	l := 0
	r := len(height) - 1
	for l < r {
		h := height[l]
		if height[l] > height[r] {
			h = height[r]
		}
		if res < h*(r-l) {
			res = h * (r - l)
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}