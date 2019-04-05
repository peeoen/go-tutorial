package main

import (
	"fmt"
	"unsafe"
)

func main() {


// bool
// Numeric Types
// 	- int8, int16, int32, int64, int
// 	- uint8, uint16, uint32, uint64, uint
// 	- float32, float64
// 	- complex64, complex128
// 	- byte
// 	- rune
// string


// int8 = -128 to 127
// int16 = -32768 to 32767
// int32 = -2147483648 to 2147483647
// int64 = -9223372036854775808 to 9223372036854775807
// int = represents 32 or 64 bit integers depending on the underlying platform. 
//		 You should generally be using int to represent integers unless there is a need to use a specific sized integer. 
// 		 size: 32 bits in 32 bit systems and 64 bit in 64 bit systems. 
// 		 range: -2147483648 to 2147483647 in 32 bit systems and -9223372036854775808 to 9223372036854775807 in 64 bit systems


	var a int = 89
	b := 95
	fmt.Println("value of a is", a, "and b is", b)


	var a1 int = 89
    b1 := 95
    fmt.Println("value of a is", a1, "and b is", b1)
    fmt.Printf("type of a is %T, size of a is %d", a1, unsafe.Sizeof(a1)) //type and size of a
	fmt.Printf("\ntype of b is %T, size of b is %d", b1, unsafe.Sizeof(b1)) //type and size of b
	


// Unsigned integers
// uint8: represents 8 bit unsigned integers 
// size: 8 bits 
// range: 0 to 255

// uint16: represents 16 bit unsigned integers 
// size: 16 bits 
// range: 0 to 65535

// uint32: represents 32 bit unsigned integers 
// size: 32 bits 
// range: 0 to 4294967295

// uint64: represents 64 bit unsigned integers 
// size: 64 bits 
// range: 0 to 18446744073709551615

// uint : represents 32 or 64 bit unsigned integers depending on the underlying platform. 
// size : 32 bits in 32 bit systems and 64 bits in 64 bit systems. 
// range : 0 to 4294967295 in 32 bit systems and 0 to 18446744073709551615 in 64 bit systems


// Floating point types
// float32: 32 bit floating point numbers 
// float64: 64 bit floating point numbers

	a3, b3 := 5.67, 8.97
    fmt.Printf("type of a3 %T b3 %T\n", a3, b3)
    sum := a3 + b3
    diff := a3 - b3
    fmt.Println("sum", sum, "diff", diff)

    no1, no2 := 56, 89
    fmt.Println("sum", no1+no2, "diff", no1-no2)

// Other numeric types
// byte is an alias of uint8 
// rune is an alias of int32



// Type Conversion
// go จะค่อนข้าง strict กับการระบุ type มันจะไม่แปลง type ให้ auto
// i := 55      //int
// j := 67.8    //float64
// sum := i + j //int + float64 not allowed
// fmt.Println(sum)

i := 10
var j float64 = float64(i) //this statement will not work without explicit conversion
fmt.Println("j", j)
}