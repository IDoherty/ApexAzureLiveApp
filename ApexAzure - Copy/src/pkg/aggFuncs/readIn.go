package aggFuncs

import (
	"log"
	//"fmt"
	"encoding/hex"
	"net"
)

func readIn(readConn net.Conn, inUDPChan chan<- string) {

	buffer := make([]byte, 1024)
	// receive message from server
	for {

		n, err := readConn.Read(buffer)

		if err != nil {
			log.Println(err)
		}

		// Encode to Hex for Transmission
		encodedStr := hex.EncodeToString(buffer[:n])

		inUDPChan <- encodedStr
	}
}
