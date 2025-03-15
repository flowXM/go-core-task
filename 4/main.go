package main

import "fmt"

func except(s1, s2 []string) []string {
	m := make(map[string]bool)
	res := make([]string, 0)

	for _, v := range s2 {
		m[v] = true
	}

	for _, v := range s1 {
		if _, ok := m[v]; !ok {
			res = append(res, v)
		}
	}

	return res
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	res := except(slice1, slice2)
	fmt.Println(res)
}
