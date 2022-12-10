package main
import "fmt"
import "sync"
func main(){
	fmt.Println("Race-conditions")

	wg:=&sync.WaitGroup{}
	mut:=&sync.Mutex{}

	var score =[]int{0}

	wg.Add(3)
	go func(wg *sync.WaitGroup,m *sync.Mutex){
		fmt.Println("1 go")
		mut.Lock()
		score =append(score,1)
		mut.Unlock()	
		wg.Done()
	}(wg,mut)
	//wg.Add(1)
	go func(wg *sync.WaitGroup,m *sync.Mutex){
		fmt.Println("2 go")
		mut.Lock()
		score =append(score,2)
		mut.Unlock()
		wg.Done()
	}(wg,mut)
	//wg.Add(1)
	go func(wg *sync.WaitGroup,m *sync.Mutex){
		fmt.Println("3 go")
		mut.Lock()
		score =append(score,3)
		mut.Unlock()
		wg.Done()
	}(wg,mut)

	wg.Wait()
		fmt.Println(score)
}