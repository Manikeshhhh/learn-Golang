package main
import "fmt"
func main(){
	fmt.Println("welcome to arrays in go lang")
	var anime [4]string
	anime[0]="AOT"
	anime[1]="Boruto"
	anime[2]="Bleach"
	anime[3]="Tokyo revengers"
	fmt.Println("My Anime list",anime)
	fmt.Println("anime number 2",anime[1])
	fmt.Println("I know it's weird but it takes memory \nno matter you use it or not ,",len(anime))
	var manga = [3]string{"bleach","boruto","chainsaw-man"}
	fmt.Println("Manga's are ", manga)

}