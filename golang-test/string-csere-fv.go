package main

import "fmt"

func add(x int, y int) int {
	return x + y
}
func swap(a, b string) (string, string) {
	return b, a
}
func main() {
	fmt.Println("Az összegük:", add(42, 53))

	c, d := swap("hello", "világ")
	fmt.Println(c, d)
}
