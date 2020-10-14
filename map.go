package main

import "fmt"

func main()  {

	person := map[interface{}]interface{}{
		"nama" : "Akbar",
		"asal" : "Tidore",
		"Sitamvan" : 1,
		"Akbar Tamvan" : true,
	}
	fmt.Println(person)
	fmt.Println(person["nama"])
	fmt.Println(person["asal"])
	fmt.Println(person["Akbar Tamvanf"])


}
