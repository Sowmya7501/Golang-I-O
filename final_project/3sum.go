func threeSum(nums []int) [][]int {
    var a [][]int                                   //Team 2:
    n := len(nums)                                  //SRN :     PES1UG19CS497
    sort.Ints(nums)                                 //Name :    Sowmya M

	for i:=0;i<(n-2);i++ {
		if (i>0) && (nums[i]==nums[i-1]) {
			continue;
		}
		l := i+1;
		r := n-1;

		for l<r {
			sum := nums[i]+nums[l]+nums[r]
			if sum<0 {
				l++;
			} else if sum>0 {
				r--;
			} else {
				a=append(a,[]int{nums[i],nums[l],nums[r]})

				for (nums[l]==nums[l+1]) {
					l++;
				}
				for (nums[r]==nums[r-1]) {
					r--;
				}
				l++;
				r--;
			}
		}
}
return a
}
