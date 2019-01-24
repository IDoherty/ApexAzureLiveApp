package aggFuncs

import (
		"fmt"
		"time"
)

func packetCounter(countChan <-chan int){
	
	var newPkt uint32
	var dupPkt uint32
	var oldPkt uint32
	
	newPkt = 0
	oldPkt = 0
	dupPkt = 0
	
	for{
		pktCase := <-countChan
		
		switch pktCase{
		case 1:
			//fmt.Print("New Packet")
			//fmt.Print(arrayLastPackets[slotNo[0]], "\r")
			//fmt.Println() 
			newPkt++
			
		case 2:                    
			//fmt.Print("Duplicate Packet")         
			//fmt.Print(arrayLastPackets[slotNo[0]], "\r")
			//fmt.Println()
			dupPkt++
			
		case 3:                    
			//fmt.Print("Old Packet")               
			//fmt.Print(arrayLastPackets[slotNo[0]], "\r")
			//fmt.Println()
			oldPkt++
		}
		
		fmt.Print("New - ", newPkt, "/Duplicate - ", dupPkt, "/Old - ", oldPkt, "\r")
		
		time.Sleep(time.Nanosecond * 1000)
	}
}