package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/rafael977/exp_quicksort/pkg/sort"
)

func main() {
	size := 10
	if len(os.Args) == 2 {
		size, _ = strconv.Atoi(os.Args[1])
	}

	arr := generateArray(size)

	start := time.Now()
	// Quicksort
	// fmt.Println("Array before sorting", arr)
	sort.QuickSort(arr)
	// fmt.Println("Array after sorting", arr)

	fmt.Printf("Execution time for quicksort %s\n", time.Since(start))
}

func generateArray(n int) []int {
	return rand.Perm(n)
}
