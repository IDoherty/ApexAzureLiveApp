package aggFuncs

import (
	"fmt"
	"log"
	"net"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func udpConnect(beaconAddr []Beacon, inUDPChan chan<- string, keepAlive []byte, localAddr string) {

	for i := 0; i < len(beaconAddr); i++ {
		RemoteAddr, err := net.ResolveUDPAddr("udp", beaconAddr[i].Address)
		CheckError(err)

		LocalAddr, err := net.ResolveUDPAddr("udp", localAddr)
		CheckError(err)

		conn, err := net.DialUDP("udp", LocalAddr, RemoteAddr)
		CheckError(err)
		// note : you can use net.ResolveUDPAddr for LocalAddr as well

		fmt.Println()
		log.Printf("Remote UDP address : %s \n", conn.RemoteAddr().String())
		log.Printf("Local UDP client address : %s \n", conn.LocalAddr().String())

		//Start Read and KA threads
		go ReadIn(conn, inUDPChan)
		go KeepAlive(conn, keepAlive)
	}
	fmt.Println()
}
