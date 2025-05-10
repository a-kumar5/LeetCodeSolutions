package main

import "fmt"

func main() {
	nums := []int{1, 3, 2, 1}
	fmt.Println(ContainsDuplicate(nums))
}

func ContainsDuplicate(nums []int) bool {
	mp := map[int]bool{}
	for _, num := range nums {
		if mp[num] {
			return true
		}
		mp[num] = true
		//fmt.Println(mp)
	}
	return false
}
