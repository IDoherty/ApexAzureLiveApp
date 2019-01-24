package aggFuncs

import (
		"encoding/hex"
        "fmt"
)

func udpTransmit(outUdpChan <-chan string){
	
    for{
		
		outUDP := <-outUdpChan
		
		fmt.Printf("%s", hex.Dump([]byte(outUDP)))
		fmt.Println()
    }
}
