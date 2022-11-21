//They are key value pairs

package main

import "fmt"

func main(){
	fmt.Println("maps in golang")
	language := make(map[string]string)
	language["JS"] = "Javascript"
	language["Py"] = "Python"
	language["Go"] = "Golang"
	language["Rb"] = "Ruby"
	fmt.Println("Python shorts for: ",language["Py"])
	delete(language,"Rb")
	fmt.Println(language)



}
