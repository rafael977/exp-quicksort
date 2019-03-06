package sort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

// TestQuickSort tests QuickSort3P function
func TestQuickSort(t *testing.T) {
	arr := rand.Perm(50)
	fmt.Println("Arr before sorting", arr)

	QuickSort(arr)
	verify(arr, t)

	fmt.Println("Arr after sorting", arr)
}

// TestQuickSort2P tests QuickSort3P function
func TestQuickSort2P(t *testing.T) {
	arr := GenRand(10000000)

	QuickSort2P(arr)
	verify(arr, t)
}

// TestQuickSort3P tests QuickSort3P function
func TestQuickSort3P(t *testing.T) {
	arr := GenRand(100)

	QuickSort3P(arr)
	verify(arr, t)
}

// TestQuickSort5P tests QuickSort5P function
func TestQuickSort5P(t *testing.T) {
	arr := GenRand(10000)
	// fmt.Println("Arr before sorting", arr)

	QuickSort5P(arr)
	verify(arr, t)

	// fmt.Println("Arr after sorting", arr)
}

// TestQuickSort7P tests QuickSort7P function
func TestQuickSort7P(t *testing.T) {
	arr := GenRand(1000)

	QuickSort7P(arr)
	verify(arr, t)
}

func verify(arr []int, t *testing.T) {
	if !sort.IsSorted(sort.IntSlice(arr)) {
		t.Errorf("Not sorted.")
	}
}
