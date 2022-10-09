package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	withData := []int{2, 6, 9}
	return withData
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index >= len(slice) || index < 0 {
		return -1
	}
	x := slice[index]
	return x

}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index >= len(slice) || index < 0 {
		return append(slice, value)
	}
	slice[index] = value
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	var emptySet []int
	filedSet := append(emptySet, values...)
	return append(filedSet, slice...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index == 0 {
		return slice[1:]
	} else if len(slice) <= index || index < 0 {
		return slice
	} else {
		firstSlice := slice[0:index]
		secondSlice := slice[index+1:]
		return append(firstSlice, secondSlice...)

	}

}

func find(num int, nums ...int) int {

	for i, v := range nums {
		if v == num {
			return i
		}
	}
	return -1
}
