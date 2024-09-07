package main
import "fmt"
func main() {

	//creating map usign make ***********
		mm1 := make(map[string]int)
		fmt.Println(mm1)  

	//NIL MAP **************************

		var mm2 map[string]string
		fmt.Println(mm2)
		fmt.Println(mm2["ran_key"])   //default values acc. to types 
		//mm2["name"] = "shivam_petwal"  //ERROR , cant write to a nil map

	//EMPTY MAP ************************
		mm3 := make(map[string]string)
		fmt.Println(mm3)
		fmt.Println(mm3["ran_key"]) //default values
		//mm3["name"] = "shivam_petwal"
		mm3["Fname"] = "shivam" // can write to a empty map
		mm3["Lname"] = "petwal"
		fmt.Println(mm3["Fname"])

	//DELETE
		delete(mm3,"Fname")
		
	//clear
		//clear(mm3)


		
	_,ok := mm3["Lname"]

	if ok{
		fmt.Println("EXIST")
	} else{
		fmt.Println("DO NOT EXIST")
	}
}
