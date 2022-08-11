package main

import "fmt"

func main() {
	i := 7
	incr(&i) //its important we use & operator for passing address
	fmt.Println(i)
}

func incr(x *int) { //its important we use * operator for receiving the address
	*x++ //dereference it as we dont want to increment the address of the passed var, but the value at the address.
}
