package aggFuncs

import (
        "log"
        "net"
		"encoding/hex"
        "fmt"
		//"pkg/udpFuncs"
)

type UDPServer struct {
	addr   string
	server *net.UDPConn
}

var udp UDPServer

func udpTransmit(outUdpChan <-chan string){
	
	hostName 	:= "192.168.187.146"
    portNum 	:= ":8080"
    service 	:= hostName + portNum
	
	// Setup local address for use in forming connections
    laddr, err := net.ResolveUDPAddr("udp", service)
    if err != nil {
            log.Fatal(err)
    }
	 
    // Setup listener for incoming UDP connection
    udp.server, err = net.ListenUDP("udp", laddr)
	if err != nil {
            log.Fatal(err)
    }
	
	fmt.Println("UDP server up and listening on", laddr)
	//KeepAliveResponse(udp.server)
	
	buf := make([]byte, 512)
	
    for {
		
		outUDP := <-outUdpChan
		
		//fmt.Printf("%s", hex.Dump([]byte(outUDP)))
		//fmt.Println()
		fmt.Println("Break 1")
				
		n, conn, err := udp.server.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Break err")
            log.Fatal(err)
		}
		
		fmt.Printf("%s", hex.Dump(buf[:n]))
		fmt.Println("Break 2")
		//fmt.Println(string(buf[:n]))
				
		udp.server.WriteToUDP([]byte(outUDP), conn)
		
		fmt.Println("Break 3")
    }
	defer udp.server.Close()
}
