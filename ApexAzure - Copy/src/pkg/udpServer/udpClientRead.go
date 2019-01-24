package udpServer

import (
	"fmt"
	"net"
	"time"
)

func udpClientRead(udpClient UdpClientStruct, server *net.UDPConn, deadChan chan<- int) {

	// Routine Channels
	readChan := make(chan *net.UDPAddr)
	closeChan := make(chan string)
	//errorChan:= make(chan string, 1)

	// Keep Alive Response Packet
	keepAliveResponse := "55dd1e0003010100f6012402bdbd1a23454a0100cd79050004003b21d2d41490efb6dd55"

	go udpReadFunc(server, readChan, closeChan) //,errorChan)

	go udpWriteFunc(udpClient, server)

	for {

		select {
		case results := <-readChan:
			fmt.Println("Keep Alive - ", results)
			udpClient.writeUdpChan <- keepAliveResponse

		case <-time.After(10 * time.Second):
			fmt.Println("Timed Out", udpClient.clientIndex)

			fmt.Println("Client Read Thread Closing")

			deadChan <- udpClient.clientIndex
			fmt.Println("deadChan output good. Routine Closing")

			time.Sleep(time.Millisecond * 10)

			closeChan <- "Closing ReadFunc"
			fmt.Println("closeChan output good")

			server.Close()
			break
			// Potential for Addition of Error Cases if required
		}
	}

}
