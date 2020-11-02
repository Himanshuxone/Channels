package main

import(
	"sync"
	"fmt"
)

var wg sync.WaitGroup

func main(){
	count := 10
	data1 := make(chan int)
	for i:=0;i<=count;i++{
		wg.Add(1)
		go load(i, &wg, data1)
	}

	// go func(){
	// 	wg.Wait()
	// 	close(data1)
	// }()

	for {
		select {
		case data, ok :=  <-data1:
			if ok {
				fmt.Println("more data", "current =>", data)
			}else{
				fmt.Println("channel empty")
				return
			}
		default:
			fmt.Println("Value not sent")
			return
		}
	}

	// select {
	// case data, ok :=  <-data1:
	// 	if ok {
	// 		fmt.Println("more data", "current =>", data)
	// 	}else{
	// 		fmt.Println("channel empty")
	// 		return
	// 	}
	// default:
	// 	fmt.Println("default condition")
	// 	return
	// }


}

func load(count int, wg *sync.WaitGroup, data1 chan int){
	defer wg.Done()
	data1 <- count
}