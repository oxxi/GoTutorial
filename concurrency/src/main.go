package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	words := []string{
		"Alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	wg.Add(len(words))

	for i, x := range words {
		go PrintSomething(fmt.Sprintf("%d : %s", i, x), &wg)
	}
	wg.Wait()
	wg.Add(1)
	PrintSomething("this is a test for second thing", &wg)
}

func PrintSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}
