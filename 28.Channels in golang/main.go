package main
import "fmt"
import "sync"
func main(){
	fmt.Println("Channels in golang")
	wg:=&sync.WaitGroup{}
	myCh:=make(chan int,2) //channel
/*	myCh <- 5 //pushing values to channel
	fmt.Println(<-myCh)*/
	wg.Add(2)
	go func(ch chan int,wg *sync.WaitGroup){
		fmt.Println(<-myCh)
		//fmt.Println(<-myCh)
		wg.Done()
}(sync,wg)      //First we are printing because channel works bit differently,  channel will only 
					//let you add value if somebody is listening to it. so we are listenig in first func

	go func(ch chan int,wg *sync.WaitGroup){
		myCh<-5
		myCh<-6
		close(myCh)
		wg.Done()
		}(sync,wg)
		wg.Wait()
}