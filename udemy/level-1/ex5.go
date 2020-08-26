package main

import "fmt"

type mytype int

func main() {
	var x mytype 
	fmt.Printf("%d %T\n", x, x)
	x=42
	fmt.Printf("%d %T\n", x, x)
	y :=  int(x)
	fmt.Printf("%d %T\n", y, y)


}