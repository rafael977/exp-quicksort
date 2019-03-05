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

	ps := append(make([]int, 0),
		arr[lo], arr[lo+1], arr[lo+2], arr[lo+3], arr[hi-2], arr[hi-1], arr[hi])
	sort.Ints(ps)
	arr[lo], arr[lo+1], arr[lo+2], arr[lo+3], arr[hi-2], arr[hi-1], arr[hi] =
		ps[0], ps[1], ps[2], ps[3], ps[4], ps[5], ps[6]

	p1, p2, p3, p4, p5, p6, p7 := arr[lo], arr[lo+1], arr[lo+2], arr[lo+3], arr[hi-2], arr[hi-1], arr[hi]
	i, j := lo+4, hi-3
	g1, g2, g3, g4, g5, g6 := i, i, i, j, j, j

	for i <= j {
		for arr[i] < p4 {
			if arr[i] > p3 {
				i++
			} else if arr[i] > p2 {
				swap(arr, i, g3)
				i++
				g3++
			} else if arr[i] > p1 {
				rotate(arr, i, g3, g2)
				i++
				g3++
				g2++
			} else {
				rotate(arr, i, g3, g2, g1)
				i++
				g3++
				g2++
				g1++
			}
		}
		for arr[j] > p4 {
			if arr[j] < p5 {
				j--
			} else if arr[j] < p6 {
				swap(arr, j, g4)
				j--
				g4--
			} else if arr[j] < p7 {
				rotate(arr, j, g4, g5)
				j--
				g4--
				g5--
			} else {
				rotate(arr, j, g4, g5, g6)
				j--
				g4--
				g5--
				g6--
			}
		}

		if i == j {
			j--
		} else if i < j {
			if arr[i] < p5 {
				if arr[j] > p3 {
					// rotate a4 and a3 elements
					swap(arr, i, j)
					i++
					j--
				} else if arr[j] > p2 {
					// rotate a4 and a2 elements
					rotate(arr, i, g3, j)
					i++
					j--
					g3++
				} else if arr[j] > p1 {
					// rotate a4 and a1 elements
					rotate(arr, i, g3, g2, j)
					i++
					j--
					g3++
					g2++
				} else {
					// rotate a4 and a0 elements
					rotate(arr, i, g3, g2, g1, j)
					i++
					j--
					g3++
					g2++
					g1++
				}
			} else if arr[i] < p6 {
				if arr[j] > p3 {
					// rotate a5 and a3 elements
					rotate(arr, i, j, g4)
					i++
					j--
					g4--
				} else if arr[j] > p2 {
					// rotate a5 and a2 elements
					rotate(arr, i, g3, j, g4)
					i++
					j--
					g3++
					g4--
				} else if arr[j] > p1 {
					// rotate a5 and a1 elements
					rotate(arr, i, g3, g2, j, g4)
					i++
					j--
					g3++
					g2++
					g4--
				} else {
					// rotate a5 and a0 elements
					rotate(arr, i, g3, g2, g1, j, g4)
					i++
					j--
					g3++
					g2++
					g1++
					g4--
				}
			} else if arr[i] < p7 {
				if arr[j] > p3 {
					// rotate a6 and a3 elements
					rotate(arr, i, j, g4, g5)
					i++
					j--
					g4--
					g5--
				} else if arr[j] > p2 {
					// rotate a6 and a2 elements
					rotate(arr, i, g3, j, g4, g5)
					i++
					j--
					g3++
					g4--
					g5--
				} else if arr[j] > p1 {
					// rotate a6 and a1 elements
					rotate(arr, i, g3, g2, j, g4, g5)
					i++
					j--
					g3++
					g2++
					g4--
					g5--
				} else {
					// rotate a6 and a0 elements
					rotate(arr, i, g3, g2, g1, j, g4, g5)
					i++
					j--
					g3++
					g2++
					g1++
					g4--
					g5--
				}
			} else {
				if arr[j] > p3 { // rotate a7 and a3 elements
					rotate(arr, i, j, g4, g5, g6)
					i++
					j--
					g4--
					g5--
					g6--
				} else if arr[j] > p2 {
					// rotate a7 and a2 elements
					rotate(arr, i, g3, j, g4, g5, g6)
					i++
					j--
					g3++
					g4--
					g5--
					g6--
				} else if arr[j] > p1 {
					// rotate a7 and a1 elements
					rotate(arr, i, g3, g2, j, g4, g5, g6)
					i++
					j--
					g3++
					g2++
					g4--
					g5--
					g6--
				} else {
					// rotate a7 and a0 elements
					rotate(arr, i, g3, g2, g1, j, g4, g5, g6)
					i++
					j--
					g3++
					g2++
					g1++
					g4--
					g5--
					g6--
				}
			}
		}
	}

	rotate(arr, lo+3, g1-1, g2-1, g3-1, i-1)
	rotate(arr, lo+2, g1-2, g2-2, g3-2)
	rotate(arr, lo+1, g1-3, g2-3)
	swap(arr, lo, g1-4)

	rotate(arr, hi-2, g6+1, g5+1, g4+1)
	rotate(arr, hi-1, g6+2, g5+2)
	swap(arr, hi, g6+3)

	return g1, g2, g3, i, g4, g5, g6
}
