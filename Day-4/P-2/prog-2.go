/*II. Create a structure of PES with the following data vals:
a. Courses offered
b. Rank
c. Location

1. Access the structure members
2. Find the size of structure
3. Perform Alignof and Offsetof the struct members
*/

package main

import (
	"fmt"
	"unsafe"
)

type PES struct {
    CoursesOffered []string
    Rank           int
    Location       string
}

func main() {
    pes := PES{
        CoursesOffered: []string{"Engineering", "Science", "Arts"},
        Rank:           1,
        Location:       "Bangalore",
    }

    fmt.Println("Courses Offered:", pes.CoursesOffered)
    fmt.Println("Rank:", pes.Rank)
    fmt.Println("Location:", pes.Location)

    fmt.Printf("Size of PES structure: %d bytes\n", unsafe.Sizeof(pes))

    fmt.Printf("Alignment of CoursesOffered: %d\n", unsafe.Alignof(pes.CoursesOffered))
    fmt.Printf("Offset of CoursesOffered: %d\n", unsafe.Offsetof(pes.CoursesOffered))
    
    fmt.Printf("Alignment of Rank: %d\n", unsafe.Alignof(pes.Rank))
    fmt.Printf("Offset of Rank: %d\n", unsafe.Offsetof(pes.Rank))
    
    fmt.Printf("Alignment of Location: %d\n", unsafe.Alignof(pes.Location))
    fmt.Printf("Offset of Location: %d\n", unsafe.Offsetof(pes.Location))
}
