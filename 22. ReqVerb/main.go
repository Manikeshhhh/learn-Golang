package main
import "fmt"
import "net/url"
import "net/http"
import "io/ioutil"
import "strings"
func main(){
	fmt.Println("making requests")
	//makegetreq()
	//makepostreq()
	formreq()
}	
func makegetreq(){
	const myurl="http://localhost:8000/get"
	response,err:=http.Get(myurl)
	errorhandling(err)
	defer response.Body.Close()
	fmt.Println("response status code is: ",response.StatusCode)
	fmt.Println("response content lenght is: ",response.ContentLength)
	content,err:=ioutil.ReadAll(response.Body)
	errorhandling(err)
	fmt.Println(string(content))
}
func makepostreq(){
	const myurl = "http://localhost:8000/post"
	//fake-json
	requestbody:=strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price": 0,
			"platform":"learnCodeOnline.in"
		}
	`)
	response,err:=http.Post(myurl,"application/json",requestbody)
	errorhandling(err)
	content,err:=ioutil.ReadAll(response.Body)
	errorhandling(err)
	fmt.Println(string(content))
	defer response.Body.Close()

}
func formreq(){
	const myurl = "http://localhost:8000/postform"
	//fake form
	data:=url.Values{}
	data.Add("Name","manikesh")
	data.Add("email","manikesh.dev")
	data.Add("Language","Golang")
	response,err:=http.PostForm(myurl,data)
	errorhandling(err)
	content,err:=ioutil.ReadAll(response.Body)
	errorhandling(err)
	fmt.Println(string(content))
	defer response.Body.Close()	

}

func errorhandling(x error){
	if x != nil {
		panic(x) 
	}
}
