package bst

// BSTArray menyimpan elemen BST dalam array yang sudah terurut
type BSTArray[T comparable] struct {
	data []T
}

// NewBSTArray membuat BST dari array yang sudah terurut
func NewBSTArray[T comparable](sortedData []T) *BSTArray[T] {
	return &BSTArray[T]{data: sortedData}
}

// Search melakukan Binary Search di BST yang disimpan dalam array
func (bst *BSTArray[T]) Search(target T, less func(a, b T) bool, equal func(a, b T) bool) *T {
	left, right := 0, len(bst.data)-1

	for left <= right {
		mid := left + (right-left)/2

		if equal(bst.data[mid], target) { // Menggunakan fungsi equal untuk pengecekan
			return &bst.data[mid]
		} else if less(target, bst.data[mid]) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return nil // Jika tidak ditemukan
}
