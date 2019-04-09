package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Part 14: Strings")

	// 1. What is a Strings?
	// A string in go is a slice of bytes.
	// Strings in Go are Unicode compliant and are UTF-8 Encoded.

    name := "Hello World"
    printBytes(name)
    fmt.Printf("\n")
	printChars(name)


	// 2. กรณีที่ access แต่ละ charecter แล้วคำบางคำ return มามากกว่า 1 byte ทำให้การแสดงผลผิดพลาด
	fmt.Printf("\n")
    name = "Señor"
	printBytes(name)
	// output: S e Ã ± o r
	// เหตุผลที่แสดงผลผิดเพราะว่า go จะ return กลับมาเป็น bytes  ซึ่ง ñ return กลับมาเป็น 2 byte  c3 and b1
	// เลยทำให้ตอนที่วนลูป มันมองว่า c3 และ b1 เป็นคนละตัวกัน c3 = Ã, b1 = ± ซึ่ง byte ที่ต้องการตือ 2 ตำแหน่ง (c3 b1 = ñ)
	// จะต้องใช้ rune ในการแก้ปัญหา
    fmt.Printf("\n")
	printChars(name)
	

	// 3. rune 
	// rune is an alias of int32
	// rune เป็น type ของ go ซึ่งจะช่วยให้แสดงผลลัพธ์ได้อย่างถูกต้องไม่ว่า charecter นั้นจะ retrun มามากกว่า 1 byte
	fmt.Printf("\n")
	printCharsWithRune(name)


	// 4. For range loop on a string
	// for loop จะ return ค่ากลับมา 2 คือ index number, rune rune
	// sample:     for index, rune := range s
	fmt.Printf("\n")
    name4 := "Señor"
	printCharsAndBytes(name4)
	

	// 5. constructing string from slice of bytes
	// convert from hexdecimal
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
    str := string(byteSlice)
	fmt.Println(str) // output: Café
	

	// convert from decimal
	byteSlice1 := []byte{67, 97, 102, 195, 169}//decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
    str1 := string(byteSlice1)
	fmt.Println(str1)
	
	// 6. Constructing a string from slice of runes
	// convert from hexdecimal
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
    str6 := string(runeSlice)
	fmt.Println(str6)
	
	// convert from decimal
	runeSlice2 := []rune{83, 101, 241, 111, 114}//decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
    str62 := string(runeSlice2)
	fmt.Println(str62)


	// 9. Length of string
	word1 := "Señor" 
    length(word1)
    word2 := "Pets"
	length(word2)
	

	// 10. String are immutable 
	// ถ้าจะเปลี่ยนแปลงค่าของ string บางตัวจะต้อง convert เป็น rune array แล้วค่อยเปลี่ยนแปลงค่า
    h := "hello"
    fmt.Println(mutate([]rune(h)))

}

func mutate(s []rune) string {  
    s[0] = 'a' 
    return string(s)
}

func printCharsAndBytes(s string) {  
    for index, rune := range s {
        fmt.Printf("%c starts at byte %d\n", rune, index)
    }
}

func printBytes(s string) {  
	// len will return number of bytes in the string
    for i:= 0; i < len(s); i++ {
		// %x format specifier for hexadecimal
        fmt.Printf("%x ", s[i])
    }
}



func printChars(s string) {  
    for i:= 0; i < len(s); i++ {
        fmt.Printf("%c ",s[i])
    }
}



func printCharsWithRune(s string) {  
    runes := []rune(s)
    for i:= 0; i < len(runes); i++ {
        fmt.Printf("%c ",runes[i])
    }
}

func length(s string) {  
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
	// fmt.Printf("length of %s is %d\n", s, len(s))
}