package routine

import(
	"sync"
	// "fmt"
)

// Channel for running routines
func Channel(wg *sync.WaitGroup) int{
	count := 10
	data1 := make(chan int)
	for i:=0;i<=count;i++{
		wg.Add(1)
		go load(i, wg, data1)
	}
	
	go func(){
		wg.Wait()
		close(data1)
	}()

	for {
		select {
		case data, ok :=  <-data1:
			if ok {
				_ = data
				// fmt.Println("more data", "current =>", data)
			}else{
				// fmt.Println("channel closed")
				return 0
			}
		}
	}
}

func load(count int, wg *sync.WaitGroup, data1 chan int){
	defer wg.Done()
	data1 <- count
}