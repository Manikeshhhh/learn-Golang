package main 

import "fmt"
import "time"
	
func  main(){
	fmt.Println("Hello, welcome to time")
	Present := time.Now()
	fmt.Println(Present)
	fmt.Println(Present.Format("01-02-2006 15:04:05 Monday "))
	created := time.Date(2021,time.August,10,23,24,0,0,time.UTC)
	fmt.Println(created)
}