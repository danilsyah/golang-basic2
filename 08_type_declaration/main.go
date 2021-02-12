package main

import "fmt"

func main() {
	// membuat alias tipe data yang sudah ada
	type NoKTP string
     type Married bool
     type Umur uint8

	var danil NoKTP
	danil = "2121010219219291"
     var marriedStatus Married = true
     var umurDanil Umur = 26

	fmt.Println(danil)
	fmt.Println(marriedStatus)

     if umurDanil > 18 {
          fmt.Println("Dewasa")
     }else{
          fmt.Println("Bocah")
     }
}