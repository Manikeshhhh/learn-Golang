package main
import "fmt"
func main(){
	fmt.Println("ifelse in golang")
	touched_grass := false
	var result string
	if touched_grass == true{
		result = "Awesome"
	}else if touched_grass==false {
		result = "Brahhhhhh, go get some beachesss"
	}else{
		result = "Dude,you alright?"
	}
	fmt.Println(result)
}