package main

import "fmt"

func main() {
	// variable adalah tempat menyimpan data
     // deklarasi awal variable
	var name string
     var age = 26
     isMarried := true
     country := "Indonesia"
     // deklarasi multiple variabel
     var (
          firstName = "Danil"
          lastname = "Syah"
     )


	name = "Danil syah"
	fmt.Println("Nama saya : ", name)

     name = "Danil"
     fmt.Println(name)
     fmt.Println(age)
     fmt.Println(country)
     fmt.Println(isMarried)
     fmt.Println(firstName, " ",lastname)
}