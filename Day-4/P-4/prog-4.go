/*IV. Create a string "Go programming" and a string "GoLang".
  1. Perform copy and find valueOf (reflection methods)
*/

package main

import (
	"fmt"
	"reflect"
)

func main() {
    str1 := "Go programming"
    str2 := "GoLang"

    strCopy := str1

    valueOfStr1 := reflect.ValueOf(str1)
    valueOfStr2 := reflect.ValueOf(str2)
    valueOfStrCopy := reflect.ValueOf(strCopy)

    fmt.Println("Original String 1:", str1)
    fmt.Println("Original String 2:", str2)
    fmt.Println("Copied String:", strCopy)

    fmt.Println("ValueOf String 1:", valueOfStr1)
    fmt.Println("ValueOf String 2:", valueOfStr2)
    fmt.Println("ValueOf Copied String:", valueOfStrCopy)
}
