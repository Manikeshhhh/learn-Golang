//mongodb+srv://manikesh:<password>@cluster0.fjsdtsc.mongodb.net/?retryWrites=true&w=majority

package main

import "fmt"
import "net/http"
import "log"
import "go.mongodb.org/mongo-driver/mongo"

func main(){
	fmt.Println("Mongo API")
	r := router.Router()
	fmt.Println("Starting")
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Listening at Port 8000	....")

	
}