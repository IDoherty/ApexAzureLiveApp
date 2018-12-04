package aggFuncs

import (
		"log"
		//"fmt"
        "net"
		"encoding/hex"
)

func readIn(readConn net.Conn, inUDPChan chan<- string){
 
	// receive message from server
	for{
        buffer := make([]byte, 1024)
        n, err := readConn.Read(buffer)

		if err != nil {
			log.Println(err)
		}

		// Encode to Hex for Transmission
        encodedStr := hex.EncodeToString(buffer[:n])
		
		inUDPChan <- encodedStr
	}
}