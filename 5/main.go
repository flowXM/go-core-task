package main

import "fmt"

func getIntersection(a, b []int) (bool, []int) {
	m := make(map[int]bool)
	res := make([]int, 0)

	for _, v := range a {
		m[v] = true
	}

	for _, v := range b {
		if _, ok := m[v]; ok {
			res = append(res, v)
		}
	}

	return len(res) > 0, res
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	fmt.Println(getIntersection(a, b))
}
