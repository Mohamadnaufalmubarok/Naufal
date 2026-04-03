package main

import "fmt"

func cetakBarisan(n int) {
	if n == 1 {
		fmt.Print(n, " ")
		return
	}
	fmt.Print(n, " ")
	cetakBarisan(n - 1)
	fmt.Print(n, " ")
}

func main() {
	var n int
	fmt.Scan(&n)
	cetakBarisan(n)
	fmt.Println()
}