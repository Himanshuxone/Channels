package main

import(
	"sync"
	"sync/atomic"
	"fmt"
)

// Publisher struct will handle the messages to be published
type Publisher struct {
	ch chan string
	closing uint32
	writersWG sync.WaitGroup
}

func (p *Publisher) Read() <-chan string {
	return p.ch
}

func (p *Publisher) write(data string) {
	p.writersWG.Add(1)
	go func(data string) {
		defer p.writersWG.Done()
		if atomic.LoadUint32(&p.closing) != 0 {
			return
		}

		p.ch <- data
	}(data)
}

// Close will be used to close the channel
func (p *Publisher) Close() {
	atomic.StoreUint32(&p.closing, 1)

	go func() {
		for range p.ch {
		}
	}()

	p.writersWG.Wait()

	close(p.ch)
}

// there are other Publisher methods
func main(){
	publish := Publisher{}
	for i:=0;i<10;i++{
		publish.write(fmt.Sprintf("Data-%d", i))
	}
	publish.Read()
}