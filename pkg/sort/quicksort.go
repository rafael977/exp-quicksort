package sort

import (
	"math/rand"
)

// QuickSort sorts input array using basic quicksort
func QuickSort(arr []int) {
	if len(arr) > 1 {
		p := partition1(arr)
		QuickSort(arr[:p])
		QuickSort(arr[p+1:])
	}
}

func partition1(arr []int) int {
	lo, hi := 0, len(arr)-1
	p := rand.Intn(hi - lo + 1)
	pivot := arr[p]
	swap(arr, p, hi)

	i, j := 0, hi-1
	for i <= j {
		for i <= j && arr[i] <= pivot {
			i++
		}
		for i <= j && arr[j] >= pivot {
			j--
		}
		if i < j {
			swap(arr, i, j)
		}
	}
	swap(arr, i, hi)

	return i
}
