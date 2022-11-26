package main
import "fmt"
import "net/url"
/*import "net/http"
import "io/ioutil"*/
const myurl string = "https://lco.dev:3000/learn?user=manikesh&secret=eydejkmea6SS"
func main(){
	fmt.Println("Handling  urls")
	fmt.Println(myurl)
	//parsing
	result,err:=url.Parse(myurl)		
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme,result.Port(),result.Host,result.RawQuery,result.Path)
	qparam:=result.Query()
	fmt.Println(qparam["user"])

	parts:=&url.URL{
		Scheme:"https",
		Host:"lco.dev",
		Path:"/learn",
		RawQuery:"user=manikesh",
	}
	anotherurl := parts.String()
	fmt.Println(anotherurl)

}	