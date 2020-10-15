package main

import "fmt"

func main()  {

	person := map[interface{}]interface{}{
		"nama" : "Akbar",
		"asal" : "Tidore",
		"Sitamvan" : 1,
		"Akbar Tamvan" : true,
	}
	fmt.Println(person["nama"])
	fmt.Println(person["asal"])
	fmt.Println(person["SiTamvan"])
	fmt.Println(person["Akbar Tamvan"])

	var book map[string]string = make(map[string] string)
	book["title"] = "belajar Golang"
	book["author"] = "sitamvan"
	book["ups"] = "salah"

	delete(book, "ups")

	fmt.Println(book)


}
