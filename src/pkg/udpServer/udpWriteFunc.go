package udpServer

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"
)

func udpWriteFunc(udpClient UdpClientStruct, server *net.UDPConn) {

	for {
		outUDP := <-udpClient.writeUdpChan
		outputData, err := hex.DecodeString(outUDP)
		if err != nil {
			fmt.Println("Break err")
			log.Fatal(err)
		}

		//fmt.Println(hex.Dump(outputData))

		server.WriteToUDP(outputData, udpClient.clientAddr)

		//fmt.Println("Output Packet")
		//fmt.Printf("%s", hex.Dump(outputData))
		//fmt.Println()
		//*/
	}

}
