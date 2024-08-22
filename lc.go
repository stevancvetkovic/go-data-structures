package main

import "fmt"

func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			answer[i] = 0
		} else {
			newArr := make([]int, len(nums))
			copy(newArr, nums)
			newArr[i] = 1
			answer[i] = multiply(newArr)
		}
	}

	return answer
}

func multiply(arr []int) int {
	fmt.Println(arr)
	out := 1
	for _, num := range arr {
		if num == 0 {
			return 0
		}
		out = out * num
	}
	return out
}

func main() {
	nums := []int{-1, 1, 0, -3, 3}
	out := productExceptSelf(nums)
	fmt.Println(out)
}
