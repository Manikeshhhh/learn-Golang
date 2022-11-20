package main
import "fmt"

//first letter capital means its public  or i would have done this "logintokenn"
const Logintokenn string = "eyrxxxxxxxxxxxxwew" 	

func main(){

	fmt.Println(Logintokenn)

	var  username string = "Manikesh"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n ", username)

	var  isloggedin bool = true
	fmt.Println(isloggedin)
	fmt.Printf("variable is of type: %T \n ", isloggedin)

	var smallvalue int = 256 
	fmt.Println(smallvalue)
	fmt.Printf("variable is of type: %T \n ", smallvalue)

	var smallfloat float64 = 255.24422222222222222 
	fmt.Println(smallfloat)
	fmt.Printf("variable is of type: %T \n ", smallfloat)

	var defvalue string
	fmt.Println(defvalue)
	fmt.Printf("variable is of type: %T \n ", defvalue)	

	var website = "manikesh.in"
	fmt.Println(website)
	fmt.Printf("This is of type: %T \n", website)

	numberofuser := 2456644
	fmt.Println(numberofuser)
}
