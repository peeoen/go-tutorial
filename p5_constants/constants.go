package main

import (  
    "fmt"
)

func main() {  
    fmt.Println("Hello, playground")
	// constant ใช้ในการระบุตัวแปร ที่ไม่สามารถแก้ไขค่าได้


	// ตัวอย่าง b ถูกระบุด้วย const จะเห็นว่่าเกิด error เพราะ const จะถูกกำหนด่าตั้งแต่ compile time 
	// ไม่สามารถใช้ function ที่จะ return ค่ากลับมาตอน runtime ได้
	// math.Sqrt(4) จะถูก evaluate ขณะ runtime

    // var a = math.Sqrt(4)//allowed
	// const b = math.Sqrt(4)//not allowed
	
	var name = "Sam"
	fmt.Printf("type %T value %v", name, name)
	
	const hello = "Hello World"  
}