package main

import (
	"reflect"
	"testing"
)

var origSlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func TestOriginalSlice(t *testing.T) {
	actual := originalSlice()
	if len(actual) != 10 {
		t.Errorf("Actual length is: %v", actual)
	}
}

func TestSliceExample(t *testing.T) {
	actual := sliceExample(origSlice)
	expected := []int{2, 4, 6, 8, 10}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Incorrect: %v", actual)
	}
}

func TestAddElements(t *testing.T) {
	actual := addElements(origSlice, 11)
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Incorrect: %v", actual)
	}
}

func TestCopySlice(t *testing.T) {
	actual := copySlice(origSlice)
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Incorrect: %v", actual)
	}

	actual[0] = 42
	expectedOrig := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if reflect.DeepEqual(actual, expectedOrig) {
		t.Errorf("Incorrect: %v", actual)
	}
}

func TestRemoveElement(t *testing.T) {
	actual := removeElement(origSlice, 0)
	expected := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Incorrect: %v", actual)
	}
}
