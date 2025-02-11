package main

import (
	"fmt"
)

func main() {
	// 1. Create a Slice with Default Values
	s := []int{1, 2, 3, 4}
	fmt.Println("Original Slice:", s)

	// 2. Slice from an Array
	arr := [5]int{10, 20, 30, 40, 50}
	sliced := arr[1:4]
	fmt.Println("Sliced from array:", sliced) // Output: [20 30 40]

	// 3. Append to a Slice
	s = append(s, 5, 6)
	fmt.Println("After append:", s)

	// 4. Append One Slice to Another
	s2 := []int{7, 8, 9}
	s = append(s, s2...) // Using ... to unpack elements
	fmt.Println("After appending another slice:", s)

	// 5. Copy a Slice
	copySlice := make([]int, len(s))
	copy(copySlice, s)
	fmt.Println("Copied Slice:", copySlice)

	// 6. Delete Element at Index 2
	i := 2
	s = append(s[:i], s[i+1:]...)
	fmt.Println("After deleting index 2:", s)

	// 7. Truncate a Slice
	s = s[:5]
	fmt.Println("After truncating to first 5 elements:", s)

	// 8. Check if a Slice is Empty
	var emptySlice []int
	if len(emptySlice) == 0 {
		fmt.Println("Slice is empty")
	}

	// 9. Use nil Instead of an Empty Slice
	var nilSlice []int
	fmt.Println("Is nilSlice nil?", nilSlice == nil)
	nilSlice = []int{1, 2, 3, 4}
	fmt.Println("Is nilSlice nil?", nilSlice == nil)
	// 10. Efficiently Reset a Slice (without reallocating memory)
	s = s[:0]
	fmt.Println("Slice after reset:", s)

	// 11. Remove First Element (Avoid Shifting)
	s = []int{1, 2, 3, 4, 5}
	s = s[1:]
	fmt.Println("After removing first element:", s)

	// 12. Remove Last Element
	s = s[:len(s)-1]
	fmt.Println("After removing last element:", s)

	// 13. Preallocate a Slice to Avoid Reallocation
	preAllocated := make([]int, 0, 10) // Capacity 10, length 0
	fmt.Println("Preallocated slice - len:", len(preAllocated), "cap:", cap(preAllocated))

	// 14. Reverse a Slice
	s = []int{1, 2, 3, 4, 5}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println("Reversed slice:", s)
}
