package main

import "fmt"

/*
	 string --> ""
	 char   --> ''
*/

func init(){
	fmt.Println("String tutorials")
}

func main(){
	var name string = "Enes"

	fmt.Printf("Ad %s\n", name)

	for index := 0; index < len(name); index++{
		fmt.Printf("%c\n", name[index])
	}

	fmt.Println()

	for _, letter := range name{
		fmt.Printf("%c", letter)
	}

	fmt.Println()

	var result string= repeat("a", 10)
	fmt.Println(result)

	result = join([]string{"tekin", "enes", "ayşe","elif"}, ",")
	fmt.Println(result)


	var surname = "KARACAER"
	fmt.Println(surname[0:len(surname)-1])

	result = reverse(surname)
	fmt.Println(result)

	resultcount := count("enes", "e")
	fmt.Println(resultcount)

	resultcount = count("enes", "ne")
	fmt.Println(resultcount)
}

func count(s, substr string) int{
	var letterCount int = 0
	for _, character := range s{
		if(string(character) == substr){
			letterCount +=1
		}
	}

	return letterCount
}


//"enes"
func reverse(str string) string{
	var reverseStr string= ""
	/*
	for index := len(str)-1; index >= 0; index--{
		reverseStr += string(str[index]) 
	}
	*/

	//"" karacaer    arak
	for _, letter := range str{
		reverseStr = string(letter) + reverseStr
	}

	return reverseStr

}



func repeat(str string, count int) string{
	
	result := ""
	for index :=0; index < count; index++{
		result += str
	}
	return result
}

func join(elems []string, sep string) string{
	var result string = ""
	for index, elem :=range elems{
		if index == 0{
			result += elem
		}else{
			result += sep + elem
		}
	}
	return result
}

