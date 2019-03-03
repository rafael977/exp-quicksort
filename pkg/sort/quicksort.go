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
	arr[p], arr[hi] = arr[hi], arr[p]

	i, j := 0, hi-1
	for i <= j {
		for i <= j && arr[i] <= pivot {
			i++
		}
		for i <= j && arr[j] >= pivot {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i], arr[hi] = arr[hi], arr[i]

	return i
}
