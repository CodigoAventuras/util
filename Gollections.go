package util

func BinarySearch(items []int, l, r, k int) (index int) {
	if l > r { return -1 }

	mid := (l + r) / 2

	if items[mid] == k { return mid }

	if items[mid] < k {
		return BinarySearch(items, mid + 1, r, k)
	} else {
		return BinarySearch(items, l, mid - 1, k)
	}
}
