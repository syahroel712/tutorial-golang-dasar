package main 

import "fmt"

func main(){

	counter := 1

	for counter <= 5 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	
	// for statement
	for i := 1; i <= 5; i++{
		fmt.Println("I ke ", i)
	}

	// for range
	names := []string{"Andi", "Ani", "Ana"}

	for i, name := range names {
		fmt.Println("index", i, "=", name)
	}

	for _, name := range names {
		fmt.Println(name)
	}


	person := make(map[string]string)
	person["name"] = "Syahrul"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}

}

// catatan

// cuma ada satu perulangan di golang yaitu for loops

// for dengan statement
// dalam for, kita bisa menambahkan statement , dimana terdapat 2 statement yg bisa ditambahkan di for
	// init statement, yaitu statement sebelum for di eksekusi
	// post statement, yaitu statement yg akan selalu dieksekusi di akhir tiap perulangan

// for range
// for bisa digunakan utk melakukan iterasi terhadap semu data collection(array, slice, map)
	// for i, name := range names {
	// 	fmt.Println("index", i, "=", name)
	// }
// jika tidak membtuhkan i di for range, maka bisa gunakan _(underscore) . Hal ini bertujuan memberitahu go bahwa variable ini tidak dibutuhkan