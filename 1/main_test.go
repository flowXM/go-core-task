package main

import (
	"reflect"
	"testing"
)

var vars = []interface{}{012, 10, 0xa, 10.5, "hello", true, complex64(1 + 2i)}
var types = []reflect.Type{typeOf[int](), typeOf[int](), typeOf[int](), typeOf[float64](), typeOf[string](), typeOf[bool](), typeOf[complex64]()}

// Generic-функция для получения типа
func typeOf[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}

func TestCreateVars(t *testing.T) {
	actual := createVars()

	if !reflect.DeepEqual(actual, vars) {
		t.Errorf("Incorrect variables created")
	}
}
func TestGetTypes(t *testing.T) {
	actual := getTypes(vars)
	if !reflect.DeepEqual(actual, types) {
		t.Errorf("Incorrect types")
	}
}

func TestMergeVars(t *testing.T) {
	actual := mergeVars(vars)
	expected := "10101010.5hellotrue(1+2i)"

	if actual != expected {
		t.Errorf("Incorrect types")
	}
}

func TestSHA(t *testing.T) {
	actual := sha([]rune("10101010.5hellotrue(1+2i)"))
	expected := "9e6c20dc69376eb34d431a6f73b265836d39297bfc8ca3cbd7488062e9518ecc"

	if actual != expected {
		t.Errorf("Incorrect types")
	}
}
