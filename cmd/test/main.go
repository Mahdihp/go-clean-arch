package main

import (
	"fmt"
	"time"
)

func main() {
	isErrorChan := make(chan bool, 1)
	stoperChan := make(chan bool, 1)

	go worker(isErrorChan)
	go stoper(stoperChan)
	for {
		select {
		case <-isErrorChan:
			fmt.Println("new worker Running..." + time.Now().Format("05"))
			go worker(isErrorChan)
		case <-stoperChan:
			fmt.Println("Stop worker Running..." + time.Now().String())
			return
			//default:
			//
			//	time.Sleep(1 * time.Second)
		}
	}
}
func stoper(stoperChan chan bool) {
	time.Sleep(10 * time.Second)
	stoperChan <- true
}
func worker(isErrorChan chan bool) {
	//for {
	fmt.Println("worker Running..." + time.Now().String())
	//select {
	//case <-isErrorChan:
	//	return
	//}
	time.Sleep(2 * time.Second)
	isErrorChan <- true
	//}
}
