package leetcode

import (
	"fmt"
	"testing"
)

func Test_TwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	ret := twoSum(nums, target)
	fmt.Println(ret)
}
