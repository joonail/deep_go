package main

import "fmt"

// Need to show solution

const (
	StatusOk    = "ok"
	StatusError = "error"
)

func notify(status string) {
	fmt.Println(status)
}

func process() {
	var status string
	defer func(s string) { // s здесь копия status.
		notify(s)
	}(status) // откладывается с заранее вычисленным аргументом, который в данном случае это пустая строка "".

	// TODO вернуться к лекции 5.2_04:40

	// processing..
	status = StatusError
}

func main() {
	process()
}
