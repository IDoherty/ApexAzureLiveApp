package aggFuncs
	
import (
		"log"
        "net"
		"time"
)

// Type Struct Definitions
type Beacon struct {
    Address 	string
    Name	  	string
    Group   	string
}

type LastPacket struct {
    seqNo	 	byte
	gpsTime		uint32
}

// Type Interface Definitions


// Connect to UDP and read in packets. Run a Keep Alive function for each Beacon to maintain connections.

// Read Beacon Addresses from CSV
func GetCSV(fileName string)([]Beacon){
	
	return getCSV(fileName)
}

// UDP connection generator and thread starter
func UdpConnect(beaconAddr  []Beacon, inUDPChan chan<- string){

	udpConnect(beaconAddr, inUDPChan)
}

// UDP Reader on specified IP and Channel
func ReadIn(readConn net.Conn, inUDPChan chan<- string){
 
	readIn(readConn, inUDPChan)
}

// Transmit Keep Alive Packets to all Beacons
func KeepAlive(readConn net.Conn, keepAlive []byte){

	for{
		_, err := readConn.Write(keepAlive)
	
		if err != nil {
				log.Println(err)
		}
		
		time.Sleep(time.Second * 10)
	}
}


// Validate the packets which have been read in over UDP and pass Valid Packets to the Output Channel.

// Packet Processing
func ProcPackets(inUDPChan <-chan string, outUDPChan chan<- string){

	procPackets(inUDPChan, outUDPChan)
}

// Packet Validity Tester
func TestValidity(lastVals *LastPacket, testNo byte, testTime uint32)(bool, int){
	
	return testValidity(lastVals, testNo, testTime)
}

// Distribute Validated Packets to other programs and services
func UdpTransmit(outUDPChan <-chan string){

	udpTransmit(outUDPChan)
}





