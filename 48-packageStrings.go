package main

import (
	"fmt"
	"strings"
)

func main(){

	fmt.Println(strings.Contains("Mbappe Neymar", "Neymar")) // hasilnya true dan caseSensitif
	fmt.Println(strings.Split("Syahrul Swimoc", " ")) // [Syahrul Swimoc]
	fmt.Println(strings.ToLower("Mbappe")) // jadi keci
	fmt.Println(strings.ToUpper("Mbappe")) // jadi besar
	fmt.Println(strings.ToTitle("mBappe kylian")) // jadi besar
	fmt.Println(strings.Trim("       Syahrul Swimoc       ", " ")) // hapus space yg ada di kanan dan kiri karakter
	fmt.Println(strings.ReplaceAll("mbappe kylian", "kylian", "neymar")) // replace

}



// catatan
// merupakan package yg berisikan function untuk memanipulasi tipe data String
