package main

import "fmt"
import "golangbot/p16_structures/structs"  


func main()  {
	fmt.Println("Part 16: Structures")

	// 1. What is a structure?
	// structure is a user defined type which represents a collection of fields.

	// 2. Creating named structures
	emp1 := Employee{
        firstName: "Sam",
        age:       25,
        salary:    500,
        lastName:  "Anderson",
    }

    //creating structure without using field names
    emp2 := Employee{"Thomas", "Paul", 29, 800}

    fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)
	
	// 3. creating anonymous structures
	emp3 := struct {
        firstName, lastName string
        age, salary         int
    }{
        firstName: "Andreah",
        lastName:  "Nikola",
        age:       31,
        salary:    5000,
    }

	fmt.Println("Employee 3", emp3)
	
	// 4. Zero Values
	// ข้อมูลแต่ละ type หากไม่ได้มีการกำหนด value แต่ละ field จะมีแค่เปน zero values ของแต่ละ type 
	var emp4 Employee //zero valued structure
	fmt.Println("Employee 4", emp4) // Employee 4 {  0 0}  
	

	// 5. Accessing individual fields of a struct
	emp6 := Employee{"Sam", "Anderson", 55, 6000}
    fmt.Println("First Name:", emp6.firstName)
    fmt.Println("Last Name:", emp6.lastName)
    fmt.Println("Age:", emp6.age)
	fmt.Printf("Salary: $%d", emp6.salary)
	
	// และสามารถ create zero struct และ ค่อยกำหนดค่าให้ภายหลังได้



	// 6. Pointers to a struct
	emp8 := &Employee{"Sam", "Anderson", 55, 6000}
    fmt.Println("\nFirst Name:", (*emp8).firstName)
	fmt.Println("Age:", (*emp8).age)
	// ถึงแม้จะเป็น type pointer แต่ถ้าหากเป็น struct เราสามารถที่จะ อ้างอิงถึง field ภายในได้เลยโดยไม่ต้องใส่ (*emp8)
	fmt.Println("\nFirst Name:", emp8.firstName)
	fmt.Println("Age:", emp8.age)


	// 7. Anonymous fields
	// เราสามารถสร้าง Anonymous fields (ไม่ต้องระบุชื่อ field)
    p := Person1{"Naveen", 50}
	fmt.Println(p)
	
	// ถึงแม้ว่าเราจะไม่ต้องระบุ field แต่ default name ของ anonymous field ก็คือ type
	var p1 Person1
    p1.string = "naveen"
    p1.int = 50
	fmt.Println(p1)
	

	// 8. Nasted Structs
	var p8 Person
    p8.name = "Naveen"
    p8.age = 50
    p8.address = Address {
        city: "Chicago",
        state: "Illinois",
    }
    fmt.Println("Name:", p8.name)
    fmt.Println("Age:",p8.age)
    fmt.Println("City:",p8.address.city)
	fmt.Println("State:",p8.address.state)
	

	// 9. Promoted fields
	// กรณีที่ เป็น anonymous field (ไม่ระบุชื่อ field) จะถือว่าชื่อ field เป็นชื่อเดียวกับ class
	var p9 Person2
    p9.name = "Naveen"
    p9.age = 50
    p9.Address = Address{
        city:  "Chicago",
        state: "Illinois",
    }
    fmt.Println("Name:", p9.name)
    fmt.Println("Age:", p9.age)
    fmt.Println("City:", p9.city) //city is promoted field
	fmt.Println("State:", p9.state) //state is promoted field
	

	// 10. Exported Structs and Fields
	// import from another file.
	var spec computer.Spec
    spec.Maker = "apple"
    spec.Price = 50000
	fmt.Println("Spec:", spec)
	

	// 11. Structs Equality
	// Structs are value types ถ้า field แต่ละตัวมีค่าเท่ากัน สามารถใช้ == ได้เลย
	name1 := name{"Steve", "Jobs"}
    name2 := name{"Steve", "Jobs"}
    if name1 == name2 {
        fmt.Println("name1 and name2 are equal")
    } else {
        fmt.Println("name1 and name2 are not equal")
    }

    name3 := name{firstName:"Steve", lastName:"Jobs"}
    name4 := name{}
    name4.firstName = "Steve"
    if name3 == name4 {
        fmt.Println("name3 and name4 are equal")
    } else {
        fmt.Println("name3 and name4 are not equal")
	}
	

	// ถ้า struct มี field ที่เป็นไม่ใช้ cpmparable เช่น map
	// จะไม่สามารถ compare ได้
	// maps are not comparable
	
	// image1 := image{data: map[int]int{
    //     0: 155,
    // }}
    // image2 := image{data: map[int]int{
    //     0: 155,
    // }}
    // if image1 == image2 {
    //     fmt.Println("image1 and image2 are equal")
    // }
}


type image struct {  
    data map[int]int
}

type name struct {  
    firstName string
    lastName string
}


type Person1 struct {  
    string
    int
}

type Person struct {  
    name string
    age int
    address Address
}

type Person2 struct {  
    name string
    age int
    Address
}

type Address struct {  
    city, state string
}
// Employee class
type Employee struct {  
    firstName, lastName string
    age, salary         int
}