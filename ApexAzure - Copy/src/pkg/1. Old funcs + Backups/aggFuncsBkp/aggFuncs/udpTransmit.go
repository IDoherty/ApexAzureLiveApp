package aggFuncs

import (
        "log"
        "net"
		"encoding/hex"
        "fmt"
)

func udpTransmit(outUdpChan <-chan string){
	
	//keepAliveResponse := "55dd1e0003010100f6012402bdbd1a23454a0100cd79050004003b21d2d41490efb6dd55"
				
	//decodedHex, err := hex.DecodeString(keepAliveResponse)
	//if err != nil {
	//	panic(err)
	//}
			
	//fmt.Printf("%s", hex.Dump(decodedHex))
	
	hostName 	:= "192.168.187.146"
    portNum 	:= ":8080"
    service 	:= hostName + portNum

    laddr, err := net.ResolveUDPAddr("udp", service)
    if err != nil {
            log.Fatal(err)
    }

    // setup listener for incoming UDP connection
    ln, err := net.ListenUDP("udp", laddr)
	if err != nil {
            log.Fatal(err)
    }

	//go KeepAliveResponse(decodedHex)
	
	fmt.Println("UDP server up and listening on port 8080")
    defer ln.Close()
	
	buf := make([]byte, 512)

    for {
		outUDP := <-outUdpChan
		
		n, udpAddr, err := ln.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Break err")
            log.Fatal(err)
		}
		
		fmt.Printf("%s", hex.Dump(buf[:n]))
		
		_, err = ln.WriteToUDP([]byte(outUDP), udpAddr)
		if err != nil {
				log.Println(err)
		}
		
		
    }

}