package main
import  "fmt"

func main(){
	fmt.Println("looooooooooooppsssssssssssssssssssssss")
	days :=	[]string{"Sunday","Monday","Wednesday","Friday","Sunday"}
	fmt.Println(days)
/*	for i:=0;i<len(days);i++{
		fmt.Println(days[i])
	}*/
/*	for i:=range days {
		fmt.Println(days[i])
	}
	*/
	//you can also return value
	for index,day:=range days{
		fmt.Printf("the index is %v and the value is %v\n",index,day) //you can use comma ok syntax to ignore any value
	}

	///Kinda while loops
	value :=4
	for value<11{
		if value==9{
			value++
			continue
		}
		fmt.Println("the value is: ",value)
		value++
	}

}