package sort

import "math/rand"

// GenPerm generates permutation of n
func GenPerm(n int) []int {
	return rand.Perm(n)
}

// GenRand generates array of random int of size n
func GenRand(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}

// GenSeq generates sorted int array of size n
func GenSeq(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
}

// GenRev generates reversed int array of size n
func GenRev(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = n - i - 1
	}
	return arr
}
