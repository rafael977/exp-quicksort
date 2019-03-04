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

	// arr := genPerm(size)
	arr := genRand(size)
	// arr := genSeq(size)
	// arr := genRev(size)

	// Quicksort
	arr1 := make([]int, len(arr))
	copy(arr1, arr)
	start := time.Now()
	// fmt.Println("Array before sorting", arr1)
	sort.QuickSort(arr1)
	// fmt.Println("Array after sorting", arr1)
	fmt.Printf("Execution time for quicksort %s\n", time.Since(start))

	// // 2 pivot Quicksort
	arr2 := make([]int, len(arr))
	copy(arr2, arr)
	start = time.Now()
	sort.QuickSort2P(arr2)
	fmt.Printf("Execution time for 2 pivot quicksort %s\n", time.Since(start))

	// // 3 pivot Quicksort
	arr3 := make([]int, len(arr))
	copy(arr3, arr)
	start = time.Now()
	sort.QuickSort3P(arr3)
	fmt.Printf("Execution time for 3 pivot quicksort %v\n", time.Since(start))
}

func genPerm(n int) []int {
	return rand.Perm(n)
}

func genRand(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}

func genSeq(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
}

func genRev(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = n - i - 1
	}
	return arr
}
