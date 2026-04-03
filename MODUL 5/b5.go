package main

import "fmt"

func cetakGanjil(n, i int) {
	if i > n {
		fmt.Println()
		return
	}
	if i%2 != 0 {
		fmt.Print(i, " ")
	}
	cetakGanjil(n, i+1)
}

func main() {
	var n int
	fmt.Scan(&n)
	cetakGanjil(n, 1)
}