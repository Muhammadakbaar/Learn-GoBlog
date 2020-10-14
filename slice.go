package main

import "fmt"

func main()  {
	var month = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	var slice1 = month[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	//month[5] = "Diubah"
	//fmt.Println(slice1)

	slice1[0] = "Mei Ubah"
	fmt.Println(month)


	var slice2 = month[2:4]
	fmt.Println(slice2) 

	var slice3 = append(slice2, "Akbar")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(month)

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Muhammad"
	newSlice[1] = "Akbar"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [5]int {1, 2, 3, 4, 5}
	iniSlice :=[]int{1, 2, 3, 4, 5 }

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
