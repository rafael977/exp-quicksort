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

	ps := append(make([]int, 0), arr[lo], arr[lo+1], arr[lo+2], arr[hi-1], arr[hi])
	sort.Ints(ps)
	arr[lo], arr[lo+1], arr[lo+2], arr[hi-1], arr[hi] = ps[0], ps[1], ps[2], ps[3], ps[4]
	p1, p2, p3, p4, p5 := arr[lo], arr[lo+1], arr[lo+2], arr[hi-1], arr[hi]
	i, j := lo+3, hi-2
	g1, g2, g3, g4 := i, i, j, j

	for i <= j {
		for arr[i] < p3 {
			if arr[i] > p2 {
				i++
			} else if arr[i] > p1 {
				swap(arr, i, g2)
				i++
				g2++
			} else {
				rotate(arr, i, g2, g1)
				i++
				g2++
				g1++
			}
		}
		for arr[j] > p3 {
			if arr[j] < p4 {
				j--
			} else if arr[j] < p5 {
				swap(arr, j, g3)
				j--
				g3--
			} else {
				rotate(arr, j, g3, g4)
				j--
				g3--
				g4--
			}
		}

		if i == j {
			j--
		} else if i < j {
			if arr[i] < p4 {
				if arr[j] > p2 {
					// rotate a3 and a2 elements
					swap(arr, i, j)
					i++
					j--
				} else if arr[j] > p1 {
					// rotate a3 and a1 elements
					rotate(arr, i, g2, j)
					i++
					j--
					g2++
				} else {
					// rotate a3 and a0 elements
					rotate(arr, i, g2, g1, j)
					i++
					j--
					g2++
					g1++
				}
			} else if arr[i] < p5 {
				if arr[j] > p2 {
					// rotate a4 and a2 elements
					rotate(arr, i, j, g3)
					i++
					j--
					g3--
				} else if arr[j] > p1 {
					// rotate a4 and a1 elements
					rotate(arr, i, g2, j, g3)
					i++
					j--
					g2++
					g3--
				} else {
					// rotate a4 and a0 elements
					rotate(arr, i, g2, g1, j, g3)
					i++
					j--
					g2++
					g1++
					g3--
				}
			} else {
				if arr[j] > p2 {
					// rotate a5 and a2 elements
					rotate(arr, i, j, g3, g4)
					i++
					j--
					g3--
					g4--
				} else if arr[j] > p1 {
					// rotate a5 and a1 elements
					rotate(arr, i, g2, j, g3, g4)
					i++
					j--
					g2++
					g3--
					g4--
				} else {
					// rotate a5 and a0 elements
					rotate(arr, i, g2, g1, j, g3, g4)
					i++
					j--
					g2++
					g1++
					g3--
					g4--
				}
			}
		}
	}

	rotate(arr, lo+2, g1-1, g2-1, i-1)
	rotate(arr, lo+1, g1-2, g2-2)
	swap(arr, lo, g1-3)

	rotate(arr, hi-1, g4+1, g3+1)
	swap(arr, hi, g4+2)

	// fmt.Println(g1, g2, i, j, g3, g4, arr)
	return g1, g2, i, g3, g4
}
