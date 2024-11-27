package main

import "fmt"

func main() {

	var berat, beratKg, sisaGr, totalBiaya int

	fmt.Scanln(&berat)

	beratKg = berat / 1000
	sisaGr = berat % 1000
	totalBiaya = beratKg * 10000

	if beratKg <= 10 {
		if sisaGr >= 500 {
			totalBiaya = totalBiaya + (sisaGr * 5)
		} else {
			totalBiaya = totalBiaya + (sisaGr * 15)

		}

	}
	fmt.Printf("Berat total: %d kg dan %d gram\n", beratKg, sisaGr)
	fmt.Printf("Biaya pengiriman: Rp %d\n", totalBiaya)
}
