/*I. Create different arrays or slices of different data types.
1. check if slices are equal using deepequal in reflection
2. Sizeof any slices
3. Number of elements in the array or slices
4. reverse the given array
*/

package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func reverseSlice(slice []int) []int {
    for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
        slice[i], slice[j] = slice[j], slice[i]
    }
    return slice
}

func main() {
    arrayInt := [5]int{1, 2, 3, 4, 5}
    slice1 := []int{1, 2, 3, 4, 5}
    slice2 := []int{1, 2, 3, 4, 6}
    sliceString := []string{"one", "two", "three"}

    fmt.Println("Slices equal (sliceInt1, sliceInt2):", reflect.DeepEqual(slice1, slice2))
    fmt.Println("Slices equal (sliceInt1, sliceInt1):", reflect.DeepEqual(slice1, slice1))

    fmt.Printf("Size of sliceInt1: %d bytes\n", unsafe.Sizeof(slice1))
    fmt.Printf("Size of sliceString: %d bytes\n", unsafe.Sizeof(sliceString))

    fmt.Printf("Number of elements in arrayInt: %d\n", len(arrayInt))
    fmt.Printf("Number of elements in sliceInt1: %d\n", len(slice1))
    fmt.Printf("Number of elements in sliceString: %d\n", len(sliceString))

    reversedSlice := reverseSlice(slice1)
    fmt.Println("Reversed sliceInt1:", reversedSlice)
}

