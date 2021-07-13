package main

import (
	"fmt"
)

func arr() {
	// grades := [3]int{97, 85, 93}
	grades := [...]int{97, 85, 93} // any size
	var students [3]string
	// printf will give a result like "students [   ]"
	students[0] = "Lisa"
	students[1] = "Lisa2"
	students[2] = "Lisa3"
	fmt.Printf("%v\n", grades)
	fmt.Printf("%v\n", students)
	fmt.Printf("Student #2: %v\n", students[1])
	fmt.Printf("Length: %v\n", len(students))

	// identity matrix
	//var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	var identityMatrix [3][3]int = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}} // suggested by extension
	// var identityMatrix [3][3]int
	// identityMatrix[0] = [3]int{1, 0, 0}
	// identityMatrix[1] = [3]int{0, 1, 0}
	// identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)

	// golang copies over the whole array
	a := [...]int{1, 2, 3}
	b := &a // if a pointer is added here, then it will refer to the address of the former array
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
}
