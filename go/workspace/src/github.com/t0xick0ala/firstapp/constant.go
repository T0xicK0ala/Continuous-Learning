package main

import (
	"fmt"
)

const e int32 = 42
const (
	g = iota
	// h = iota
	// i = iota
	h
	i // pattern
) // counter, enumerated constant

const (
	error = iota // _ = iota set as default, cannot get zero value here
	// _ = iota + 5 then j equals to 6
	j
	k
	l
)

const (
	isAdmin = 1 << iota
	isHeadquaters
	canSeeFinancials

	canSeeA
	canSeeB
	canSeeC
	canSeeD
	canSeeE
)

func cons() {
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)

	const a int = 14
	const b string = "foo"
	const c float32 = 3.14
	const d bool = true
	const e int = 14 // global constant can be shadowed internally
	var f int = 27

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v %T\n", e, e)
	fmt.Printf("%v %T\n", e+f, e+f)

	fmt.Printf("%v\n", g)
	fmt.Printf("%v\n", h)
	fmt.Printf("%v\n", i)

	var conType int
	fmt.Printf("%v\n", k == conType)

	var role byte = isAdmin | canSeeFinancials | canSeeD
	fmt.Printf("%v\n", role)
	fmt.Printf("Is Admin? %v\n", isAdmin&role == isAdmin)
	fmt.Printf("Is HQ? %v\n", isHeadquaters&role == isHeadquaters)
}
