package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {
	//Init
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}
	floats := map[string]float64{
		"first":  35.98,
		"Second": 12.45,
	}

	fmt.Printf("no Genericos %v, and %v\n", SumInit(ints), Sumfloat(floats))
	fmt.Printf("With Generics %v, and %v\n", SumIntOrFloats(ints), SumIntOrFloats(floats))
	fmt.Printf("Generic Sums with Constraint  %v, and  %v\n", SumNumbers(ints), SumNumbers(floats))
}

func SumInit(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

//Sumfloat
func Sumfloat(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func SumIntOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}

	return s
}
