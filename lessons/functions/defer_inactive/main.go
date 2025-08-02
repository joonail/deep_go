package main

import "fmt"

func main() {
	defer fmt.Println("code")
	if false {
		defer fmt.Println("unreacheable code") // этот defer будет отложен, только если будет выполнен вход в эту точку кода.
	}

	return
	defer fmt.Println("unreacheable code")
}
