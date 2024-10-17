package main

import (
	"fmt"
)

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x * x
}

func h(x int) int {
	return x + 1
}

func fogoh(a int) int {
	return f(g(h(a)))
}

func gohof(b int) int {
	return g(h(b))
}

func hofog(c int) int {
	return h(f(g(c)))
}

func main() {
	var a, b, c int
	fmt.Print("masukkan tiga buah bilangan bulat: ")
	fmt.Scan(&a, &b, &c)

	fmt.Println("hasil: ")
	fmt.Println("fogoh(a) = ", fogoh(a))
	fmt.Println("gohof(b) = ", gohof(b))
	fmt.Println("hofog(c) = ", hofog(c))
}
