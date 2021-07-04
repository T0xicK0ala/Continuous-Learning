package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j float32
	j = float32(i)
	fmt.Printf("%v, %T\n", j, j)

	var k string
	k = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", k, k)

	// primitives
	var n bool = true
	fmt.Printf("%v, %T\n", n, n)

	n1 := 1 == 1
	n2 := 1 == 2
	fmt.Printf("%v, %T\n", n1, n1)
	fmt.Printf("%v, %T\n", n2, n2)

	var n3 bool
	fmt.Printf("%v, %T\n", n3, n3) // zero value is false, every primitive has a zero value

	a := 10             // 1010
	b := 3              // 0011
	fmt.Println(a & b)  // and 0010
	fmt.Println(a | b)  // or 1011
	fmt.Println(a ^ b)  // exclusive OR 1001
	fmt.Println(a &^ b) // and not (opposite to OR) 0100

	// bit shifting
	a1 := 8              // 2^3
	fmt.Println(a1 << 3) // 2^3 * 2^3 = 2^6 = 64
	fmt.Println(a1 >> 3) // 2^3 / 2^3 = 2^0 = 1

	// floating point numbers
	n11 := 3.14 // var n float32 or float64
	n11 = 13.7e72
	n11 = 2.1e14
	fmt.Printf("%v, %T\n", n11, n11)

	// complex number
	var nc complex64 = 1 + 2i //var nc complex64 = complex(5, 12)
	fmt.Printf("%v, %T\n", nc, nc)
	fmt.Printf("%v, %T\n", real(nc), real(nc))
	fmt.Printf("%v, %T\n", imag(nc), imag(nc))

	// string - immutable
	s := "this is a string"
	s2 := "this is also a string"
	bs := []byte(s)
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", string(s[2]), s[2]) // letter "i"
	fmt.Printf("%v, %T\n", s+s2, s+s2)
	fmt.Printf("%v, %T\n", bs, bs)

	r := 'a' // var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)
}
