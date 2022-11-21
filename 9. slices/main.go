package main
import "fmt"
import "sort"
func main(){
	fmt.Println("welcome to slices in go lang")
	var anime = []string{"bleach","boruto","chainsaw-man"}
	anime = append(anime,"tokyo revengers","Hello world")
	fmt.Println("anime's are ", anime)
	anime = append(anime[1:]) //this will remove element with index 1
	fmt.Println("anime's are ", anime)
	anime = append(anime[1:3]) 
	fmt.Println("anime's are ", anime) //This will print boruto and tkr


	///Trying to use make
	highscore := make([]int,4)  //4 here is default memory allocation, you can definitely use append
	highscore[0] = 245
	highscore[1] = 645
	highscore[2] = 845
	highscore[3] = 945
	//highscore[4]  = 414 (you can't do this because the default is 4 but you can defintely append more values)
	highscore=append(highscore,555,2222,3333,444) //This works, here memory allocation happens again 
	fmt.Println(highscore)

	sort.Ints(highscore)
	fmt.Println(highscore)

	//How to remove value from slice based on index
	var skills = []string{"Python","Golang","Notmyskill","Appsec","Netsec"}
	var index int = 2
	//skills = append(skills[:index])  //This will only give python and golang
	skills = append(skills[:index],skills[index+1:]...)//The three dots (…) in Golang is termed as Ellipsis in Golang which is used in the variadic function. The function that called with the varying number of arguments is known as variadic function. Or in other words, a user is allowed to pass zero or more arguments in the variadic function
	fmt.Println(skills)
}