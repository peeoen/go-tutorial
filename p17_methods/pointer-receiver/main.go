package main

import (  
    "fmt"
)

type rectangle struct {  
    length int
    width  int
}

func perimeter(r *rectangle) {  
    fmt.Println("perimeter function output:", 2*(r.length+r.width))

}

func (r *rectangle) perimeter() {  
    fmt.Println("perimeter method output:", 2*(r.length+r.width))
}

func main() {  
    r := rectangle{
        length: 10,
        width:  5,
    }
	p := &r //pointer to r
	
    perimeter(p)
    p.perimeter()

    /*
        cannot use r (type rectangle) as type *rectangle in argument to perimeter
    */
    //perimeter(r)

	perimeter(&r);
    r.perimeter()//calling pointer receiver with a value

}