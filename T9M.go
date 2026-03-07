package main

import (
	"fmt"
)

func main() {
	var k int

	fmt.Print("Nilai K = ")
	fmt.Scan(&k)
	hampiranAkar2 := 1.0

	for i := 0; i <= k; i++ {
		kf := float64(i)
		
		pembilang := (4*kf + 2) * (4*kf + 2)
		penyebut := (4*kf + 1) * (4*kf + 3)
		fk := pembilang / penyebut
		hampiranAkar2 *= fk
	}

	fmt.Printf("Nilai akar 2 = %.10f\n", hampiranAkar2)
}