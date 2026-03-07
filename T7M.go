package main

import (
	"fmt"
	"strings"
)

func main() {
	var bunga string
	var pita string
	var jumlah int

	for i := 1; ; i++ {
		fmt.Print("Bunga ", i, ": ")
		fmt.Scan(&bunga)

		if strings.ToUpper(bunga) == "SELESAI" {
			break
		}

		pita = pita + bunga + " - "
		jumlah++
	}

	fmt.Println("Pita:", pita)
	fmt.Println("Bunga:", jumlah)
}