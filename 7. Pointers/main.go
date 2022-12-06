//pointers gives us surety that whatever we are passing on as a pointer is infact the same variable and not the copy of it
//We pass memory address for pointers
//Pointer is just a direct reference to memory address of the variable 
package main
import "fmt"
func main(){
	fmt.Println("Welcome to pointersssss")
	// var point *int
	// fmt.Println("Default value of pointer is: ", point)
	mynumber := 245
	var point = &mynumber  //& is referencing a pointer
	fmt.Println("address of  pointers is: ", point)
	fmt.Println("value of  pointers is: ", *point)  //This is value inside pointer

	*point = *point * 2  
	fmt.Println("Multiplied value is: ", mynumber)
}