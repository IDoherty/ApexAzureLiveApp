package aggFuncs

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"net"
)

type UDPServer struct {
	addr   string
	server *net.UDPConn
}

var udp UDPServer

func udpTransmit(outUdpChan <-chan string) {

	var keepAliveTestVal uint32 = 65795

	keepAliveResponse := "55dd1e0003010100f6012402bdbd1a23454a0100cd79050004003b21d2d41490efb6dd55"

	decodedHex, err := hex.DecodeString(keepAliveResponse)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s", hex.Dump(decodedHex))

	hostName := "192.168.187.131:8808"
	//portNum 	:= ":8808"
	udp.addr = hostName //+  portNum

	laddr, err := net.ResolveUDPAddr("udp", udp.addr)
	if err != nil {
		log.Fatal(err)
	}

	// setup listener for incoming UDP connection
	udp.server, err = net.ListenUDP("udp", laddr)
	if err != nil {
		log.Fatal(err)
	}

	//go KeepAliveResponse(decodedHex)

	defer udp.server.Close()

	fmt.Println("UDP server up and listening on: ", laddr)

	buf := make([]byte, 1024)
	cnt := 0

	for {

		n, udpAddr, err := udp.server.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Break err")
			log.Fatal(err)
		}

		if udpAddr == nil {
			continue
		}

		fmt.Println("udpAddr: ", udpAddr)

		/*/ UDP Address Printout
		//fmt.Println(udpAddr)
		//fmt.Println(buf)
		//*/

		if n == 4 && binary.LittleEndian.Uint32(buf[0:4]) == keepAliveTestVal {

			cnt++
			fmt.Println("Keep Alive Packet: Listener")
			fmt.Println(hex.Dump(buf[:n]), " Packet Number - ", cnt)
			fmt.Println()

			udp.server.WriteToUDP(decodedHex, udpAddr)

			//*/
			for i := 0; i < 100; i++ {

				outUDP := <-outUdpChan
				outputData, err := hex.DecodeString(outUDP)
				if err != nil {
					fmt.Println("Break err")
					log.Fatal(err)
				}

				udp.server.WriteToUDP(outputData, udpAddr)
				//*/

				/*/ Output Packet Hex Dump
				//fmt.Println("Output Packet")
				//fmt.Printf("%s", hex.Dump(outputData))
				//fmt.Println()
				//*/
			}
		} else {
			//*/
			outUDP := <-outUdpChan
			outputData, err := hex.DecodeString(outUDP)
			if err != nil {
				fmt.Println("Break err")
				log.Fatal(err)
			}

			udp.server.WriteToUDP(outputData, udpAddr)
		}
		//*/
	}

}
