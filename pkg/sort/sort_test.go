package sort

import (
	"fmt"
	"math/rand"
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
	arr := rand.Perm(50)
	fmt.Println("Arr before sorting", arr)

	QuickSort2P(arr)
	verify(arr, t)

	fmt.Println("Arr after sorting", arr)
}

// TestQuickSort3P tests QuickSort3P function
func TestQuickSort3P(t *testing.T) {
	arr := rand.Perm(50)
	fmt.Println("Arr before sorting", arr)

	QuickSort3P(arr)
	verify(arr, t)

	fmt.Println("Arr after sorting", arr)
}

func verify(arr []int, t *testing.T) {
	for i := range arr {
		if i < len(arr)-1 && arr[i]+1 != arr[i+1] {
			t.Errorf("Not sorted. Index: %d", i)
			break
		}
	}
}
