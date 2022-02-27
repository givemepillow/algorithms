package rt

func FindMax(array []int, left int, right int) int {
	if left == right {
		return array[left]
	}
	mid := (left + right) / 2
	if array[left] > array[mid+1] {
		return FindMax(array, left, mid)
	} else {
		return FindMax(array, mid+1, right)
	}
}
