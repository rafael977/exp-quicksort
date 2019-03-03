package sort

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func rotate(arr []int, index ...int) {
	tmp := arr[index[0]]
	for i := range index {
		if i == len(index)-1 {
			arr[index[i]] = tmp
		} else {
			arr[index[i]] = arr[index[i+1]]
		}
	}
}
