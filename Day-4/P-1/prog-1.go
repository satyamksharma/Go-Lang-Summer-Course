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

func Reverse(slice interface{}) {
	value := reflect.ValueOf(slice)
	if value.Kind() != reflect.Slice {
		panic("Reverse: not a slice")
	}

	length := value.Len()
	swap := reflect.Swapper(slice)
	for i := 0; i < length/2; i++ {
		swap(i, length-i-1)
	}
}

func main() {
	intSlice := []int{1, 2, 3, 4, 5}
	floatSlice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	stringSlice := []string{"a", "b", "c", "d", "e"}

	fmt.Println("Are intSlice and stringSlice equal?", reflect.DeepEqual(intSlice, stringSlice))
	fmt.Println("Are intSlice and floatSlice equal?", reflect.DeepEqual(intSlice, floatSlice))
	fmt.Println("Are floatSlice and stringSlice equal?", reflect.DeepEqual(floatSlice, stringSlice))

	fmt.Printf("Size of intSlice: %d bytes\n", unsafe.Sizeof(intSlice))
	fmt.Printf("Size of floatSlice: %d bytes\n", unsafe.Sizeof(floatSlice))
	fmt.Printf("Size of stringSlice: %d bytes\n", unsafe.Sizeof(stringSlice))

	fmt.Printf("Number of elements in intSlice: %d\n", len(intSlice))
	fmt.Printf("Number of elements in floatSlice: %d\n", len(floatSlice))
	fmt.Printf("Number of elements in stringSlice: %d\n", len(stringSlice))

	Reverse(intSlice)
	Reverse(floatSlice)
	Reverse(stringSlice)
	fmt.Println("Reversed intSlice:", intSlice)
	fmt.Println("Reversed floatSlice:", floatSlice)
	fmt.Println("Reversed stringSlice:", stringSlice)
}
