package udpServer

import (
	//"fmt"
	"net"
)

func udpClientAdd(udpClient []UdpClientStruct, clientAddr *net.UDPAddr) (newUdpClient []UdpClientStruct, tempIndex int) {

	// Declare empty Struct to Append
	var newClient UdpClientStruct

	// Add new Struct to slice
	tempIndex = len(udpClient)
	//fmt.Println("Temp Index - ", tempIndex)
	newUdpClient = append(udpClient, newClient)

	// Fill new Struct's fields
	newUdpClient[tempIndex].clientAddr = clientAddr
	newUdpClient[tempIndex].clientIndex = ClientIndex
	newUdpClient[tempIndex].writeUdpChan = make(chan string, 1)

	ClientIndex++

	//fmt.Println("UDP Client Struct - ", newUdpClient)

	return newUdpClient, tempIndex
}
