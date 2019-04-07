package main 

import (
	"fmt"
)

func main() {
	fmt.Println("Part 9: Loops")

	//syntax: 
	// for initialisation; condition; post { 
	// }

	// 1. normal loop
	for a := 1; a <= 10; a++ {
        fmt.Printf("%d ", a)
	}

	fmt.Println("")

	// 2. break
	for b := 1; b <= 10; b++ {
        if b > 5 {
            break
        }
        fmt.Printf("%d ", b)
	}
	
	fmt.Println("")

	// 3. continue
	for c := 1; c <= 10; c++ {
        if c%2 == 0 {
            continue
        }
        fmt.Printf("%d ", c)
	}
	
	fmt.Println("")


	// 4. nesated For
	f := 5
    for d := 0; d < f; d++ {
        for e := 0; e <= d; e++ {
            fmt.Print("*")
        }
        fmt.Println()
	}

	fmt.Println("")

	// 5. Labels
	// label ใช้สำหรับการ break outer loop จากภายใน inner loop
	// ตัวอย่าง หาก g กับ h เท่ากับให้ออกจาก loop นอกสุด

	// จากโค้ดด้านล่าง ถ้าใส่ break แบบนี้มันจะเป็น break ของ loop h
	// for g := 0; g < 3; g++ {
    //     for h := 1; h < 4; h++ {
	// 		fmt.Printf("g = %d , h = %d\n", g, h)
	// 		fmt.Printf("i = %d , j = %d\n", i, j)
    //         if g == h {
    //             break
    //         }
    //     }
	// }
	
	// ถ้าต้องการให้ break loop นอกสุด (g)
	// เมื่อ g == j จะออกมา loop นอกทันที
	// outer สามารถตั้งเป็นชื่ออะไรก็ได้
	outer:  
    for g := 0; g < 3; g++ {

        for j := 1; j < 4; j++ {
            fmt.Printf("g = %d , j = %d\n", g, j)
            if g == j {
                break outer
            }
        }

	}


	// 6. More examples
	// 6.1 prints all even numbers from 0 to 10.

	// i := 0
    // for ;i <= 10; { // initialisation and post are omitted
    //     fmt.Printf("%d ", i)
    //     i += 2
	// }

	// หรือ 
    // i := 0
    // for i <= 10 { //semicolons are ommitted and only condition is present หากไม่มี semicolon จะถือว่าอันนั้นเป็น condition ไป
    //     fmt.Printf("%d ", i)
    //     i += 2
    // }
	

}