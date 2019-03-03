package sort

// QuickSort2P sorts input array with 2 pivot quickdsort
func QuickSort2P(arr []int) {
	if len(arr) > 1 {
		lp, rp := partition2(arr)
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
