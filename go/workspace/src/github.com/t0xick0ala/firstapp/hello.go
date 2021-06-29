package main

import "fmt"

var l int = 55 // global level variable

func main() {

	var i float32 = 42
	var j int
	j = 27
	k := 99
	//fmt.Printf("hello, world\n")
	fmt.Println(i, j, k)
	fmt.Printf("%v, %T\n", j, i) // printf: print format, %v - value, %T - type
	fmt.Printf("%v, %T", l, l)
}
