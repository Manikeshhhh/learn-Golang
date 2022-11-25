package main
import "os"
import "fmt"
import "io"
import "io/ioutil"
func main(){
	fmt.Println("Helllllllllllllllooooooooooooooooo")
	content := "Heyyyyyy therreeeeeeeeeeeeeeeee, how's it goinnnnn"
	file,err:=os.Create("./myfile.txt")
	if err!=nil{
		panic(err)
	}
	length,err:=io.WriteString(file,content)
	if err!=nil{
		panic(err)
	}
	fmt.Println("The length is: ",length)
	defer file.Close()
	readfile("./myfile.txt")
}
func readfile(filename string){
	databyte,err:=ioutil.ReadFile(filename)
	if err!=nil{
		panic(err)
	}
	fmt.Println("Text data inside the file is: \n",string(databyte))
}