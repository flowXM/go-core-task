package main

import (
	"fmt"
	"math/rand/v2"
)

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func originalSlice() []int {
	var numbers = make([]int, 10)
	for i := 0; i < 10; i++ {
		numbers[i] = randRange(1, 100)
	}

	return numbers
}

func sliceExample(origin []int) []int {
	var results []int
	for i := 0; i < 10; i++ {
		if origin[i]%2 == 0 {
			results = append(results, origin[i])
		}
	}

	return results
}

func addElements(origin []int, number int) []int {
	return append(origin, number)
}

func copySlice(origin []int) []int {
	result := make([]int, len(origin))
	_ = copy(result, origin)
	return result
}

func removeElement(origin []int, index int) []int {
	if index < 0 || index >= len(origin) {
		return origin
	}

	result := make([]int, len(origin)-1)

	for i, j := 0, 0; i < len(origin); i++ {
		if i != index {
			result[j] = origin[i]
			j++
		}
	}

	return result
}

func main() {
	numbers := originalSlice()
	fmt.Printf("Original Slice: %v\n", numbers)

	slice := sliceExample(numbers)
	fmt.Printf("Slice: %v\n", slice)

	addElementsSlice := addElements(numbers, 10)
	fmt.Printf("Added number slice: %v\n", addElementsSlice)

	copiedSlice := copySlice(numbers)
	fmt.Printf("Copy slice: %v\n", copiedSlice)

	removeSlice := removeElement(numbers, 2)
	fmt.Printf("Copy slice: %v\n", removeSlice)
}
