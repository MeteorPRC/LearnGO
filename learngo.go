package main

import (
	"fmt"
)

func main() {
	var name string="fdbf" 
	var age int
	var isok bool
	var on float64
	fmt.Println(name,isok,on,age)
	m:=15
	fmt.Print(m)
	const (
		a,b=iota+1 ,iota
		c,d=iota,iota
	)
	fmt.Print(a,b,c,d)
}