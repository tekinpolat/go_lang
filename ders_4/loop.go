package main

import "fmt"

var a int = 16;

func init(){
	fmt.Println("Loop starting ...")
}


func main() {
	

	//int 0
	//double 0.0
	//string ""
	//boolean false

	

	var ids [7]int
	ids[0] = 1
	ids[1] = 2
	ids[6] = 4 

	for index, value := range ids{
		fmt.Println(index, value)
	}

	for index := 0; index < 10; index++{
		fmt.Printf("%d ", index*2)
	}
	fmt.Println()

	var count int = 1
	for{
		fmt.Printf("%v ", count)
		count++

		if count > 10{
			break
		} 
		
	}

	//var numbers int[5] = {1,2,3,4,5} wrong 

	var nums = [4]int{1, 2, 3, 4} 

	//nums[0]

	for index := 0; index < len(nums); index++{
		fmt.Printf("%d", nums[index])
	}
	fmt.Println()

	for index := range nums{
		fmt.Println(index, nums[index])
	}
	//fmt.Println()

	for index, value := range nums{
		fmt.Println(index, value) 
	}

	//fmt.Println(index, value)  not access

	// for 97  while 3

	i := 0
    for i<5 {
        fmt.Println(i)       // prints 0 2 4
        i += 2              
	}
	
	for _, value := range nums{
		fmt.Println(value) 
	}

	//var items []int = []int{1, 2, 3, 4, 5}
	//var number int = int 3  wrong


	var names = [4]string{"tekin", "enes", "ayşe", "ali"}

	for _, name := range names{
		fmt.Println(name)
	}

	fmt.Printf("%v\n", names)
	fmt.Printf("%#v\n", names)
	fmt.Printf("%+v\n", names)

}
