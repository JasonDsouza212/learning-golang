package main

import "fmt"

func main() {
	mybill := newBill("jasons bill")
	fmt.Println(mybill.format())
}
