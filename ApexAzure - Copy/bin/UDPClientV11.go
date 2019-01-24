package main

import (
	"encoding/hex"
	"fmt"

	//"log"
	"net"
	"os"
	"time"

	//"encoding/binary"
	"pkg/aggFuncs"
	"pkg/metricFuncs"
)

/* A Simple function to verify error */
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

type UDPServer struct {
	addr   string
	server *net.UDPConn
}

var udp UDPServer

// Start of Main Body of Code
func main() {

	// Build/Read Variables

	keepAlive, err := hex.DecodeString("03010100")

	if err != nil {
		panic(err)
	}

	// Get Beacon Addresses & Details
	fileLocation := "NetworkInfo.txt"
	beaconAddr := aggFuncs.GetCSV(fileLocation)

	// Setup Communication Channels
	// Input Channel - Pass from ReadIn routines to Processing Thread
	inUDPChan := make(chan string, 256)

	// Output Channel - Pass Validated Data to UDP Output Thread
	outUDPChan := make(chan string, 256)

	// Metric Channel - Pass Validated Packet to Processing Thread
	metricChan := make(chan string, 256)

	// Start UDP Connection Thread for each beacon connected to the system. Addresses taken from CSV file.
	// Add Activity check for each? Use Keep Alive response packets?
	aggFuncs.UdpConnect(beaconAddr, inUDPChan, keepAlive)

	// Start Packet Processing Thread
	go aggFuncs.ProcPackets(inUDPChan, outUDPChan, metricChan, keepAlive)

	// Start UDP Distribution Thread
	go aggFuncs.UdpTransmit(outUDPChan, keepAlive)

	// Start Metric Processing Thread
	go metricFuncs.MetricFunc(metricChan)

	// Run Indefinitely until Break. Add Monitoring for Functions and Routines?	Break Function to Kill code?
	for {
		time.Sleep(time.Second * 10)
	}
}
