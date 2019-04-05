package main

import "fmt"

func main() {
	
	// 1.Declaring a single variable
	// zero values of int = 0 หากไม่ระบุ value go จะกำหนด default value ให้ในแต่ละ type
	// declare by not initial value
	var age int
	fmt.Println("My age is ", age)
	age = 29
	fmt.Println("my age is ", age)
	age = 54
	fmt.Println("my age is ", age)

	// 2. Declaring a variable with initial value
	// declare a variable with initial value
	var age2 int = 29
	fmt.Println("my age2 is ", age2)


	// 3. Type interface 
	// ถ้ามีการกำหนด value ให้กับตัวแปร go จะนำข้อมูลที่กำหนดไป initial tpye ให้กับตัวแปรตัวนั้น
	// age3 type = int
	var age3 = 29
	fmt.Println("my age3 is ", age3)


	// 4. Multiple variable declaration
	// สามารถระบุตัวแปรหลายตัวได้ใน statment เดียว

	var width, height int = 100, 50
	fmt.Println("width is", width, " height is", height)

	// หรือไม่ต้องระบุ type ก็ได้หาก มีการกำหนด value
	
	var width1, height1 = 100, 50
	fmt.Println("width is", width1, " height is", height1)

	// หรืออาจจะ declare variable หลายตัวโดยระบุหลาย type
	var (
		name4 = "pon"
		age4 = 25
		height4 int
	)
	fmt.Println("my name4 is", name4, ", age4 is", age4, "and height4 is", height4)


	// 5. Short hand declaration
	// go มี format สำหรับการกำหนดตัวแปรแบบ กระชับๆ name := initialValue

	name5, age5 := "pon", 29
	fmt.Println("my name5 is ", name5, " age is ", age5);

	// กรณีที่ declare โดยใช้ := แล้วหากมีการเรียกตัวแปรตัวเดิมแล้วใช้ := ในเซ็ตของการประกาศนั้นจะต้องมีตัวแปรใหม่อย่างน้อย 1 ตัว ที่ประกาศใหม่เพิ่มขึ้นมา
	// ตัวอย่าง b ถูกประกาศไปแล้ว แต่เมื่อประกาศอีก ต้องมีตัวแปรใหม่อยู่ภายในนั้นด้วย ก็คือ c
	
	a, b := 20, 30 // declare variables a and b
    fmt.Println("a is", a, "b is", b)
    b, c := 40, 50 // b is already declared but c is new
    fmt.Println("b is", b, "c is", c)
    b, c = 80, 90 // assign new values to already declared variables b and c
    fmt.Println("changed b is", b, "c is", c)


}