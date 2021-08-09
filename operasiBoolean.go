package main

import "fmt"

func main(){

	var ujian = 90
	var absensi = 80

	var lulusNilaiAkhir bool = ujian >= 80
	var lulusAbsensi bool = absensi >= 80

	var lulus =  lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulus)

}



// catatan

// operasi boolean hanya utk data boolean
	// && => and 
		// => true && true =  true
		// => true && false =  false
		// => false && true =  false
		// => false && false =  false
	// || => or 
		// => true || true = true
		// => true || false = true
		// => false || true = true
		// => false || false = false
	// ! => not 
		// => !true = false
		// => !false = true
