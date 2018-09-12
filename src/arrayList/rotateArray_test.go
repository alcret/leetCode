package arrayList

import (
	"log"
	"testing"
)

/*
输入: [-1,-100,3,99] 和 k = 2
输出: [3,99,-1,-100]
解释:

*/
func rotate(nums []int, k int) {
	//nums[0 : len(nums)-k] //前部分
	//nums[len(nums)-k:]    //后部分
	for i := 0; i < (len(nums)-k)/2; i++ {
		nums[i], nums[len(nums)-k-i-1] = nums[len(nums)-k-i-1], nums[i]
	}
	log.Println(nums)
	for i := 0; i < k/2; i++ {
		nums[len(nums)-k-i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[len(nums)-k-i]
	}
	log.Println(nums)
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
	log.Println(nums)
}

//测试
func TestArray(t *testing.T) {
	rotate([]int{-1, -2}, 2)
}
