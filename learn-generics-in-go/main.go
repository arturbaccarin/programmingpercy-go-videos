// https://youtu.be/fpLz-SRZ2ho
// 28:40
package main

import "fmt"

type Substractable interface {
	int | int32 | int64 | float32 | float64 | uint | uint32 | uint64
}

type Moveable[S Substractable] interface {
	Move(S)
}

func Move[V Moveable[S], S Substractable](v V, distance, meters S) S {
	v.Move(meters)
	return Subtract(distance, meters)
}

type Person[S Substractable] struct {
	Name string
}

func (p Person[S]) Move(meters S) {
	fmt.Printf("%s moved %v meters\n", p.Name, meters)
}

type Car[S Substractable] struct {
	Name string
}

func (c Car[S]) Move(meters S) {
	fmt.Printf("%s moved %v meters\n", c.Name, meters)
}

func main() {
	p := Person[float64]{Name: "John"}
	c := Car[int]{Name: "Ferrari"}

	milesToDestination := 100

	distanceLeft := Move(c, milesToDestination, 95)

	fmt.Println(distanceLeft)
	fmt.Println(p)

	newDistanceLeft := Move[Person[float64], float64](p, float64(distanceLeft), 100)

	fmt.Println(newDistanceLeft)
}

func Subtract[V Substractable](a, b V) V {
	return a - b
}

/*
// ~int any date type that derives from int is accepted

type MyOwnInteger int

// Results is a SLICE of results that are Subtractable
type Results[T Substractable] []T

type Results2[T interface{}] []T

type ResultsC[T comparable] []T

func main() {
	var a int = 20
	var b int = 10

	var c float32 = 20
	var d float32 = 10.5

	var e uint = 20
	var f uint = 10

	var g MyOwnInteger = 20
	var h MyOwnInteger = 10

	result := Subtract(a, b)

	// We need to cast data type into int here
	resultFloat := Subtract(c, d)

	resultUint := Subtract[uint](e, f)

	resultMyOwnInteger := Subtract(g, h)

	// Will return 10
	fmt.Println("Result: ", result)
	// Will return 10 - Which is wrong
	fmt.Println("ResultFloat: ", resultFloat)

	fmt.Println("ResultUint: ", resultUint)

	var resultStorage Results[int] // we need to use the type parameter

	resultStorage = append(resultStorage, result, int(resultFloat))

	var resultStorage2 Results[float32] // you can't pass interface substractable
	resultStorage2 = append(resultStorage2, resultFloat)

	var resultStorage Results2[any]
	resultStorage = append(resultStorage, result, resultFloat, resultUint, resultMyOwnInteger)
}

// Subtraction will subtract two numbers
// [V int] - type constraint, need to be a interface or multiple type
// int | float32: anonymous interface

func Subtract[V int | float32](a, b V) V {
	return a - b
}


func Subtract[V Substractable](a, b V) V {
	return a - b
}

*/

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
