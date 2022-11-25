//In Golang, the defer keyword is used to delay the execution of a function or a statement until the nearby function returns.
// In simple words, defer will move the execution of the statement to the very end inside a function.
package main
import "fmt"
func main(){
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three") ///why?cuzz whyy nottttt
	//jk, it works on LIFO 
	//defer mydefer()
	fmt.Println("broooooooooooooo")
	mydefer()
}

func mydefer(){
	for i:=0;i<6;i++{
		defer fmt.Println(i)
	}
}