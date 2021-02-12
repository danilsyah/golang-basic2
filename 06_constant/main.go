package main

import "fmt"

func main() {
	// constant adalah variable yang nilainya
	// tidak bisa diubah lagi setelah pertama kali diberi nilai
	// saat pembuatan constant, wajib langsung menginisialisasi datanya

	const firstName string = "Danil"
	const age = 26

     // multiple variabel const
     const (
          lastName = "Syah"
          isMarried = true
     )

	// age = 20 error

	fmt.Println(firstName)
	fmt.Println(age)
     fmt.Println(lastName)
     fmt.Println(isMarried)

}