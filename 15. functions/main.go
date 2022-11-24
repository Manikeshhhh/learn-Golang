package main
import  "fmt"
func main(){
	fmt.Println("functionssssssssssssss in golang")
	hello()
	result :=sub(6,13)
	fmt.Println("result is: ",result)

	addval :=  proadder(56,44,44,112,223)
	fmt.Println(addval)

	subval,message :=  prosub(5,44)
	fmt.Println(subval)	
	fmt.Println(message)	
}

func sub( a int,b int) int{ 
	//the third int which is outside is a signature, which tells us which type of value we are expecting from user
	return b-a
}


func hello(){
	fmt.Println("hellooooooooooooooooooo")
}

//When you have no idea how many values to expect
func proadder(values ...int)int{
	total := 0
	for _,val := range values {
		total  += val
	}	
	return total

}

func prosub(values2 ...int)(int,string){
	total2:= 0
	for _,val2:=range values2{
		total2 += val2
	}
	return total2,"hellooo"
}