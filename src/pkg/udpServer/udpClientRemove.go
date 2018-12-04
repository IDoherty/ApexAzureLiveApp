package udpServer

import (
	"fmt"
)

// Prototype Function for Removing Clients from Address Slice
func udpClientRemove(udpClient []UdpClientStruct, deadIndex int) (newUdpClient []UdpClientStruct) {
	// Locate Desired Element
	var i int

	for i, _ = range udpClient {
		if udpClient[i].clientIndex == deadIndex {
			fmt.Println("Dead Client - ", udpClient[i].clientAddr)
			newUdpClient = append(udpClient[:i], udpClient[i:]...)
			break
		}
	}
	return newUdpClient
}
