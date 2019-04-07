package main

import (
	"fmt"
)

func main() {
	fmt.Println("Part 10: Switch Statement")

	finger := 4
    switch finger {
    case 1:
        fmt.Println("Thumb")
    case 2:
        fmt.Println("Index")
    case 3:
        fmt.Println("Middle")
    case 4:
        fmt.Println("Ring")
    case 5:
        fmt.Println("Pinky")
    default: //default case
        fmt.Println("incorrect finger number")
	}
	
	// 1. Multiple expression in case
	letter := "i"
    switch letter {
    case "a", "e", "i", "o", "u": //multiple expressions in case
        fmt.Println("vowel")
    default:
        fmt.Println("not a vowel")
	}
	
	// 2. Expressionless switch
	// ไม่ต้องใส่ expression หากไม่ใส่มันจะเลือกทำงาน case ที่เป็น true อันแรก (เหมือน if- else if)
	num := 75
    switch { // expression is omitted
    case num >= 0 && num <= 50:
        fmt.Println("num is greater than 0 and less than 50")
    case num >= 51 && num <= 100:
		fmt.Println("num is greater than 51 and less than 100")
	case num >= 52 && num <= 100:
        fmt.Println("num is greater than 51 and less than 100 (2)")
    case num >= 101:
        fmt.Println("num is greater than 100")
    }
}
