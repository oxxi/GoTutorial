package main

import "fmt"

func main() {
	var a [4]int

	a[3] = 555

	for i, x := range a {
		fmt.Printf("index %d, value %v\n", i, x)
	}

	fmt.Printf("capacity of slice is %v and his len is %v\n", cap(a), len(a))
	//made an slice with make
	fmt.Println("===========================================================================")
	mySlice := make([]string, 5, 6)

	fmt.Printf("capacity of slice is %v and his len is %v\n", cap(mySlice), len(mySlice))
	for inx, v := range mySlice {
		fmt.Printf("index %d, value %v\n", inx, v)
	}
	mySlice = mySlice[:6] //extends the length
	mySlice[5] = "tada .."
	fmt.Println("===========================================================================")

	for inx, v := range mySlice {
		fmt.Printf("index %d, value %v\n", inx, v)
	}
}
