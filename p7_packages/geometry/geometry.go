package main

// root directory จะอยู่ที่ user/go/
import (
	"fmt"
	"golangbot/p7_packages/geometry/rectangle"
	"log"
)

var rectLen, rectWidth float64 = -1, 7 

func init() {  
    println("main package initialized")
    if rectLen < 0 {
        log.Fatal("length is less than zero")
    }
    if rectWidth < 0 {
        log.Fatal("width is less than zero")
    }
}



func main() {
	// var rectLen, recWidth float64 = 6, 7

	fmt.Println("Geometrical shape properties")


	// 1. Exported Names
	// การตั้งชื่อโดยใช้พิมพ์ใหญ่ go จะถือว่าตัวแปรหรือฟังก์ชั้นนั้นสามารถ export ไปใช้ที่อื่นได้ (หากเป็นตัวเล็กจะไม่ยอม)
    fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
    fmt.Printf("diagonal of the rectangle %.2f ",rectangle.Diagonal(rectLen, rectWidth))
}



	// 2. init function
	// init function ไม่ควรจะมีการ return อะไรกลับไป 
	// init function ไม่ควรมี parameters 
	// init function ไม่สามารถ call ได้จาก package อื่นๆ
	// init function เป็น function สำหรับการทำงานหรือตรวจสอบอะไรบ้างอย่างก่อนที่จะดำเนินการ execution start
	// init ถูก call ก่อน main


	// 3. Blank Identifier
	// หากมีการ import packge อื่นๆ เข้ามาแล้ว ไม่มีการใช้งาน go จะแจ้งเตือนขณะ compile time 
	// แต่ถ้าหากเราต้องการจะ import โดยยังไม่เรียกใช้ สามารถใช้ Blank identifier ช่วยได้
	// var _ = rectangle.Area 

	// หรือถ้าหากเราต้องการจะ import package โดยไม่อ้างอิงถึงใน package นั้น
	// โดยมีข้อแม่้ว่าจะไม่มีการเรียกใช้อะไรเลย ภายแต่ package นั้น 
	// ทำมาเพื่อ import package ที่ยังไม่มีได้มีการใช้งาน