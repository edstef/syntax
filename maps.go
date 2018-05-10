package main

import (
	"fmt"
)

func printMap(m map[string]int) {
	fmt.Println("map:", m)
}

func printValues(m map[string]int) {
	for k := range m {
		fmt.Println("val:", m[k])
	}
}

func printKeys(m map[string]int) {
	for k := range m {
		fmt.Println("key:", k)
	}
}

func printKeyVals(m map[string]int) {
	for k, v := range m {
		fmt.Println("key:", k, "val:", v)
	}
}

func countElements(m map[string]int) {
	fmt.Println("# Of elements:", len(m))
}

func main() {
	m := make(map[string]int)

	m["key1"] = 1
	m["key2"] = 2

	printMap(m)
	fmt.Println(m["key1"])
	printValues(m)
	printKeys(m)
	printKeyVals(m)
	countElements(m)
}