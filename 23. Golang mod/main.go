package main
import (
	"fmt"
 	"net/http"
 	"log"
	"github.com/gorilla/mux"
)
func main(){
	fmt.Println("Golang mod")
	greeter()
	r := mux.NewRouter()
    r.HandleFunc("/", serveHome).Methods("GET")
    http.ListenAndServe(":4000",r)
    log.Fatal(http.ListenAndServe(":4000",r))
}
func greeter(){
	fmt.Println("Greetings")

}
func serveHome(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("<h1>welcome to golang</h1>"))
}	