package main

import (
	"fmt"
)

func main() {
	a := make([]int, 0)
	b := new([]int)
	
	c := []int{1,2,3,4}
	
	a = append(a, c...)
	*b = append(*b, c...)
	
	fmt.Println(a)
	fmt.Println(b)
	
	fmt.Println(a[0])
	fmt.Println((*b)[0])
}
