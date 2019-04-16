package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Part 17: Methods")

	emp1 := Employee{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	emp1.displaySalary() //Calling displaySalary() method of Employee type
	displaySalary2(emp1)

	r := Rectangle{
		length: 10,
		width:  5,
	}
	fmt.Printf("Area of rectangle %d\n", r.Area())
	c := Circle{
		radius: 12,
	}
	fmt.Printf("Area of circle %f\n", c.Area())

	// pointer receive vs value receivers
	e := Employee{
		name: "Mark Andrew",
		age:  50,
	}
	fmt.Printf("Employee name before change: %s", e.name)
	e.changeName("Michael Andrew")
	fmt.Printf("\nEmployee name after change: %s", e.name)

	fmt.Printf("\n\nEmployee age before change: %d", e.age)
	(&e).changeAge(51)
	fmt.Printf("\nEmployee age after change: %d", e.age)
	// ถ้า function ที่ call ตั้งแต่เป็น pointer receiver ไม่จำเป็นต้องใส่ (&e) เพราะว่า go จะตีความว่า instance ที่ไป call function นั้นคือ pointer
	e.changeAge(55)
	fmt.Printf("\nEmployee age after change: %d\n", e.age)
	


	// Methods of anonymous fields
	p := person{
        firstName: "Elon",
        lastName:  "Musk",
        address: address {
            city:  "Los Angeles",
            state: "California",
        },
    }

	// same output
	p.address.fullAddress();
    p.fullAddress() //accessing fullAddress method of address struct
}

// method on struct
func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

// normal method
func displaySalary2(e Employee) {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

type Employee struct {
	name     string
	salary   int
	currency string
	age      int
}

type Rectangle struct {
	length int
	width  int
}

type Circle struct {
	radius float64
}

// method can same name but different types
func (r Rectangle) Area() int {
	return r.length * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

/*
Method with value receiver
*/
func (e Employee) changeName(newName string) {
	e.name = newName
}

/*
Method with pointer receiver
*/
func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}



	// Methods of anonymous fields
type address struct {  
    city  string
    state string
}

func (a address) fullAddress() {  
    fmt.Printf("Full address: %s, %s\n", a.city, a.state)
}

type person struct {  
    firstName string
    lastName  string
    address
}