package main

import (
	"belajar-golang-dasar/packageimport"
)

func main() {
	packageimport.SayHello("Syahrul")
}




// cattaan

// Package
// package adalah tempat yg bisa digunakan utk mengorganisir kode program yg kita buat di golang
// dengan menggunakan packeage, kita bisa merapikan kode program yg kita buat
// package sendiri sebenarnya hanya direktori folder di sistem operasi kita
// kata kuncinya package
// didalam direktori yg sama tidak boleh ada nama function yg sama, jika beda direktori boleh

// Import
// secara standart, file go hanya bisa mengakses file go lainya yg berda dalam package yg sama
// jika ingin mengkases file go yg berada di luar package maka kita bisa menggunakan import
// caranya kita harus panggil nama_folder_utama/nama_folder_import
// cara menggunakan nama_folder_import.nama_func


//Membuat golang bisa di jalankan di folder lain tanpa harus di GOROOT
// go env -w GO111MODULE=off

// atau juga bisa menggunakan perintah di bawah dan menghasilkan file go.mod 
// go mod init

