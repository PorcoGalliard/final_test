package main

import (
	"fmt"
	"strings"
)

func main() {
	arr := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(soalPertama(arr, 12))

	str := "kaSur iNi ruSak"
	fmt.Println(soalKedua(str))

	str2 := "xhixhix"
	str3 := "mic"
	str4 := "haha"
	str5 := "xxxxyz"

	fmt.Println(soalKetiga(str2, "x"))
	fmt.Println(soalKetiga(str2, "hi"))
	fmt.Println(soalKetiga(str3, "mic"))
	fmt.Println(soalKetiga(str4, "ho"))
	fmt.Println(soalKetiga(str5, "xx"))

	nums1 := []int{1,3,2,4}
	nums2 := []int{1,4,8,9}
	nums3 := []int{9,9,9,9}
	fmt.Println(soalKeempat(nums1))
	fmt.Println(soalKeempat(nums2))
	fmt.Println(soalKeempat(nums3))

	soalKelima(7, 1)
	soalKelima(7, 5)

}

// Saya rasa ini soal two sum, tapi di contoh sepertinya case pertama salah, harusnya 7 dan disitu 9
// Jadi saya kerjakan ini sebagai soal TwoSum

func soalPertama(nums []int, target int) []int {
	result := []int{}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{nums[i], nums[j]}
			}
		}
	}

	return result
}

func soalKedua(str string) bool {
	n := len(str)
	if n == 0 {
		return false
	}

	loweredStr := strings.ToLower(str)
	left := 0
	right := n - 1

	for left < right {
		if loweredStr[left] != loweredStr[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func soalKetiga(str1, str2 string) int {
	result := 0

	if str2 == "" {
		return result
	}

	for i := 0; i <= len(str1)-len(str2); {
		if str1[i:i+len(str2)] == str2 {
			result++
			i += len(str2)
		} else {
			i++
		}
	}

	return result
}

func soalKeempat(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}

		digits[i] = 0
	}

	result := make([]int, len(digits)+1)
	result[0] = 1

	return result
}

func soalKelima(n, start int) {
	for i := 1; i <= n; i++ {
		current := start + (i - 1)
		for j := 0; j < i; j++ {
			fmt.Print(current + j, " ")
		}
		fmt.Println()
	}
}