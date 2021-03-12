package main 

import ("fmt")
type user struct{
	id int
	name, email string
	userAddress address
	userSchool school
}

type address struct{
	street int
	city, country string
}

type school struct{
	universty string
	high_school string
}


func init(){
	fmt.Println("Struct pass parameter to function")
}

func setName(user1 user){
	fmt.Printf("%p\n", &user1)
	user1.name = "Enes"
}

func printUser(user1 user){
	fmt.Printf("Id: %d Name:%s Email:%s\n", user1.id, user1.name, user1.email)
	fmt.Printf("Address %d %s %s\n",user1.userAddress.street, user1.userAddress.city, user1.userAddress.country)
	fmt.Printf("School Uni:%s High Scool:%s\n", user1.userSchool.universty, user1.userSchool.high_school)
}

func testpass(numberPtr *int){
	*numberPtr = 6

	//*numberPtr  gösterdiği değişkenin değeri
	//&numberPtr  değişkenin adresi
	//numberPtr	  gösterdiğin değişkenin adresi		
}

func main(){
	var number int = 5  //0x2540A0
	fmt.Printf("%d\n", number)
	testpass(&number)   //0x2540A0

	user1 := user{
		id:1, 
		name:"Tekin", 
		email:"tekin.polat.dpu@gmail.com",
		userAddress: address{street:288, city:"İzmir", country:"Türkiye"},
		userSchool : school{universty:"Dpu", high_school:"500 Evler Lisesi"},
	}

	fmt.Printf("%p\n", &user1)

	printUser(user1)
	setName(user1)
	printUser(user1)
}