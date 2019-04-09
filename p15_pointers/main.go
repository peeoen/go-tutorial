package main

import (
	"fmt"
)

func main() {
	fmt.Println("Part 15: Pointers")

	// 1. What is a pointer?
	// pointer is a variable which stores the memory address of another variable

	// 2. Declaring pointers
	// *T คือ type ของ pointer variable ซึ่งจะชี้ไปยัง value ของ type T
	// & operator จะใช้ในการอ้างอิงถึง address ของตัวแปร
	b := 255
    var a *int = &b
    fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)
	

	// 3. Zero value of a pointer
	// The zero value of a pointer is "nil"
	a3 := 25
    var b3 *int
    if b3 == nil {
        fmt.Println("b3 is", b3)
        b3 = &a3
        fmt.Println("b3 after initialization is", b3)
	}
	

	// 4. Deferencing a pointer
	// การเข้าถึง value จาก pointer โดยใช้ *a syntax สำหรับการ deference
	b4 := 255
    a4 := &b4
    fmt.Println("address of b4 is", a4)
	fmt.Println("value of b4 is", *a4)
	//	นอกจากนี้ยังสามารถแก้ไข value จาก pointer ได้
	*a4++
	fmt.Println("new value of b is", b4)
	

	// 5. Passing pointer to a function
	a5 := 58
    fmt.Println("value of a before function call is",a5)
    b5 := &a5
    change(b5)
    fmt.Println("value of a after function call is", a5)

}

func change(val *int) {  
    *val = 55
}