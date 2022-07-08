package main

import "fmt"

func main() {
	inp := "Hola Mundo"
	rv := Reverse(inp)
	fmt.Println("original: ", inp)
	fmt.Println("Reverse: ", rv)
	fmt.Println("reverse again: ", Reverse(rv))
}

func Reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}
