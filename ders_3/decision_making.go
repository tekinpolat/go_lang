package main

import "fmt"

//https://golangdocs.com/ternary-operator-in-golang

//number1 := 21 invalid
var number1 int = 35

func main(){
	fmt.Println("Decision Making...")

	/* 
		if 
		if-else 
		if-else if .......-else
		if-else if .......
	*/
	var number2 float32 = 35.32

	if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
	}

	{
		var number3 int = 45
	}
	// number3 scope alanının dışında
	//fmt.Printf("Number --> %d\n", number3)
	


	if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
	}
	

	var age int = 19

	if( age < 18){
		fmt.Printf("Yaşınız %v. Reşit değilsiniz\n", age)
	}

	if( age >= 18){
		fmt.Printf("Yaşınız %v. Reşitsiniz\n", age)
	}else{
		fmt.Printf("Yaşınız %d. Reşit değilsiniz\n",age)
	}

	var not int = 96

	if(not > 85){
		fmt.Printf("Notunuz %d. Harf notu A\n", not)
	}else if(not > 70){
		fmt.Printf("Notunuz %d. Harf notu B\n", not)
	}else if(not > 60){
		fmt.Printf("Notunuz %d. Harf notu C\n", not)
	}else{
		fmt.Printf("Notunuz %d. Harf notu F\n", not)
	}

	switch{
		case not > 85:
			fmt.Printf("Notunuz %d. Harf notu A\n", not)
		case not > 70:
			fmt.Printf("Notunuz %d. Harf notu B\n", not)
		case not > 60:
			fmt.Printf("Notunuz %d. Harf notu C\n", not)
		default:
			fmt.Printf("Notunuz %d. Harf notu F\n", not)
	}

	//switch case de break yok
	var letterNot string = "B"
	switch {
		case letterNot == "A":
			fmt.Printf("%s Excellent!\n",letterNot) 
		case letterNot == "B":
			fmt.Printf("%s  Well done\n",letterNot)  
		case letterNot == "C":
			fmt.Printf("%s  You passed\n",letterNot) 
		default:
			fmt.Printf("%s  Invalid grade\n",letterNot);
	}

	var marks int = 90
	var grade string = ""
	switch marks {
		case 90: grade = "A"
		case 80: grade = "B"
		case 50,60,70 : grade = "C"
		default: grade = "D"  
	}

	fmt.Printf("Your grade is  %s\n", grade ); 
}