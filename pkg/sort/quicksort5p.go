package sort

import (
	"sort"
)

// QuickSort5P sorts the input array using 5 pivot quicksort
func QuickSort5P(arr []int) {
	if len(arr) <= 5 {
		sort.Ints(arr)
	} else {
		p1, p2, p3, p4, p5 := partition5(arr)
		QuickSort5P(arr[:p1-3])
		QuickSort5P(arr[p1-2 : p2-2])
		QuickSort5P(arr[p2-1 : p3-1])
		QuickSort5P(arr[p3 : p4+1])
		QuickSort5P(arr[p4+2 : p5+2])
		QuickSort5P(arr[p5+3:])
	}
}

func partition5(arr []int) (int, int, int, int, int) {
	lo, hi := 0, len(arr)-1

	var ps []int
	ps = append(ps, arr[lo], arr[lo+1], arr[lo+2], arr[hi-1], arr[hi])
	sort.Ints(ps)
	arr[lo], arr[lo+1], arr[lo+2], arr[hi-1], arr[hi] = ps[0], ps[1], ps[2], ps[3], ps[4]
	p1, p2, p3, p4, p5 := arr[lo], arr[lo+1], arr[lo+2], arr[hi-1], arr[hi]
	i1, i2, i3, i4, i5, i6 := lo+3, lo+3, lo+3, hi-2, hi-2, hi-2

	for i3 <= i4 {
		for arr[i3] < p3 && i3 <= i4 {
			if arr[i3] < p1 {
				rotate(arr, i3, i2, i1)
				i1++
				i2++
			} else {
				if arr[i3] < p2 {
					swap(arr, i2, i3)
					i2++
				}
			}
			i3++
		}
		for arr[i4] > p3 && i3 <= i4 {
			if arr[i4] > p5 {
				rotate(arr, i4, i5, i6)
				i6--
				i5--
			} else {
				if arr[i4] > p4 {
					swap(arr, i4, i5)
					i5--
				}
			}
			i4--
		}

		if i3 <= i4 {
			if arr[i3] < p4 {
				if arr[i4] > p2 {
					// a3 and a2 element
					swap(arr, i3, i4)
				} else {
					if arr[i4] > p1 {
						// a3 and a1
						rotate(arr, i3, i2, i4)
						i2++
					} else {
						// a3 and a0
						rotate(arr, i3, i2, i1, i4)
						i2++
						i1++
					}
				}
			} else {
				if arr[i3] < p5 {
					if arr[i4] > p2 {
						// a4 and a2 element
						rotate(arr, i3, i4, i5)
						i5--
					} else {
						if arr[i4] > p1 {
							// a4 and a1
							rotate(arr, i3, i2, i4, i5)
							i2++
							i5--
						} else {
							// a4 and a0
							rotate(arr, i3, i2, i1, i4, i5)
							i2++
							i1++
							i5--
						}
					}
				} else {
					if arr[i4] > p2 {
						// a5 and a2 element
						rotate(arr, i3, i4, i5, i6)
						i5--
						i6--
					} else {
						if arr[i4] > p1 {
							// a5 and a1
							rotate(arr, i3, i2, i4, i5, i6)
							i2++
							i5--
							i6--
						} else {
							// a5 and a0
							rotate(arr, i3, i2, i1, i4, i5, i6)
							i2++
							i1++
							i5--
							i6--
						}
					}
				}
			}
			i4--
			i3++
		}
	}

	rotate(arr, lo+2, i1-1, i2-1, i3-1)
	rotate(arr, lo+1, i1-2, i2-2)
	swap(arr, lo, i1-3)

	rotate(arr, hi-1, i6+1, i5+1)
	swap(arr, hi, i6+2)

	return i1, i2, i3, i5, i6
}
