package slices

// Your only reference:
// https://blog.golang.org/go-slices-usage-and-internals

// Insert should add the value 'value' at the index 'index' of the slice 'slice'
func Insert(slice []int, index, value int) []int {
	return append(slice[:index], append([]int{value}, slice[index:]...)...)
}

// Delete should remove the index item from slice
func Delete(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
