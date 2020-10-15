package main

import "fmt"

func main() {
	nama := "Muhammad"

	switch nama {
	case "Akbar":
		fmt.Println("Hallo Akbar")
		fmt.Println("Hallo Akbar")
	case "Muhammad":
		fmt.Println("Hallo Muhammad")
		fmt.Println("Hallo Muhammad")
	default:
		fmt.Println("Namanya Siapa")
		fmt.Println("Namanya Siapa")

	}
	switch  panjang := len(nama); panjang > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Terlalu Pendek")

	}
	panjang := len(nama)
	switch {
	case panjang > 10:
		fmt.Println("Nama Terlalu Panjang")
	case panjang > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Terlalu Pendek")

	}
}
