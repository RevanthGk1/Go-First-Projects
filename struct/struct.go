package main

import "fmt"

type Address struct {
	streetName string
	streetNum  int
}

type Person struct {
	name    string
	age     int
	address Address
}

func main() {

	addres := Address{streetName: "ostera", streetNum: 2}
	p := Person{name: "revanth", age: 22, address: addres}

	fmt.Println(p)
}
