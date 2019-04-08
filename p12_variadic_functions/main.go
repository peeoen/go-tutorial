package main

import (
	"fmt"
)

func main() {
	fmt.Println("Part 12: Variadic Functions")

	// Variadic Functions เป็น function ที่สามารถระบุ parameters ได้ไม่รู้จบ

	find(89, 89, 90, 95)
    find(45, 56, 67, 45, 90, 109)
    find(78, 38, 56, 98)
	find(87)
	
	// หรืออาจจะระบุไปทั้ง array โดยใช้ ... เลยก็ได้
	nums := []int{89, 90, 95}
    find(89, nums...)
}


func find(num int, nums ...int) {  
    fmt.Printf("type of nums is %T\n", nums)
    found := false
    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            found = true
        }
    }
    if !found {
        fmt.Println(num, "not found in ", nums)
    }
    fmt.Printf("\n")
}