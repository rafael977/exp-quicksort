package sort

// QuickSort3P sorts the input array using 3 pivot quicksort
func QuickSort3P(arr []int) {
	if len(arr) <= 1 {
		return
	} else if len(arr) <= 2 {
		if arr[0] > arr[1] {
			swap(arr, 0, 1)
		}
	} else {
		p1, p2, p3 := partition3(arr)
		QuickSort3P(arr[:p1-2])
		QuickSort3P(arr[p1-1 : p2-1])
		QuickSort3P(arr[p2 : p3+1])
		QuickSort3P(arr[p3+2:])
	}
}

func partition3(arr []int) (int, int, int) {
	lo, hi := 0, len(arr)-1

	if arr[lo] > arr[lo+1] {
		swap(arr, lo, lo+1)
	}
	if arr[lo] > arr[hi] {
		rotate(arr, lo+1, hi, lo)
	} else if arr[lo+1] > arr[hi] {
		swap(arr, lo+1, hi)
	}

	i, j, k, l := lo+2, lo+2, hi-1, hi-1
	p, q, r := arr[lo], arr[lo+1], arr[hi]

	for j <= k {
		for arr[j] < q {
			if arr[j] < p {
				swap(arr, i, j)
				i++
			}
			j++
		}

		for arr[k] > q {
			if arr[k] > r {
				swap(arr, k, l)
				l--
			}
			k--
		}

		if j <= k {
			if arr[j] > r {
				if arr[k] < p {
					rotate(arr, j, i, k, l)
					i++
				} else {
					rotate(arr, j, k, l)
				}
				l--
			} else {
				if arr[k] < p {
					rotate(arr, j, i, k)
					i++
				} else {
					swap(arr, j, k)
				}
			}
			j++
			k--
		}
	}
	rotate(arr, lo+1, i-1, j-1)

	swap(arr, lo, i-2)
	swap(arr, hi, l+1)

	// fmt.Println(i, j, k, l, arr)
	return i, j, l
}
