package main

import (
	"fmt"
)

func sli() {
	// slice
	a1 := []int{1, 2, 3}
	b1 := a1 // slice is reference type
	b1[1] = 5
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Printf("Length: %v\n", len(a1))
	fmt.Printf("Capacity: %v\n", cap(a1))

	// another example
	fmt.Println()
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// slice operations can be used in arrays too, see below
	// a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]   // slice of all elements
	c := a[3:]  // slcie from 4th element to end
	d := a[:6]  // slice first 6 elements
	e := a[3:6] // slice the 4th, 5th, 6th elements
	// a[5] = 42 - it will change all values
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	// make function
	a2 := make([]int, 3, 100) // it takes two/three parameters, eg. type and length, and third: capacity
	fmt.Println(a2)
	fmt.Printf("Length: %v\n", len(a2))
	fmt.Printf("Capacity: %v\n", cap(a2))
	a2 = append(a2, 1)
	fmt.Println(a2)
	fmt.Printf("Length: %v\n", len(a2))
	fmt.Printf("Capacity: %v\n", cap(a2))
	a2 = append(a2, []int{3, 4, 5, 6}...) // spread operator
	fmt.Println(a2)
	fmt.Printf("Length: %v\n", len(a2))
	fmt.Printf("Capacity: %v\n", cap(a2))

	// stack operation
	a3 := []int{1, 2, 3, 4, 5}
	b3 := a3[1:]                      // trim the first element
	b33 := a3[:len(a3)-1]             // trim the end element
	b333 := append(a3[:2], a3[3:]...) // trim the middle element
	fmt.Println(b3)
	fmt.Println(b33)
	fmt.Println(b333)
	fmt.Println(a3) // pointing to the same underlying array
}
