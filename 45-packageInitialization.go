package main

import (
	"belajar-golang-dasar/database"
	"fmt"
)

func  main(){
	result := database.GetDatabase()
	fmt.Println(result)
}


// catatan

// package initialization
// saat kita membuat package, kita bisa membuat sebuah function yg akan diakses ketika package di jalankan
// ini sangat cocok ketika contohnya, jika package kita berisi function2 utk berkomunikasi degnan database, kita membuat function inisialisai utk membuka koneksi ke database
// utk membuat function yg diakses secara otomatis ketika package diakses, kita cukup membuat function dengan nama init
// contoh	
	// package database
	// var connection string
	// func init() {
	// 	connection = "MySQL"
	// }
	// func GetDatabase() string{
	// 	return connection
	// }


// blank identifier
// kadang kita hanya ingin menjalankan init function di package tanpa harus mengeksekusi salah satu funcion yg ada di package
// secara defaul, golang akan komplen ketika ada package yg diimport namun tidak digunakan
// untuk menangani hal tsb, kita bisa menggunakan blank indentifier (_) sebelum nama package ketika melakukan import
// intiny saat membuat sesuatu dan tidak digunakan maka gunakan _ sebagai tanda ini belum digunakan

