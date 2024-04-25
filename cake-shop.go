package main

import (
	"fmt"
)

func main() {
	// Data pesanan dalam setiap batch
	order := map[string][]string{
		"batch1": {"brownies", "nastar", "muffin", "muffin", "brownies", "pie", "nastar", "muffin"},
		"batch2": {"brownies", "pie", "pie", "nastar"},
	}

	// Simpan batch
	batchNames := make([]string, 0, len(order))
	for batch := range order {
		batchNames = append(batchNames, batch)
	}

	// Iterasi setiap nama batch dan print pesanan
	for _, batch := range batchNames {
		// print pesan bahwa sedang memproses batch
		fmt.Println("Processing ", batch)

		// Ambil daftar pesanan dalam batch
		items := order[batch]

		// Hitung jumlah setiap jenis pesanan dalam batch
		counts := make(map[string]int)
		for _, item := range items {
			counts[item]++
		}

		// Iterasi setiap jenis pesanan dan print jumlahnya
		for item, count := range counts {
			fmt.Printf("%d %s is ready!\n", count, item)
		}

		// print baris kosong buat spasi aja
		fmt.Println()
	}
}
