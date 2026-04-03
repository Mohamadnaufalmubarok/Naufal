package main

import "fmt"

func cetakFaktor(n, i int) {
	if i > n {
		fmt.Println()
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	cetakFaktor(n, i+1)
}

func main() {
	var n int
	fmt.Scan(&n)
	cetakFaktor(n, 1)
}