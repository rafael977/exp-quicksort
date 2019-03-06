package sort

// QuickSort2P sorts input array with 2 pivot quickdsort
func QuickSort2P(arr []int) {
	if len(arr) > 1 {
		lp, rp := partition2V2(arr)
		QuickSort2P(arr[:lp])
		QuickSort2P(arr[lp+1 : rp])
		QuickSort2P(arr[rp+1:])
	}
}

func partition2(arr []int) (int, int) {
	lo, hi := 0, len(arr)-1
	if arr[lo] > arr[hi] {
		swap(arr, lo, hi)
	}

	i, j, k, p, q := lo+1, lo+1, hi-1, arr[lo], arr[hi]
	for j <= k {
		if arr[j] < p {
			swap(arr, j, i)
			i++
		} else if arr[j] >= q {
			for arr[k] > q && j < k {
				k--
			}
			swap(arr, j, k)
			k--
			if arr[j] < p {
				swap(arr, j, i)
				i++
			}
		}
		j++
	}
	i--
	k++

	swap(arr, lo, i)
	swap(arr, hi, k)

	return i, k
}

func partition2V2(arr []int) (int, int) {
	lo, hi := 0, len(arr)-1
	if arr[lo] > arr[hi] {
		swap(arr, lo, hi)
	}

	p1, p2 := arr[lo], arr[hi]
	i, j := lo+1, hi-1
	g1 := i

	for i <= j {
		for arr[i] < p2 {
			if arr[i] > p1 {
				i++
			} else {
				swap(arr, i, g1)
				i++
				g1++
			}
		}
		for arr[j] > p2 {
			j--
		}

		if i == j {
			i++
		}
		if i < j {
			if arr[j] > p1 {
				// rotate a2 and a1 elements
				swap(arr, i, j)
				i++
				j--
			} else {
				// rotate a2 and a0 elements
				rotate(arr, i, g1, j)
				i++
				j--
				g1++
			}
		}
	}

	swap(arr, lo, g1-1)
	swap(arr, hi, j+1)

	return g1 - 1, j + 1
}
