package main

import(
	"sync"
	"fmt"
	"github.com/Himanshuxone/Channels/routine"
)

var wg sync.WaitGroup

func main(){
	routine.Channel(&wg)
}