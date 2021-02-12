package main

import "fmt"

func main() {
	fmt.Println("Danil Syah")
     fmt.Println("Ini adalah String")

     // method build in String
     fmt.Println("berjumlah = ", len("danil"))
     fmt.Println("udin"[1])
     fmt.Println("danil syah"[3])

     // konversi byte to string
     fmt.Println(string("udin"[1]))
     fmt.Println(string("danil syah"[3]))
}