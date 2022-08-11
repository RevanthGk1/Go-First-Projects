package main

import "fmt"

func main() {
	//Arrays
	var a [5]int
	b := [5]int{3, 2, 5, 1, 3}
	c := []int{3, 2, 5, 1, 3} //slices or dynamic arrays without fixed length, we can append to this later
	c = append(c, 12)

	b[2] = 7
	fmt.Println("Arrays demo")
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(b)
	fmt.Println(c)

	//Dictionaries/Maps
	myMap := make(map[string]int)
	fmt.Println("Maps demo")
	myMap["item1"] = 11
	myMap["item2"] = 23
	myMap["item3"] = 34
	delete(myMap, "item2")
	println(myMap["item3"])

	//Diff type of for loop for array & maps
	for index, value := range b {
		println("index:", index, "value", value)
	}

	//for map
	for key, value := range myMap {
		println("key:", key, "value", value)
	}

}
