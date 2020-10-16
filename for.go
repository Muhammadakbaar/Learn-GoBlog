package main

import "fmt"

func main() {
	counter1 := 1

	for counter1 <= 5{
		fmt.Println("Perulangan Counter 1 = ", counter1)
		counter1++
	}
	fmt.Println("")

	for counter2 := 1; counter2 <= 5; counter2++{
		fmt.Println("Perulangan Counter 2 = ", counter2)
	}
	slice := [] string{"Muhammad", "Akbar", "Hakim"}

	fmt.Println(slice)
} 
