package main
import (
	"encoding/json"
	"fmt"
)
func main(){
	fmt.Println("Welcome to Json processing")
	jsonencode()
	DecodeJson()
}

type course struct{
	Name string 
	Course string `json:"coursename"`  //this makes alias
	Completed bool `json:"completed"`
	Secret string `json:"-"`
	Tag []string
}

func jsonencode(){

	lcocourses:=[]course{
		{"Manikesh","React JS",false,"Key***********",[]string{"react  js","Facebook"}},
		{"Manikesh","Angular JS",false,"Key***********",[]string{"Angular  js","Google"}},
		{"Manikesh","Golang",true,"Key***********",nil},
	}
	//Package data as json data
	finaljson,err:=json.MarshalIndent(lcocourses,"","\t")
	errorhandling(err)
	fmt.Printf("%s\n",finaljson)

}
	
func errorhandling(x error){
	if x != nil {
		panic(x) 
	}
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev","js"]
	}
	`)

	var lcoCourses course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourses)
		fmt.Printf("%#v\n", lcoCourses)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v  and value is %v and Type is: %T\n", k, v, v)
	}

}