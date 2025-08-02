package main

import "fmt"

func Generator(number int) func() int {
	return func() int {
		// здесь r это ссылка на number, а не локальная копия переменной, т.к. мы ссылаемся на переменную из вне функции,
		// Поэтому запоминается значение number. Т.е. принцип замыкания.
		r := number
		number++
		return r
	}
}

func main() {
	generator := Generator(100)
	for i := 0; i <= 200; i++ {
		fmt.Println(generator())
	}

	fmt.Println(generator()) // 301
	fmt.Println(generator()) // 302
}
