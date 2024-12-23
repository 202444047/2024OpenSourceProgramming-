package main

import (
	"fmt"
)

func main() {
	var emptyslice []bool
	//emptyslice = make([]bool, 5)
	fmt.Printf("%#v %d\n", emptyslice, len(emptyslice)) // slice zero value (nill), 0
	if len(emptyslice) == 0 {
		emptyslice = append(emptyslice, true)
	}
	fmt.Printf("%#v %d\n", emptyslice, len(emptyslice)) //[]bool true, 1

	gpas := [5]float64{3.5, 4.1, 4.5, 3.9, 4.23} // array := array literal
	gpa_slice := gpas[1:4]                       // [4.1, 4.5, 3.9]
	//gpa_slice[1] = 2.71
	gpas[2] = 2.71
	//gpa_slice = append(gpa_slice, 4.3)
	gpa_slice = append(gpa_slice, 4.3, 5.55)
	fmt.Println(len(gpa_slice), gpa_slice, gpas)

}
