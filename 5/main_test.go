package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetIntersection(t *testing.T) {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	actualState, actualData := getIntersection(a, b)
	expectedState, expectedData := true, []int{64, 3}

	if actualState != expectedState {
		t.Errorf("getIntersection returned wrong result: got %v want %v", actualState, expectedState)
	}

	if !reflect.DeepEqual(actualData, expectedData) {
		t.Errorf("getIntersection returned wrong result: got %v want %v", actualData, expectedData)
	}

	fmt.Println(getIntersection(a, b))
}
