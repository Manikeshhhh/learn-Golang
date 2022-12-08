package main 
import "fmt"
import "time"
func main(){
	fmt.Println("Why are you here")
	go greeter("Hello")
	greeter("World")

}

func greeter(s string){

	for i:=0;i<10;i++{
		time.Sleep(5*time.Millisecond)
		fmt.Println(s)
	}
}
