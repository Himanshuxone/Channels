package main

import(
	"sync"
	"fmt"
	"time"
)

var wg sync.WaitGroup

func main(){
	// Defining time value 
    // of Since method 
    now := time.Now() 
	for i:=0;i<1000000;i++{
		wg.Add(1)
		goroutine(i, &wg)
	}
	// done := make(chan bool)
	go func(){
		// close(done)
		wg.Wait()
	}()
	// Prints time elapse 
	fmt.Println("time elapse:", time.Since(now)) 
}

func goroutine(count int, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println(fmt.Sprintf("Go routines:- %d",count))
}