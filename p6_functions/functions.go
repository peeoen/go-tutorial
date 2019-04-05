package main

import (  
    "fmt"
)

func main() {  

	// 
	name := "pon";
	fmt.Println(getName(name))


	// 1. กรณี paramter มี type เดียวกัน
	name1 := "Theerawat";
	last1 := "Burawat";
	fmt.Println(getFullName(name1, last1))


	// 2. Multiple return values
	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f", area, perimeter) 
	

	// 3. Named return values
	// ไม่ต้อง return ตัวแปร แต่ go จะสร้างตัวแปรให้แล้ว reurn ตัวแปรตามชื่อที่ได้ตั้งไว้ใน return type
	area3, perimeter3 := rectProps1(10.8, 5.6)
	fmt.Println("") 
	fmt.Printf("Area %f Perimeter %f", area3, perimeter3) 


	// 4. Blank Identifier
	// _ คือ Blank Identifier ใน go คือเป็นการระบุตัวแปรว่างเปล่า หรือไม่ได้ต้องการใช้
	// เช่น rectProps1 จะ return area กับ perimeter กลับมา หากต้องการแค่ area ก็สามารถระบุ perimeter เป็น Blank Identifier ได้

	area4, _ := rectProps(10.8, 5.6)
	fmt.Println("") 
	fmt.Printf("Area %f", area4) 
}

func getName(name string) string {
	return "Mr. " + name;
}


func getFullName(name, last string) string {
	return "Mr. " + name + " " + last;
}

func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}


func rectProps1(length, width float64)(area, perimeter float64) {  
    area = length * width
    perimeter = (length + width) * 2
    return //no explicit return value
}