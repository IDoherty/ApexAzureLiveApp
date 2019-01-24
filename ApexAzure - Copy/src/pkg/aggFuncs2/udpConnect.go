package aggFuncs

import (
        "log"
        "net"
        "fmt"
)

func udpConnect(beaconAddr  []Beacon, inUDPChan chan<- string){

	keepAlive := make([]byte, 4)
	
	keepAlive[0] = 0x03
	keepAlive[1] = 0x01
	keepAlive[2] = 0x01
	keepAlive[3] = 0x00
	
	for i:=0;i<len(beaconAddr);i++{
	    RemoteAddr, err := net.ResolveUDPAddr("udp", beaconAddr[i].Address)
		if err != nil {
			log.Fatal(err)
		}
		
		conn, err := net.DialUDP("udp", nil, RemoteAddr)
		if err != nil {
			log.Printf("Could not connect to Beacon at the address : %s \n", conn.RemoteAddr().String())
            log.Fatal(err)
		}else{
		// note : you can use net.ResolveUDPAddr for LocalAddr as well
				
		log.Printf("Remote UDP address : %s \n", conn.RemoteAddr().String())
		log.Printf("Local UDP client address : %s \n", conn.LocalAddr().String())
		}
	//Start Read and KA threads
		go ReadIn(conn, inUDPChan)
		go KeepAlive(conn, keepAlive)
	}
	fmt.Println()
}
