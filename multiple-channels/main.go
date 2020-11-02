package main

import (
	"fmt"
)

func main(){
	done, err := make(chan bool), make(chan error)
	ticker := time.NewTicker(refreshInterval)
	for i:=0;i<10;i++{
		err <-Debug()
	}
	select {
	case <-done:
	case <-err:
	}
}

// Debug will return error on channel
func Debug() <-chan error{
	var err chan error
	err <- fmt.Errorf("Create an error")
	return err
}

func worker(assignment int, done chan bool){
	for i:=0;i<10;i++{

	}
}