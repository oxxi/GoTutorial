package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	evenOrOdd()
	whatTimeIs()
	loopsFor()
	switchControl()

	//goto
	var K int
	i := 0
Next:
	if i >= 5 {
		goto Exit
	}
	K = i + i
	fmt.Println(K, " in Goto")
	i++
	goto Next
Exit:
}

func whatTimeIs() {
	h := time.Now().Hour()

	if h < 12 {
		fmt.Printf("Now is AM time.")
	} else if h > 19 {
		fmt.Printf("Now is evening time.")
	} else {
		fmt.Printf("Now is afternoon time.")
	}
}

func evenOrOdd() {

	rand.Seed(time.Now().UnixNano())

	if n := rand.Int(); n%2 == 0 {
		fmt.Printf("es un numero par %v\n", n)
	} else {
		fmt.Printf("No es un numero par %v\n", n)
	}
}

func loopsFor() {
	//Loops
	for i := 0; i < 10; i++ {
		fmt.Printf("Index loop for %v\n", i)
	}
	//loop with statement
	var control = 0
	for {
		if control > 10 {
			fmt.Println("Saliendo de for con break")
			break
		}

		control++
	}

	//loop with statement and continue key
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("mod %v\n", i)
	}
}

func switchControl() {
	//switch statemnt
	rand.Seed(time.Now().UnixNano())
	switch n := rand.Intn(100); n % 9 {
	case 0:
		fmt.Println(n, " is a multiple of 9")
	case 1, 2, 3:
		fmt.Println(n, "Mod 9 is 1, 2, 3.")

	case 4, 5, 6:
		fmt.Println(n, "mod 9 is 4, 5 or 6")

	default:
		fmt.Println(n, "mod 9 is 7 or 8")
	}
	fmt.Println("----------------------fallthrough-------------------------")
	//switch whit fallthrough
	//fallthrough hace que pase por el siguiente bloque aunque este no cumpla con la condicion del switch
	rand.Seed(time.Now().UnixNano())
	switch n := rand.Intn(100) % 5; n {
	case 0, 1, 2, 3, 4:
		fmt.Println("n =", n)
		fallthrough
	case 5, 6, 7, 8:
		n := 99
		fmt.Println("n =", n)
		fallthrough
	default:
		fmt.Println("n =", n)
	}

	//switch without condition always true
	switch {
	case true:
		fmt.Println("Hello ")
	default:
		fmt.Println("golang")
	}
}
