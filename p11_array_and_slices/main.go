package main

import (
	"fmt"
)

func main() {
	fmt.Println("Part 11: Array and Slices")

	// 1. Specific length
	a := [3]int { 12, 78 ,50}

	fmt.Println(a)

	// 2. ignore length of the array with (...)

	b := [...]int{12, 78, 50, 10 ,9 ,8} // ... makes the compiler determine the length
	fmt.Println(b)
	

	// 3. array cann't be resize 

	// c := [3]int{5, 78, 8}
    // var d [5]int
    // d = c  //not possible since [3]int and [5]int are distinct types

	// ใช้ ... เพิ่มเป็นการไม่ระบุจำนวนของ array


	c := [...]string{"USA", "China", "India", "Germany", "France"}
    d := c // c copy of c is assigned to d
    d[0] = "Singapore"
    fmt.Println("c is ", c)
	fmt.Println("d is ", d) 
	

	// 4. array of go is value types not reference types. **
	num := [...]int{5, 6, 7, 8, 8}
    fmt.Println("before passing to function ", num)
    changeLocal(num) //num is passed by value
	fmt.Println("after passing to function ", num)
	
	// 5. length of array
    e := [...]float64{67.7, 89.8, 21, 78}
	fmt.Println("length of e is",len(e))
	

	// 6. iterating array using range
	f := [...]float64{67.7, 89.8, 21, 78}
    for g := 0; g < len(f); g++ { //looping from 0 to the length of the array
        fmt.Printf("%d th element of f is %.2f\n", g, f[g])
	}
	

	h6 := [...]float64{67.7, 89.8, 21, 78}
    sum6 := float64(0)
    for i6, v6 := range h6 {//range returns both the index and value
        fmt.Printf("%d the element of h is %.2f\n", i6, v6)
        sum6 += v6
    }
	fmt.Println("\nsum of all elements of h",sum6)
	
	// or ignore index with _ blank identifier
	// for _, v := range a { //ignores index  
	// }


	// 7. Multidimensional Arrays
	a7 := [3][2]string{
        {"lion", "tiger"},
        {"cat", "dog"},
        {"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
    }
    printarray(a7)
    var b7 [3][2]string
    b7[0][0] = "apple"
    b7[0][1] = "samsung"
    b7[1][0] = "microsoft"
    b7[1][1] = "google"
    b7[2][0] = "AT&T"
    b7[2][1] = "T-Mobile"
    fmt.Printf("\n")
	printarray(b7)
	


	// Slices จะ reference กับ array ตัวเดิมไมไ่ด้ clone ขึ้นมาใหม่
	// 8. Create a slice
	a8 := [5]int{76, 77, 78, 79, 80}
    var b8 []int = a8[1:4] //creates a slice from a[1] to a[3]
	fmt.Println(b8)
	
	// 9. Modify slic
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
    dslice := darr[2:5]
    fmt.Println("array before",darr)
    for i := range dslice {
        dslice[i]++
    }
	fmt.Println("array after",darr) 
	


	// [:] เป็นการเลือกข้อมูลทั้งจาก array ต้นทาง
	numa := [3]int{78, 79 ,80}
    nums1 := numa[:] //creates a slice which contains all elements of the array
    nums2 := numa[:]
    fmt.Println("array before change 1",numa)
    nums1[0] = 100
    fmt.Println("array after modification to slice nums1", numa)
    nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)
	
	// 10. Length and capacity of a slice
	// len คือจำนวน element ที่ได้จาก slice
	// cap คือจำนวน element ที่ได้จาก array ต้นทาง
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
    fruitslice := fruitarray[1:3]
	fmt.Printf("-->length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice))
	
	// slice สามารถ re-sliced ได้
	fruitarray2 := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
    fruitslice2 := fruitarray2[1:3]
    fmt.Printf("length of slice %d capacity %d\n", len(fruitslice2), cap(fruitslice2)) //length of is 2 and capacity is 6
    fruitslice2 = fruitslice2[:cap(fruitslice2)] //re-slicing furitslice till its capacity
	fmt.Println("After re-slicing length is",len(fruitslice2), "and capacity is",cap(fruitslice2))
	

	// 11. Create a slice using make
	// func make([]T, len, cap) ระบุ len กับ cap ไป
	i11 := make([]int, 5, 5)
	fmt.Println(i11)
	
	// 12. Appending to a slice
	// func append(s []T, x ...T) []T
	cars := []string{"Ferrari", "Honda", "Ford"}
    fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
    cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6
	fmt.Println(cars)

	// zero values of array is nil
	// A nil slice has length and capacity 0. It is possible to append values to a nil slice using the append function.
	var names []string //zero value of a slice is nil
    if names == nil {
        fmt.Println("slice is nil going to append")
        names = append(names, "John", "Sebastian", "Vinay")
        fmt.Println("names contents:",names)
	}
	
	veggies := []string{"potatoes","tomatoes","brinjal"}
    fruits := []string{"oranges","apples"}
    food := append(veggies, fruits...)
	fmt.Println("food:",food)
	


	// 13. Passing a slice to a function
	// slice is reference types. even though it's passed by value, the pointer variable will refer to the same underlying array. 
	nos := []int{8, 7, 6}
    fmt.Println("slice before function call", nos)
    subtactOne(nos)                               //function modifies the slice
	fmt.Println("slice after function call", nos) //modifications are visible outside
	

	// 14. Multidimensional slices
	// Similar to arrays, slices can have multiple dimensions.
	pls := [][]string {
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
		}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}


	// 15. Memory Optimisation
	// slices เป็นชิ้นส่วนของ array ตราบใดที่ slices ยังคงใช้งานอยู่ array ไม่สามารถถูกเคลียร์ทิ้งได้ (garbage collected)
	// สมมุติว่ามี array ขนาดใหญ่แต่เราใช้งานจริงๆ แค่บางส่วนใน array นั้น
	// เราสามารถ copy ส่วนนั้นออกมาเป็น array อันใหม่ได้ (pointer ชี้คนละที่)
	// func copy(dst, src []T)
	countriesNeeded := countries()
    fmt.Println(countriesNeeded)
}

func countries() []string {  
    countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
    neededCountries := countries[:len(countries)-2]
    countriesCpy := make([]string, len(neededCountries))
    copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
    return countriesCpy
}

func subtactOne(numbers []int) {  
    for i := range numbers {
        numbers[i] -= 2
    }

}

func changeLocal(num [5]int) {  
    num[0] = 55
    fmt.Println("inside function ", num)

}

func printarray(a [3][2]string) {  
    for _, v1 := range a {
        for _, v2 := range v1 {
            fmt.Printf("%s ", v2)
        }
        fmt.Printf("\n")
    }
}

