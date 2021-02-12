package main

import "fmt"

func main() {
	// konversi tipe data

	// konversi antar integer
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai64)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

     var name = "Danil"
     // konversi byte to string
     var i byte  = name[3]
     var iString string = string(i)

     fmt.Println(i)
     fmt.Println(iString)

}