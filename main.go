package main

import (
	"fmt"
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

	arr := sort.GenPerm(size)
	// arr := sort.GenRand(size)
	// arr := sort.GenSeq(size)
	// arr := sort.GenRev(size)

	// Quicksort
	runSort(sort.QuickSort, 1, arr)
	// 2 pivot Quicksort
	runSort(sort.QuickSort2P, 2, arr)
	// 3 pivot Quicksort
	runSort(sort.QuickSort3P, 3, arr)
	// 5 pivot Quicksort
	runSort(sort.QuickSort5P, 5, arr)
	// 7 pivot Quicksort
	runSort(sort.QuickSort7P, 7, arr)
}

func runSort(fn func([]int), numPivots int, arr []int) {
	cpy := make([]int, len(arr))
	copy(cpy, arr)
	start := time.Now()
	fn(cpy)
	fmt.Printf("Execution time for %d pivot quicksort %.3fms\n", numPivots, time.Since(start).Seconds()*1000)
}
