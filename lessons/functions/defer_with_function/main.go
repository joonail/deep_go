package main

import "fmt"

func get() string {
	fmt.Println("1")
	return ""
}

func handle(string) {
	fmt.Println("3")
}

func process() {
	// здесь функция откладывается с заранее вычисленным аргументами. В данном случае это get() и отображается "1".
	// Затем эта функция откладывается и вызывается в конце функции process().
	defer handle(get())
	fmt.Println("2") // Отображается "2".
	// Затем происходит вызов handle и отображается "3".
}

func main() {
	process()
}
