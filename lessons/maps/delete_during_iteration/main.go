package main

import "fmt"

func main() {
	data := map[string]int{"foo": 0, "bar": 1, "baz": 2}
	for key := range data {
		if key == "foo" {
			delete(data, "bar")
		}
		if key == "bar" {
			delete(data, "foo")
		}
	}

	// Проход по мапе рандомизирован, поэтому порядок удаления может отличаться.
	// Например, в при одном проходе удалится "foo", в другом "bar".
	fmt.Println(data)
}
