package main

import "fmt"

type Number interface {
	int | float32 | float64
}

func main() {
	minInt := []int{10, 2, 3, 4, 5}
	fmt.Println("Min number : ", Min(minInt))

	minFloat := []float64{100, 2.2, 3.3, 4.4, 5.5}
	fmt.Println("Min number (float): ", Min(minFloat))
}

func Min[T Number](nums []T) T {
	if len(nums) == 0 {
		panic("No numbers provided")
	}

	min := nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return min
}
