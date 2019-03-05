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

	// arr := sort.GenPerm(size)
	arr := sort.GenRand(size)
	// arr := sort.GenSeq(size)
	// arr := sort.GenRev(size)

	// Quicksort
	arr1 := make([]int, len(arr))
	copy(arr1, arr)
	start := time.Now()
	// fmt.Println("Array before sorting", arr1)
	sort.QuickSort(arr1)
	// fmt.Println("Array after sorting", arr1)
	fmt.Printf("Execution time for quicksort %s\n", time.Since(start))

	// 2 pivot Quicksort
	arr2 := make([]int, len(arr))
	copy(arr2, arr)
	start = time.Now()
	sort.QuickSort2P(arr2)
	fmt.Printf("Execution time for 2 pivot quicksort %s\n", time.Since(start))

	// 3 pivot Quicksort
	arr3 := make([]int, len(arr))
	copy(arr3, arr)
	start = time.Now()
	sort.QuickSort3P(arr3)
	fmt.Printf("Execution time for 3 pivot quicksort %v\n", time.Since(start))

	// 5 pivot Quicksort
	arr5 := make([]int, len(arr))
	copy(arr5, arr)
	start = time.Now()
	sort.QuickSort5P(arr5)
	fmt.Printf("Execution time for 5 pivot quicksort %v\n", time.Since(start))

	// 7 pivot Quicksort
	arr7 := make([]int, len(arr))
	copy(arr7, arr)
	start = time.Now()
	sort.QuickSort7P(arr7)
	fmt.Printf("Execution time for 7 pivot quicksort %v\n", time.Since(start))
}
