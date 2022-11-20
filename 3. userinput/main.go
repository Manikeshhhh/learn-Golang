package main
import "fmt"
import  "bufio"
import "os"

func main(){
	message := "Hey there, how's it going "
	fmt.Println(message)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please rate our service: ")

	input,_ := reader.ReadString('\n')
	fmt.Println("Thank for rating ", input)
	fmt.Printf("Thank for rating %T", input)

}