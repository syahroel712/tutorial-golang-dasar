package main

import (
	"fmt"
	"regexp"
)

func main(){
	var regex *regexp.Regexp = regexp.MustCompile("e([a-z])o")

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("edo"))
	fmt.Println(regex.MatchString("eTo"))
	// fmt.Println(regex.MatchString("enso"))

	search := regex.FindAllString("eko eka edo eto eno eki", 10)
	fmt.Println(search)

}

// catatan
// utilitas di go untuk melakukan pencarian regular expression
// regular expression di go mengunakan libray C yg dibuat Google bernama RE2