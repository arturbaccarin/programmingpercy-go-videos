// https://youtu.be/fpLz-SRZ2ho
package main

import "fmt"

type Substractable interface {
	int | int32 | int64 | float32 | float64 | uint | uint32 | uint64
}

func main() {
	var a int = 20
	var b int = 10

	var c float32 = 20
	var d float32 = 10.5

	var e uint = 20
	var f uint = 10

	result := Subtract(a, b)

	// We need to cast data type into int here
	resultFloat := Subtract(c, d)

	resultUint := Subtract[uint](e, f)

	// Will return 10
	fmt.Println("Result: ", result)
	// Will return 10 - Which is wrong
	fmt.Println("ResultFloat: ", resultFloat)

	fmt.Println("ResultUint: ", resultUint)
}

// Subtraction will subtract two numbers
// [V int] - type constraint, need to be a interface or multiple type
// int | float32: anonymous interface
/*
func Subtract[V int | float32](a, b V) V {
	return a - b
}
*/

func Subtract[V Substractable](a, b V) V {
	return a - b
}

/*
func main() {
	var a int = 20
	var b int = 10

	var c float32 = 20
	var d float32 = 10.5

	result := Subtract(a, b)

	// We need to cast data type into int here
	resultFloat := Subtract(int(c), int(d))

	// Will return 10
	fmt.Println("Result: ", result)
	// Will return 10 - Which is wrong
	fmt.Println("ResultFloat: ", resultFloat)
}

// Subtraction will subtract two numbers
func Subtract(a, b int) int {
	return a - b
}

func SubtractFloat32(a, b float32) float32 {
	return a - b
}
*/
