package main 

import (
	"fmt"
	"errors"
)
func init(){
	fmt.Println("Error Tutorials")
}

//stdin
//stdout
//stderr

func divide(pay, payda float32)(float32, error){
	//fmt.Errorf("user %q (id %d) not found", name, id)
	if payda == 0{
		return 0.0, errors.New("Payda 0 olamaz..")
	}

	return pay/payda, nil
}

func ageCheck(age int)(bool, error){
	if age < 18{
		return false, errors.New("Age not check")
	}

	return true, nil
}

func main(){

	//var age1 int = 17
	age1 := 17
	_, err := ageCheck(age1)

	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Ok...")
	}

	result, err := divide(1.5, .4)
	
	if err == nil{
		fmt.Printf("Result : %f\n", result)
	}else{
		fmt.Println(err)
	}

	var err_s error = fmt.Errorf("errorf %d\n", age1)
	fmt.Println(err_s)

	//2 + "tekin"  wrong
	//2+3.4        wrong
	//float32 + float64   wrong
	//arc

	//SPrinf
}