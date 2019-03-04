package sort

import (
	"sort"
)

// QuickSort7P is basic implementation of 7-pivot quicksort
func QuickSort7P(arr []int) {
	if len(arr) <= 7 {
		sort.Ints(arr)
	} else {
		p1, p2, p3, p4, p5, p6, p7 := partition7(arr)
		QuickSort7P(arr[:p1-4])
		QuickSort7P(arr[p1-3 : p2-3])
		QuickSort7P(arr[p2-2 : p3-2])
		QuickSort7P(arr[p3-1 : p4-1])
		QuickSort7P(arr[p4 : p5+1])
		QuickSort7P(arr[p5+2 : p6+2])
		QuickSort7P(arr[p6+3 : p7+3])
		QuickSort7P(arr[p7+4:])
	}
}

func partition7(arr []int) (int, int, int, int, int, int, int) {
	lo, hi := 0, len(arr)-1

	var ps []int
	ps = append(ps, arr[lo], arr[lo+1], arr[lo+2], arr[lo+3], arr[hi-2], arr[hi-1], arr[hi])
	sort.Ints(ps)
	arr[lo], arr[lo+1], arr[lo+2], arr[lo+3], arr[hi-2], arr[hi-1], arr[hi] =
		ps[0], ps[1], ps[2], ps[3], ps[4], ps[5], ps[6]

	p1, p2, p3, p4, p5, p6, p7 := arr[lo], arr[lo+1], arr[lo+2], arr[lo+3], arr[hi-2], arr[hi-1], arr[hi]
	i1, i2, i3, i4, i5, i6, i7, i8 := lo+4, lo+4, lo+4, lo+4, hi-3, hi-3, hi-3, hi-3

	for i4 <= i5 {
		for arr[i4] < p4 && i4 <= i5 {
			if arr[i4] < p2 {
				if arr[i4] < p1 {
					rotate(arr, i4, i3, i2, i1)
					i1++
				} else {
					rotate(arr, i4, i3, i2)
				}
				i2++
				i3++
			} else {
				if arr[i4] < p3 {
					swap(arr, i4, i3)
					i3++
				}
			}
			i4++
		}

		for arr[i5] > p4 && i4 <= i5 {
			if arr[i5] > p6 {
				if arr[i5] > p7 {
					rotate(arr, i5, i6, i7, i8)
					i8--
				} else {
					rotate(arr, i5, i6, i7)
				}
				i6--
				i7--
			} else {
				if arr[i5] > p5 {
					swap(arr, i5, i6)
					i6--
				}
			}
			i5--
		}

		if i4 <= i5 {
			if arr[i4] < p6 {
				if arr[i4] < p5 {
					if arr[i5] < p2 {
						if arr[i5] < p1 {
							// a4 and a0
							rotate(arr, i4, i3, i2, i1, i5)
							i1++
							i2++
							i3++
						} else {
							// a4 and a1
							rotate(arr, i4, i3, i2, i5)
							i2++
							i3++
						}
					} else {
						if arr[i5] < p3 {
							// a4 and a2
							rotate(arr, i4, i3, i5)
							i3++
						} else {
							// a4 and a3
							swap(arr, i4, i5)
						}
					}
				} else {
					if arr[i5] < p2 {
						if arr[i5] < p1 {
							// a5 and a0
							rotate(arr, i4, i3, i2, i1, i5, i6)
							i1++
							i2++
							i3++
							i6--
						} else {
							// a5 and a1
							rotate(arr, i4, i3, i2, i5, i6)
							i2++
							i3++
							i6--
						}
					} else {
						if arr[i5] < p3 {
							// a5 and a2
							rotate(arr, i4, i3, i5, i6)
							i3++
							i6--
						} else {
							// a5 and a3
							rotate(arr, i4, i5, i6)
							i6--
						}
					}
				}
			} else {
				if arr[i4] < p7 {
					if arr[i5] < p2 {
						if arr[i5] < p1 {
							// a6 and a0
							rotate(arr, i4, i3, i2, i1, i5, i6, i7)
							i1++
							i2++
							i3++
							i6--
							i7--
						} else {
							// a6 and a1
							rotate(arr, i4, i3, i2, i5, i6, i7)
							i2++
							i3++
							i6--
							i7--
						}
					} else {
						if arr[i5] < p3 {
							// a6 and a2
							rotate(arr, i4, i3, i5, i6, i7)
							i3++
							i6--
							i7--
						} else {
							// a6 and a3
							rotate(arr, i4, i5, i6, i7)
							i6--
							i7--
						}
					}
				} else {
					if arr[i5] < p2 {
						if arr[i5] < p1 {
							// a7 and a0
							rotate(arr, i4, i3, i2, i1, i5, i6, i7, i8)
							i1++
							i2++
							i3++
							i6--
							i7--
							i8--
						} else {
							// a7 and a1
							rotate(arr, i4, i3, i2, i5, i6, i7, i8)
							i2++
							i3++
							i6--
							i7--
							i8--
						}
					} else {
						if arr[i5] < p3 {
							// a7 and a2
							rotate(arr, i4, i3, i5, i6, i7, i8)
							i3++
							i6--
							i7--
							i8--
						} else {
							// a7 and a3
							rotate(arr, i4, i5, i6, i7, i8)
							i6--
							i7--
							i8--
						}
					}
				}
			}
			i4++
			i5--
		}
	}

	rotate(arr, lo+3, i1-1, i2-1, i3-1, i4-1)
	rotate(arr, lo+2, i1-2, i2-2, i3-2)
	rotate(arr, lo+1, i1-3, i2-3)
	swap(arr, lo+0, i1-4)

	rotate(arr, hi-2, i8+1, i7+1, i6+1)
	rotate(arr, hi-1, i8+2, i7+2)
	swap(arr, hi-0, i8+3)

	return i1, i2, i3, i4, i6, i7, i8
}
