package leetcode

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int, len(nums))
	for idx1, num := range nums {
		idx2, ok := hashmap[target-num]
		if ok {
			return []int{idx2, idx1}
		}
		hashmap[num] = idx1
	}
	return nil
}
