package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"math/rand"
	"time"
	"github.com/gorilla/mux"
)

//model for course
type course struct{
	CourseId int `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int	`json:"price"`
	Author *Author	`json:"author"`
}

type Author struct{
	Fullname string `json:"fullname"`
	Website string `json:"website"`

}

//fake db
var courses []course

//middleware or helper

func (c *Course) IsEmpty() bool{
	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}	

func main(){

}


//controller

//serve home route
func servehome(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("<h1>Golang API</h1>"))
}

func GetAllCourses(c){
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	//Grab id from Request
	params := mux.Var(r)

	//loop  through courses and find matching id which user is sending and return the response 
	for _,course:=range courses{
		if course.CourseId==params["id"] {
			json.NewEncoder(w).Encode(course)
			return 
		} else {
			json.NewEncoder(w).Encode("No Course Found matching the id")
			return
		}
	}
}

func createonecourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")	

	//what if body is empty
	if r.Body==nil {
		json.NewEncoder(w).Encode("Please send some data")
	}
	var course Course
	_=json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("Please send some data")
		return
	}
	//generate unique id and convert that into string 
	//append this into course
	rand.Seed(time.Now().UnixNano())
	course.CourseId = rand.Intn(100)
	courses=append(courses,course)
	json.NewEncoder(w).Encode(course)
	return

}

