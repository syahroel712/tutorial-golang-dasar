package main

import (
	"fmt"
	"time"
)

func main(){
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(2021, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	// time.RFC3339
	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "2021-09-14")
	fmt.Println(parse)

}


// catatan
// package yg berisikan fungsionalitas utk management waktu di go