package main
import "fmt"

func main() {
	var n int
	var bunga string
	pita := ""

	fmt.Print("N: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Print("Bunga ", i, ": ")
		fmt.Scan(&bunga)

		pita = pita + bunga + " - "
	}

	fmt.Println("Pita:", pita)
}