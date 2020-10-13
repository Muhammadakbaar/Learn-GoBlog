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
}
