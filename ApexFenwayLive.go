package main

import (
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"time"

	"pkg/aggFuncs"
	"pkg/azureFuncs"
	"pkg/metricFuncs"

	//	"pkg/udpServer"

	structs "pkg/structPrototypes"
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

	// Read in Config Variables
	//sessionWrite, sessionWrite2, keepAlive, devList, localAddr, sessionName
	config := aggFuncs.GetConfigCSV()

	// Build/Read Variables
	// Keep Alive Packet Identifier - Create a String of Hex Bytes for Comparison
	keepAlive, err := hex.DecodeString("03010100")
	if err != nil {
		panic(err)
	}

	// Setup Communication Channels for passing between sections of code
	//*/  Input Channel - Pass from ReadIn routines to Processing Thread
	inUDPChan := make(chan string, 64)
	//*/

	//*/  Output Packet Channel - Pass Validated Packets to File Writing Thread
	outFileChan := make(chan string, 64)
	outFileChan2 := make(chan string, 64)
	//*/

	//*/  Metric Channel - Pass Validated Packet to Processing Thread
	metricChan := make(chan string, 64)
	//*/

	/*/  UDP Output Channel - Pass Validated Data to UDP Output Thread - Disabled at this Time.
		outUDPChan := make(chan string, 64)
	//*/

	//*/  Azure Output Channel - Pass Message Data and Device Identifier to Azure Upload Thread
	outAzureChan := make(chan structs.AzureChanStruct, 64)
	//*/

	//*/  Start Packet Processing Threads
	// outUDPChan provides a direct bypass to the metric functions if this program is to serve as an Aggregator only, or a way to pass unmodified packets out over UDP in parallel to the metric processing.
	go aggFuncs.ProcPackets(inUDPChan /*, outUDPChan*/, outFileChan, metricChan, keepAlive) // Requires aggFuncs.UdpTransmit and udpServer.UdpServer for outUDPChan
	//*/

	//*/  Start File Writing Threads - Write Aggregated Output to File
	if config.WriteOn == true {
		go aggFuncs.WriteToFile(outFileChan, config.SessionWrite)
	}
	if config.Write2On == true {
		go aggFuncs.WriteToFile(outFileChan2, config.SessionWrite2)
	}
	//*/

	/*/ Start UDP Distribution Threads
	if config.UDPOutOn == true {
		go aggFuncs.UdpTransmit(outUDPChan)
		go udpServer.UdpServer(outUDPChan)
	}
	//*/

	//*/  Start Metric Processing Threads - Process Aggregated Packets and Output Metrics (Azure and UDP are options)
	go metricFuncs.MetricFunc(metricChan /*, outUDPChan*/, outAzureChan, outFileChan2, config.Write2On, config.DevList)
	//*

	//*/  Start Azure Output Threads - Send Packets up to Azure
	go azureFuncs.AzureUpload(outAzureChan, config.DevList)
	//*/

	//*/  Start UDP Connection Thread for each beacon connected to the system. Addresses taken from CSV file.
	if config.UDPInOn == true {
		beaconAddr := aggFuncs.GetCSV(config.BeaconAddr)

		aggFuncs.UdpConnect(beaconAddr, inUDPChan, keepAlive, config.LocalAddr)
	}
	//*/

	//*/	 Start Reading Packets in from an existing session record

	if config.ReadInOn == true {
		go aggFuncs.ReadFromFile(inUDPChan, config.SessionName)
	}
	//*/

	// Run Indefinitely until Break. Add Monitoring for Functions and Routines?	Break Function to Kill code?
	for {
		time.Sleep(time.Second * 10)
		//fmt.Println("ping")
	}
}
