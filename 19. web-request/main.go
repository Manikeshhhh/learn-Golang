package main
import "fmt"
import "net/http"
import "io/ioutil"
const url = "https://lco.dev"
func main(){
	fmt.Println("Web requests")
	response,err:=http.Get(url)
	if err != nil {
		panic(err)
	}	
	fmt.Println(response)
	data,err:=ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	defer response.Body.Close()
}