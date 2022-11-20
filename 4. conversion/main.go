package main
import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

func main(){
	fmt.Println("Hello, Thanks for using our service")
	fmt.Println("Please rate our service: ")
	reader := bufio.NewReader(os.Stdin)
	input,_:= reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)

	fmt.Println("we will be adding 1 to your input now ")
	numrating,err:= strconv.ParseFloat(strings.TrimSpace(input),64)
	if err!=nil {
		fmt.Println(err)
	}  else {
			fmt.Println("added 1 to rating: ", numrating+1)
	}

}
