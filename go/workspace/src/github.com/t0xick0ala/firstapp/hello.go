package main

import "fmt"

var (
	actorname  string = "Elisabeth Sladen"
	companion  string = "Sarah Jane Smith"
	doctorName int    = 3
	season     int    = 11
)

var (
	counter int = 0
)

func main() {

	var i float32 = 42
	var j int
	j = 27
	k := 99
	//fmt.Printf("hello, world\n")
	fmt.Println(i, j, k)
	fmt.Printf("%v, %T\n", j, i) // printf: print format, %v - value, %T - type
	//fmt.Printf("%v, %T", l, l)
}
