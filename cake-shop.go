package main

import (
	"fmt"
)

func main() {
	order := map[string][]string{
		"batch1": {"brownies", "nastar", "muffin", "muffin", "brownies", "pie", "nastar", "muffin"},
		"batch2": {"brownies", "pie", "pie", "nastar"},
	}

	for batch, items := range order {
		fmt.Println("Processing", batch)
		counts := make(map[string]int)
		for _, item := range items {
			counts[item]++
		}

		for item, count := range counts {
			fmt.Printf("%d %s is ready!\n", count, item)
		}
		fmt.Println()
	}
}
