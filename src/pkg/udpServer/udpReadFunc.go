package udpServer

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
)

func udpReadFunc(server *net.UDPConn, readchan chan<- *net.UDPAddr, closeChan <-chan string) {

	for {

		select {
		case closeRead := <-closeChan:
			fmt.Println(closeRead)
			server.Close()
			break
		default:
			buf := make([]byte, 1024)

			n, addr, err := server.ReadFromUDP(buf)
			if err != nil {
				fmt.Println("Break err")
				log.Fatal(err)
			}

			if n == 4 && binary.LittleEndian.Uint32(buf[0:4]) == KeepAliveTestVal {

				readchan <- addr

			}
			// Add Error Cases here via elseif or Switch w. specific readChan values
		}
	}

}
