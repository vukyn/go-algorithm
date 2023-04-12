package sort

// SelectSort - (For Asc) Find the highest value, swap place with last index and keep continue.
func SelectSort(nums []int) []int {
	lg := len(nums)
	for lg != 0 {
		tmp := 0
		for i := 0; i < lg; i++ {
			// get the highest value index
			if nums[tmp] < nums[i] {
				tmp = i
			}
			// swap the values
			if i == (lg - 1) {
				nums[lg-1], nums[tmp] = nums[tmp], nums[lg-1]
				lg -= 1
			}
		}
	}
	return nums
}
