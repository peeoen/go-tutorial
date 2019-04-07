package main

import (
	"fmt"
)

func main() {
	// if syntax ก็เหมือนภาษาทั่วๆ ไปมี if else แต่ไม่ต้องใส่วงเล็บ ()
	if 1 == 1 {
		fmt.Println("true")
	}

	// if ของ go สามารถใส่ statement ก่อนที่จะเข้าเงื่อนไข (condition )ได้
	// โดย statement จะถูก execute ก่อน condition
	// if statement; conditon { }

	if num := 10; num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}


	// สังเกตุุว่าหาก else เว้นบรรทัดไปอีก 1 บรรทัด จะเกิด error  syntax error: unexpected else, expecting }  
	// เพราะว่าเวลาที่ go compile มันจะเติม semicolon ให้อัตโนมัติหลัง } เลยเกืด error ขึ้น
	// if num := 10; num%2 == 0 {
	// 	fmt.Println(num, "is even")
	// } 
	// else {
	// 	fmt.Println(num, "is odd")
	// }
}
