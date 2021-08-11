package main

import "fmt"

func main(){

	var syahrul Customer

	syahrul.Name = "Syahrul"
	syahrul.Address = "Padang"
	syahrul.Age = 20

	fmt.Println(syahrul)
	fmt.Println(syahrul.Name)
	fmt.Println(syahrul.Address)
	fmt.Println(syahrul.Age)


	joko := Customer{
		Name : "Joko",
		Address : "Jakarta",
		Age : 20,
	}

	fmt.Println(joko)
	
	// rentan error
	budi := Customer{"Budi", "Indo", 20}
	fmt.Println(budi)

}

type Customer struct {
	Name, Address string
	Age int
}





// catatan
// struct adalah sebuah template data yg digunakan untk menggabungkan nol / lebih tipe data lainya dalam satu kesatuan
// struct biasanya representasi data lam program aplikasi yg kita buat
// data di struct disimpan di dalam field
// sederhananya struct adalah kumpulan dari field
	// contoh
		// type Customer struct {
		// 	name, address 	string
		// 	age 			int
		// }


// membuat data struct
// struct adalah template data / prototype (atau di php seperti class)
// struct tidak bisa langsung di pakai
// namaun kita bisa membuat data/object dari struct yg telah di buat
	// contoh
		// func main(){ 
		// 	var syahrul Customer
		// 	syahrul.Name = "Syahrul"
		// 	syahrul.Address = "Padang"
		// 	syahrul.Age = 20
		// }


// struct literals
// sebelumnnya kita tealah membuat data struct, namun sebenarnya ada banyak cara yg bisa digunakan utk membuat data dari struct
			// contoh 1
				// joko := Customer{
				// 	Name : "Joko",
				// 	Address : "Padang"
				// 	Age : 20
				// }

			// contoh 2, tapi harus sesuai urutan deklarasi structnya
				// budi := Customer{"budi", "Indo", 20}
