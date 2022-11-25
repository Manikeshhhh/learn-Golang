package main
import "fmt"
func main(){
	fmt.Println("Fuck structs, we having a bad day today")
	//No inheritence in golang
	user := User{"Manikesh","manikesxx@wearehackerone.com",19,true}
	fmt.Println(user)
	fmt.Printf("My details are: %+v\n", user)
	fmt.Printf("My name is: %v ,this is my age: %v and here is my email address for contacting: %v\n", user.Name,user.Age,user.Email)
	user.getstatus()
	user.Newmail()
	user.Newname()
	fmt.Printf("My name is: %v ,this is my age: %v and here is my email address for contacting: %v\n", user.Name,user.Age,user.Email)
}

type User struct{
	//we are use U as capital beacause to show that it is public and can be exported
	Name string	
	Email string 
	Age int 
	LoggedIn bool
}	
	
func (u User) getstatus(){
	fmt.Println("Is user active: ",	u.LoggedIn)
}

//manipulate property
func (u User) Newmail(){
	u.Email = "manikesh@go.dev"
	fmt.Printf("New email is: %v\n", u.Email)
} //This is passing a copy, but you can use pointers to manipulate the actual string 

func (u User) Newname(){
	u.Name = "Manikeshhhhhhhhhhhhhhhhh"
	fmt.Printf("New name is: %v\n", u.Name)
}