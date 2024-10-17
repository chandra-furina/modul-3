package main

import (
	"fmt"
)

func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * faktorial(n-1)
}

func permutasi(n, r int) int {
	if n >= r {
		return faktorial(n) / faktorial(n-r)
	}
	return 0
}

func kombinasi(n, r int) int {
	if n >= r {
		return faktorial(n) / (faktorial(r) * faktorial(n-r))
	}
	return 0
}

func main() {
	var a, b, c, d int
	fmt.Print("masukkan bilangan: ")
	fmt.Scan(&a, &b, &c, &d)

	permutasiAC := permutasi(a, c)
	kombinasiAC := kombinasi(a, c)
	permutasiBD := permutasi(b, d)
	kombinasiBD := kombinasi(b, d)

	fmt.Printf("hasil permutasi a c: %d\n", permutasiAC)
	fmt.Printf("hasil kombinasi a c: %d\n", kombinasiAC)
	fmt.Printf("hasil permutasi b d: %d\n", permutasiBD)
	fmt.Printf("hasil kombinasi b d: %d\n", kombinasiBD)
}
