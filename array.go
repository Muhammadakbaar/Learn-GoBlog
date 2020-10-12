package main

import "fmt"

func main()  {

	var nama [3] string
	nama[0] = "Muhammad"
	nama[1] = "Akbar"
	nama[2] = "Hakim"

	fmt.Println(nama[0])
	fmt.Println(nama[1])
	fmt.Println(nama[2])

	var nilai = [3] int{
		10,
		20,
		30,
	}
	fmt.Println(nilai)

	fmt.Println(len(nama))
	fmt.Println(len(nilai))

	var coba[] int
	fmt.Println(len(coba))

}
