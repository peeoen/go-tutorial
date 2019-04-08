package main

import (
	"fmt"
)

func main() {
	fmt.Println("Part 13: Maps")

	// 1. What is a map ?
	// map is a buitin type in Go which associates a value to a key. The value can be retrived using the corresponding key
	// Like a Dictionary in C#

	// make(map[type of key]type of value)

	// zero value of map is "nil"
	personSalary := make(map[string]int)
    personSalary["steve"] = 12000
    personSalary["jamie"] = 15000
    personSalary["mike"] = 9000
	fmt.Println("personSalary map contents:", personSalary)
	

	// หรืออาจะกำหนดค่าตอนที่สร้างตัวแปรเลย
	personSalary1 := map[string]int {
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary1["mike"] = 9000
	fmt.Println("personSalary1 map contents:", personSalary1)
	

	// 2. Accessing items of a map
	personSalary2 := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary2["mike"] = 9000
    employee := "jamie"
	fmt.Println("Salary of", employee, "is", personSalary2[employee])
	
	
	// ในกรณีหา key ไม่เจอใน map มันจะส่งเป็น zero value กลับมา
	// กรณีนี้หากไม่มีจะส่ง 0 กลับมาเพราะ value type คือ int
	personSalary3 := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary3["mike"] = 9000
    employee3 := "jamie"
    fmt.Println("Salary of", employee, "is", personSalary3[employee3])
	fmt.Println("Salary of joe is", personSalary3["joe"]) // joe will return 0
	

	// ถ้าหากจะดู state ของ map ว่ามีหรือไม่ go จะ return ข้อมูล status กลับมาอีกตัว (ตัวแปร ok)
	// ซึ่งสามารถตรวจสอบได้จากตัวแปนดังกล่าว
	personSalary4 := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary4["mike"] = 9000
    newEmp := "joe"
    value, ok := personSalary4[newEmp]
    if ok == true {
        fmt.Println("Salary of", newEmp, "is", value)
    } else {
        fmt.Println(newEmp,"not found")
	}
	

	// 3. range 
	// สามารถวนลูปเพื่อดูข้อมูลของ map แต่ละตัวได้

	personSalary5 := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary5["mike"] = 9000
    fmt.Println("All items of a map")
    for key, value := range personSalary5 {
        fmt.Printf("personSalary5[%s] = %d\n", key, value)
	}
	

	// 4. deleting items
    personSalary6 := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary6["mike"] = 9000
    fmt.Println("map before deletion", personSalary6)
    delete(personSalary6, "steve")
    fmt.Println("map after deletion", personSalary6)


	// 5. length of map
	personSalary7 := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary7["mike"] = 9000
	fmt.Println("length is", len(personSalary7))
	
	// 6. Maps are reference types
	personSalary8 := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary8["mike"] = 9000
    fmt.Println("Original person salary", personSalary8)
    newPersonSalary8 := personSalary8
    newPersonSalary8["mike"] = 18000
	fmt.Println("Person salary changed", personSalary8)
	

	// 7. Maps equality 
	// map ไม่สามารถตรวจสอบโดยใช้ == ได้
	// got error: invalid operation: map1 == map2 (map can only be compared to nil)
	// map1 := map[string]int{
    //     "one": 1,
    //     "two": 2,
    // }

    // map2 := map1

    // if map1 == map2 {
    // }
}

