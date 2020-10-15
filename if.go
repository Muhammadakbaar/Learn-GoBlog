package main

import "fmt"

func main(){
	 var nama = "sitamvan"

	 if nama == "sitamvan"{
	 	fmt.Println("Halo tamvan")
	 }else if nama == "akbar  "{
	 	fmt.Println("Halo akbar")
	 }else{
	 	fmt.Println("default")
	 }

	 if panjang := len(nama); panjang < 5 {
	 	fmt.Println("Nama Terlalu Panjang")
	 }else {
	 	fmt.Println("Nama Terlalu Pendek")
	 }
}