package leetcode

import (
	"fmt"
	"sort"
)

// leetcode453
func MinMoves(nums []int) int {
	sum := 0
	min := 0xFFFFFFFF
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if nums[i] < min {
			min = nums[i]
		}
	}
	return sum - min*len(nums)
}

// leetcode454
func FourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	mp := make(map[int]int)
	for _, val1 := range nums1 {
		for _, val2 := range nums2 {
			mp[val1+val2]++
		}
	}
	cnt := 0
	for _, val1 := range nums3 {
		for _, val2 := range nums4 {
			if _, ok := mp[-(val1 + val2)]; ok {
				cnt++
			}
		}
	}
	return cnt
}

// leetcode459
func RepeatedSubstringPattern(s string) bool {
	next := make([]int, len(s)+1)
	getNext := func(str string) {
		i, j := 0, -1
		next[0] = -1
		for i < len(str) {
			if j == -1 || str[i] == str[j] {
				i++
				j++
				next[i] = j
			} else {
				j = next[j]
			}
		}
	}
	getNext(s)
	last := next[len(s)]
	if last == 0 {
		return false
	}
	return len(s)%(len(s)-last) == 0
}

// leetcode475
func FindRadius(houses []int, heaters []int) int {
	sort.Ints(heaters)
	binarySearchFirst := func(nums []int, target int) int {
		low, high := 0, len(nums)-1
		for low <= high {
			mid := (low + high) / 2
			if nums[mid] >= target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		return low
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	res := 0
	for i := 0; i < len(houses); i++ {
		r := binarySearchFirst(heaters, houses[i])
		l := r - 1
		if l < 0 {
			if houses[i] < heaters[r]-res {
				res = heaters[r] - houses[i]
			}
			continue
		}
		if r >= len(heaters) {
			if houses[i] > heaters[l]+res {
				res = houses[i] - heaters[l]
			}
			continue
		}
		if houses[i] < heaters[r]-res && houses[i] > heaters[l]+res {
			//不在右边加热器的范围      并且         不在左边加热器的范围
			res = min(houses[i]-heaters[l], heaters[r]-houses[i])
		}
	}
	return res
}

// leetcode476
func FindComplement(num int) int {
	res := []int{}
	for num > 0 {
		res = append(res, num&1)
		num >>= 1
	}
	ans := 0
	fmt.Println(res)
	for i := len(res) - 1; i >= 0; i-- {
		if res[i] == 0 {
			ans |= 1
		}
		ans <<= 1
	}
	return ans >> 1
}

// leetcode488
func FindMinStep(board string, hand string) int {
	for i := 0; i < len(board); i++ {

	}
	return 1
}

// leetcode491
func FindSubsequences(nums []int) [][]int {
	var backtrace func(start int)
	res := [][]int{}
	cur := []int{}

	backtrace = func(start int) {
		if len(cur) >= 2 {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
		}
		if start == len(nums) {
			return
		}
		history := make([]int, 201)
		for i := start; i < len(nums); i++ {
			if len(cur) > 0 && nums[i] < cur[len(cur)-1] || history[nums[i]+100] == 1 {
				continue
			}
			history[nums[i]+100] = 1
			cur = append(cur, nums[i])
			backtrace(i + 1)
			cur = cur[:len(cur)-1]
		}
	}
	backtrace(0)
	return res
}

// leetcode492
func ConstructRectangle(area int) []int {
	num1 := 1
	for i := 1; i*i <= area; i++ {
		if area%i == 0 {
			num1 = i
		}
	}
	res := []int{area / num1, num1}
	return res
}

// leetcode495
func FindPoisonedDuration(timeSeries []int, duration int) int {
	begin, end := 0, -1
	res := 0
	for _, time := range timeSeries {
		if time > end {
			res += (end - begin + 1)
			end = time + duration - 1
			begin = time
		} else {
			end = time + duration - 1
		}
	}
	res += (end - begin + 1)
	return res
}

// leetcode496
func NextGreaterElement1(nums []int) []int {
	// 给定一个数组，返回一个等长的数组，对应索引存储着下一个更大元素，如果没有更大的元素，就存 -1。
	// 输入：[2, 1, 2, 4, 3]
	// 输出：[4, 2, 4, -1, -1]

	res := make([]int, len(nums))
	stack := []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) != 0 && stack[len(stack)-1] <= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[i] = -1
		} else {
			res[i] = stack[len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return res
}

func NextGreaterElement2(nums []int) []int {
	// 给定一个数组 T = [73, 74, 75, 71, 69, 72, 76, 73]
	// 该数组存放是近几天的天气气温（华氏度）
	// 返回一个数组，计算：对于每一天，至少等多少天才能等到一个更暖和的气温；如果等不到那一天，填 0 。
	// 输入 T = [73, 74, 75, 71, 69, 72, 76, 73]
	// 输出 R = [1, 1, 4, 2, 1, 1, 0, 0]。
	// 解释：第一天 73 华氏度，第二天 74 华氏度，比 73 大，所以对于第一天，只要等一天就能等到一个更暖和的气温。后面的同理。

	res := make([]int, len(nums))
	stack := []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) != 0 && nums[stack[len(stack)-1]] <= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[i] = 0
		} else {
			index := stack[len(stack)-1]
			res[i] = index - i
		}
		stack = append(stack, i)
	}
	return res
}

func NextGreaterElement3(nums1, nums2 []int) []int {
	res := make([]int, len(nums1))
	stack := []int{}
	mp := map[int]int{}
	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]
		for len(stack) > 0 && num >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			mp[num] = stack[len(stack)-1]
		} else {
			mp[num] = -1
		}
		stack = append(stack, num)
	}
	for i, num := range nums1 {
		res[i] = mp[num]
	}
	return res
}
