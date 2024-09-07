package main

import "fmt"

// //this function can only recieve int type of args-----------------
// 	func vfunc(args...int)int{
// 		red := 0
// 		for _,v := range args{
// 			red +=v
// 		}
// 		return red
// 	}

// func main() {
// 	fmt.Println(1, 2, 3, 4) // 1 2 3 4 
// 	fmt.Println(vfunc(9,9,9,9,9))
// }


//types string and void-----------------------------
	// func printMessages(messages ...string) {
	// 	for _, message := range messages {
	// 		fmt.Println(message)
	// 	}
	// }

	// func main() {
	// 	msgs := []string{"Hello", "Iam", "Shivam", "Petwal"}
	// 	printMessages(msgs...)  //// Unpacks the slice into separate arguments
	// }

//Anonymous function
// Go program to illustrate how 
// to create an anonymous function 

//Error at package level 
	// func(){ 

	// 	fmt.Println("inside the anonymous function") 
	// }() 
	
	func main() { 
		
	// func(){ 

	// 	fmt.Println("inside the anonymous function") 
	// }() 
		
// Assigning to a variable
	add := func(a int, b int) int {
		return a + b
	}
	fmt.Println(add)
	
	} 
			
	// package main
	// import "fmt"
	
	// func add(x int, y int)int{
	// 	return x + y
	// }
	
	// //multilpe return values
	// func oper(x int, y int) (int,int){
	// 	return x+y, x*y    
	// }
	
	// //multi return 2
	// func name() (string, int){
	// 	return "shivam", 22 
	// }
	
	
	// func main() {
	//   fmt.Println(add(10,20)) //30
	//   fmt.Println(oper(4,4)) //8 16
	//   n, a := name()
	//   fmt.Println("name : ",n,",age : ",a)
	  
	// }
