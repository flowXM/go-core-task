package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"
	"strings"
)

func createVars() []interface{} {
	var (
		i8, i10, i16 int       = 012, 10, 0xa
		f            float64   = 10.5
		s            string    = "hello"
		b            bool      = true
		c            complex64 = 1 + 2i
	)

	return []interface{}{i8, i10, i16, f, s, b, c}
}

func getTypes(vars []interface{}) []reflect.Type {
	types := make([]reflect.Type, len(vars))

	for i, t := range vars {
		types[i] = reflect.TypeOf(t)
	}

	return types
}

func mergeVars(vars []interface{}) string {
	var strBuilder strings.Builder

	for _, t := range vars {
		switch t := t.(type) {
		case int:
			strBuilder.WriteString(fmt.Sprintf("%d", t))
		case float64:
			strBuilder.WriteString(fmt.Sprintf("%g", t))
		case string:
			strBuilder.WriteString(t)
		case bool:
			strBuilder.WriteString(fmt.Sprintf("%t", t))
		case complex64:
			strBuilder.WriteString(fmt.Sprintf("%g", t))
		}
	}

	return strBuilder.String()
}

func sha(runes []rune) string {
	str := string(runes)
	saltStr := fmt.Sprintf("%sgo-2024%s", str[0:len(str)/2], str[len(str)/2:])
	sum := sha256.Sum256([]byte(saltStr))
	return fmt.Sprintf("%x", sum)
}

func main() {
	vars := createVars()
	types := getTypes(vars)
	fmt.Println(types)
	str := mergeVars(vars)
	fmt.Println(str)
	runes := []rune(str)
	fmt.Println(runes)
	sha := sha(runes)
	fmt.Println(sha)
}
