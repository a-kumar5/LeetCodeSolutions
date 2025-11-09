// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sCount := make(map[byte]int)
	tCount := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		sCount[s[i]]++
	}

	for j := 0; j < len(t); j++ {
		tCount[t[j]]++
	}
	for k, v := range sCount {
		if v != tCount[k] {
			return false
		}
	}
	return true
}

func twoSum(nums []int, target int) []int {
	indexNums := make(map[int]int)
	for i, v := range nums {
		if _, ok := indexNums[target-v]; ok {
			return []int{indexNums[v], i}

		}
		indexNums[v] = i
	}
	return []int{}
}

func main() {
	fmt.Println("Hello, 世界")
	//s := "racecar"
	//t := "racecar"
	nums := []int{3, 4, 5, 6}
	target := 7
	//fmt.Println(isAnagram(s, t))
	fmt.Println(twoSum(nums, target))
}
