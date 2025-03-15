package main

import (
	"reflect"
	"testing"
)

func TestNewStringIntMap(t *testing.T) {
	m := NewStringIntMap()
	if m == nil || m.data == nil {
		t.Errorf("Incorrect map initialized")
	}
}

func TestStringIntMap_Add(t *testing.T) {
	m := NewStringIntMap()
	m.Add("key", 42)

	_, ok := m.data["key"]

	if !ok {
		t.Errorf("Incorrect adding to map")
	}
}

func TestStringIntMap_Get(t *testing.T) {
	m := NewStringIntMap()
	m.data["key"] = 42

	if m.Get("key") != 42 {
		t.Errorf("Incorrect get map item")
	}
}

func TestStringIntMap_Remove(t *testing.T) {
	m := NewStringIntMap()
	m.data["key"] = 42
	m.Remove("key")

	_, ok := m.data["key"]

	if ok {
		t.Errorf("Incorrect remove from map %d", m.data["key"])
	}
}

func TestStringIntMap_Copy(t *testing.T) {
	m := NewStringIntMap()
	m.data["key"] = 42
	c := m.Copy()
	if !reflect.DeepEqual(m.data, c) {
		t.Errorf("Incorrect copy of map")
	}
}

func TestStringIntMap_Exists(t *testing.T) {
	m := NewStringIntMap()
	m.data["key"] = 42
	if !m.Exists("key") {
		t.Errorf("Incorrect exists of map")
	}
}
