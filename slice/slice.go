package slice

import "fmt"

// Normal array initialization :
// var arr [5]int
func TryingSliceCode() {
	var arr = [5]int{1, 2, 3, 4}
	var slice = []int{1, 3, 4, 5, 6, 7}
	fmt.Printf("This is arr %v , and this is slice %v", arr, slice)
	var newSlice = make([]int, 8, cap(slice)*2)
	itemsCopied := copy(newSlice, slice)
	fmt.Printf("Items Copied %d", itemsCopied)
	updatedSlice := append(newSlice, 19990, 2389, 2321, 4324732)
	fmt.Printf("Slice  %v , Previous Slice %v", updatedSlice, newSlice)
	// testSlice:=[]int{newSlice...}
}

func SumAll(lists ...[]int) {
	fmt.Println(lists)
}
