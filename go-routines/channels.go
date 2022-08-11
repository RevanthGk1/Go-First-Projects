package main

func SendDataToChannel(ch chan int, val int) {
	ch <- val
}

func main() {
	var v int
	ch := make(chan int)
	go SendDataToChannel(ch, 101)
	v = <-ch //to recieve the changed value from the func like out parameter!
	println(v)
}
