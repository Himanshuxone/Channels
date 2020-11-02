package main

import (
	"time"
	"fmt"
)

// Publisher will contain the data to be published
type Publisher struct{
	ch chan string
	closingCh chan interface{}
}

func (p *Publisher) write(data string){
	select {
	case <-p.closingCh:
		return
	default:
	}

	go func(data string) {
		select {
			case <-p.closingCh:
			case p.ch <- data:
		}
	}(data)
	go func(data string){
		
	}(data)
}

func (p *Publisher) Read() <-chan string {
	return p.ch
}

// Close will close the channel
func(p *Publisher) Close(){
	close(p.closingCh)
	go func() {
		for range p.ch {
		}
	}()

	<-time.After(1 * time.Second)
	close(p.ch)
}

func main(){
	// publish to channel
	publish := Publisher{
		ch: make(chan string),	
	}
	for i:=0;i<10;i++{
		publish.write(fmt.Sprintf("Data-%d",i))
	}

	// read from the channel
	for {
		select {
		case dh, ok := <-publish.Read():
			if ok {
				fmt.Println("Value received on channel", dh)
			}else{
				fmt.Println("channel closed")
				return
			}
		}
	}
}