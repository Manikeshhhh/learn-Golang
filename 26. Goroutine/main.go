package main 
import "fmt"
import "net/http"
import "log"
import "sync"
/*import "time"*/

var wg sync.WaitGroup //pointer

var signals =  []string{"test"}
var mut sync.Mutex //pointer



func main(){
	fmt.Println("Why are you here")
/*	go greeter("Hello")
	greeter("World")*/

	websitelist:=[]string{
		"https://manikeshhhh.github.io",
		"https://lco.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}
	for _,web:=range websitelist{
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)

}	

/*func greeter(s string){s

	for i:=0;i<10;i++{
		time.Sleep(5*time.Millisecond)
		fmt.Println(s)
	}
}
*/

func getStatusCode(endpoint  string){
	defer wg.Done()
	res,err:=http.Get(endpoint)
	if err != nil {
		log.Fatal(err) 
	}else{
	mut.Lock()
	signals = append(signals,endpoint)
	mut.Unlock()
	fmt.Printf("%d Status code for the %s\n",res.StatusCode,endpoint)
}
}