package udpServer

import (
	"fmt"
)

func removeFlag(deadChan <-chan int, removeChan chan<- int, removalPtr *bool) {

	for {
		tempIndex := <-deadChan
		*removalPtr = true
		fmt.Println("Removal of ", tempIndex)
		removeChan <- tempIndex
		*removalPtr = false
	}
}
