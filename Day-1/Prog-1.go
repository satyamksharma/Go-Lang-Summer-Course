package main

import "fmt"

type Rectangle struct{
	len int
	wid int
}

var (
	v1 = Rectangle{12,4}
	v2 = Rectangle{13,7}
	v3 = Rectangle{16, 9}
	v4 = Rectangle{}
)

func main() {
	a := Rectangle{12, 4}
	area := a.len * a.wid

	fmt.Println(v1, v2, v3, v4)

	fmt.Println(a)

	fmt.Printf("Area of the Rectangle is %v", area)

}