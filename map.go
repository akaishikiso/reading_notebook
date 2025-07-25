package main

import (
	"fmt"
	"maps"
)

func MapExample() {
	var m = make(map[string]int)
	m["key1"] = 1
	m["key2"] = 2

	fmt.Println(m) // map[key1:1 key2:2]

	n := map[string]int{
		"keyA": 10,
		"keyB": 20,
	}

	maps.Copy(m, n)
	fmt.Println(m) // map[key1:1 key2:2 keyA:10 keyB:20]

	v, ok := m["key1"]
	fmt.Println(v, ok) // 1 true

	v, ok = m["keyC"]
	fmt.Println(v, ok) // 0 false

	delete(m, "key1")
	fmt.Println(m) // map[key2:2 keyA:10 keyB:20]
}