package twoSum

func twoSum(nums []int, target int) []int {
	result := make(map[int]int)
	for index, value := range nums {
		right := target - value
		if i, ok := result[right]; ok {
			return []int{index, i}
		}
		result[value] = index
	}
	return nil
}
