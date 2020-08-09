package num_identical

func numIdenticalPairs(nums []int) int {
	ret := 0
	valueMap := make(map[int]int)
	for _, val := range nums {
		if v, ok := valueMap[val]; ok == true {
			ret += v
		}
		valueMap[val] += 1
	}
	return ret
}
