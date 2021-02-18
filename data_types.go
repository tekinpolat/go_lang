package main

//import "os"
//import "fmt"
//import "io"

import (
	"os"
	"fmt"
	"io"
)

func main(){
	//data types 
	/* 
		boolean types 
		numeric types 
		string types

		numeric types 
		 -int8  (-128 to 127)     uint8     unsigned(işaret)  0-256       int8  --> 8bit    -
		 -int16    uint16
		 -int32    uint32
		 -int64    uint64

		 float32
		 float64 
		 complex64
		 complex128
	*/
	//a := 21
	var a uint8 = 21
	var number1  int8 = -24
	fmt.Println("Number1 ", number1)
	fmt.Println("Merhaba Go", a)

	number2 := 45.5
	fmt.Println("Number2 ", number2)

	var number3 int = 35 

	fmt.Printf("Number3: %d %o Type:%T\n", number3,number3, number3)

	/* 
		https://golang.org/pkg/fmt/
		%T	a Go-syntax representation of the type of the value
		%d	base 10
		%o	base 8
		%t	the word true or false

		%f	decimal point but no exponent, e.g. 123.456
		%F	synonym for %f

		%f     default width, default precision
		%9f    width 9, default precision
		%.2f   default width, precision 2
		%9.2f  width 9, precision 2
		%9.f   width 9, precision 0
	*/

	var isMaried bool = true 
	fmt.Printf("İs Maried : %t Type:%T\n", isMaried, isMaried)

	var salary float32 = 5980.66
	fmt.Printf("Salary %f %F Type:%T\n", salary, salary, salary)
	fmt.Printf("Salary %9.3f\n", salary)

	var number4 int = 21 
	var number5 int = 43 

	//Sprintf  --> String printf
	result := fmt.Sprintf("Number5 => %[2]d Number4 => %[1]d   result => %[1]d", number4, number5)
	//result := fmt.Sprintf("%[2]d %[1]d\n", 11, 22)


	fmt.Printf("%s\n", result)


	//Stdout 
	//Stdio
	//Stderr   

	//https://www.geeksforgeeks.org/fmt-sprintf-function-in-golang-with-examples/
	const num1, num2, num3 = 5, 10, 15 

  
    // Calling Sprintf() function 
    s := fmt.Sprintf("%d + %d = %d\n", num1, num2, num3) 
  
    // Calling WriteString() function to write the 
    // contents of the string "s" to "os.Stdout" 
	io.WriteString(os.Stderr, s) 
	
	var number6, number7 int = 21, 43 
	fmt.Println(number6, number7)

	//https://medium.com/faun/golangs-fmt-sprintf-and-printf-demystified-4adf6f9722a2
	test := 5
	hello := "Hello"
	isPro := false 
	fmt.Printf("%v %v %v", test, hello, isPro)

	
}