/*
III. Create an embedded structure - Company{name, location}
    1. Access the embedded structure members.
    2. Access structure members
    3. Accessing structure members randomly.
    4. Find the structure members are equal or not (deepequal)
*/

package main

import (
	"fmt"
	"reflect"
)

type Company struct {
    Name     string
    Location string
}

type Employee struct {
    Company
    EmployeeName string
    Position     string
}

func main() {
    emp1 := Employee{
        Company: Company{
            Name:     "TechCorp",
            Location: "New York",
        },
        EmployeeName: "John Doe",
        Position:     "Software Engineer",
    }

    emp2 := Employee{
        Company: Company{
            Name:     "TechCorp",
            Location: "New York",
        },
        EmployeeName: "Jane Smith",
        Position:     "Project Manager",
    }

    fmt.Println("Company Name:", emp1.Name)
    fmt.Println("Company Location:", emp1.Location)

    fmt.Println("Employee Name:", emp1.EmployeeName)
    fmt.Println("Employee Position:", emp1.Position)

    fmt.Println("Random Access - Employee Position:", emp2.Position)
    fmt.Println("Random Access - Company Location:", emp2.Company.Location)

    fmt.Println("Are emp1 and emp2 Companies equal?", reflect.DeepEqual(emp1.Company, emp2.Company))
    fmt.Println("Are emp1 and emp2 Employees equal?", reflect.DeepEqual(emp1, emp2))
}
