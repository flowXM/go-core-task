package main

import (
	"reflect"
	"testing"
)

func TestExcept(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	actual := except(slice1, slice2)
	excepted := []string{"apple", "cherry", "43", "lead", "gno1"}

	if !reflect.DeepEqual(actual, excepted) {
		t.Errorf("except actual is %v, but excepted is %v", actual, excepted)
	}
}
