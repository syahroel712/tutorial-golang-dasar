package main

import (
	"flag"
	"fmt"
)

func main(){
	// hasil dari fla.string adalah pointer
	var host *string = flag.String("host", "localhost", "Put your database host")
	var username *string = flag.String("username", "root", "Put your database username")
	var password *string = flag.String("password", "root", "Put your database password")
	var number *int = flag.Int("number", 100, "Put your number")

	// untuk melakukan proses parsing gunakan flag.Parse()
	flag.Parse()

	// hasilnya berupa code
	// fmt.Println("Host : ", host)
	// fmt.Println("Username : ", username)
	// fmt.Println("Password : ", password)

	// jika ingin ambil value, maka gunakan pointer 
	fmt.Println("Host : ", *host)
	fmt.Println("Username : ", *username)
	fmt.Println("Password : ", *password)
	fmt.Println("Number : ", *number)
}



// catatan

// package flag erat kaitannya dengan package os
// package flag berisikan funsionalitas  untuk memparsing command line argument yg ada di os.Args
// hasil dari flag.String() adalah berupa pinter jadi gunakan * untuk memanggilnya
	// contoh
		// func main(){
		// 	host := flag.String("host", "localhost","Put your database host")
		// 	username := flag.String("username", "username","Put your database username")
		// 	password := flag.String("password", "password","Put your database password")
		
		// 	flag.Parse()

		// 	fmt.Println(*host, *username, *password)
			
		// }
// pastikan memanggil flag.Parse() setelah mendeklarasikan flag
