package main

import "fmt"

func main() {

	// break
	for i := 0; i < 5; i++{
		if i == 3{
			break
		}
		fmt.Println("Perulangan ke", i)
	}

	// continue
	for a := 0; a < 10; a++{
		if a % 2 == 1 {
			continue
		}

		fmt.Println("Bil Genap =", a)

	}

}

// catatan

// break dan continue adalah kata kunci yg bisa digunakan dalam perulangan
// break digunakan untuk menghentikan seluruh perulangan
// continue digunakan untuk menghentikan perulangan yg berjalan, dan langsung melanjutkan ke perulangan selanjutnya
