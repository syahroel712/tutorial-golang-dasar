package main

import "fmt"

func main(){
	name := "Syahrul"
	counter := 0

	increment := func(){
		// closure
		name = "Nama diganti"
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()

	fmt.Println(name)
	fmt.Println(counter)
}










// catatan closure
// closure adalah kempuan function berinteraksi dengan data2 disekitarnya dalam scope yg sama
// harap gunakan fitur closure ini dengan bijak saat membuat kita aplikasi

