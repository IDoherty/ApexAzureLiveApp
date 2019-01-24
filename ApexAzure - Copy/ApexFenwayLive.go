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

	//"pkg/udpServer"

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
	sessionWrite, sessionWrite2, keepAlive, devList, localAddr, sessionName
	
	// Build/Read Variables
	// Keep Alive Packet Identifier - Create a String of Hex Bytes for Comparison
	keepAlive, err := hex.DecodeString("03010100")
	if err != nil {
		panic(err)
	}

	// Write Session Options
	sessionWrite := "Sessions/TestSessions/Session - "
	sessionWrite2 := "Sessions/TestSessions/CutSession - "

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

	/*/  UDP Output Channel - Pass Validated Data to UDP Output Thread
	outUDPChan := make(chan string, 64)
	//*/

	//*/  Azure Output Channel - Pass Message Data and Device Identifier to Azure Upload Thread
	outAzureChan := make(chan structs.AzureChanStruct, 64)
	//*/

	//*/  Start Packet Processing Threads
	// outUDPChan provides a direct bypass to the metric functions if this program is to serve as an Aggregator only, or a way to pass unmodified packets out over UDP in parallel to the metric processing.
	go aggFuncs.ProcPackets(inUDPChan, outFileChan, metricChan, keepAlive)
	//go aggFuncs.ProcPackets(inUDPChan, outUDPChan, outFileChan, metricChan, keepAlive) // Requires aggFuncs.UdpTransmit and udpServer.UdpServer
	//*/

	//*/  Start File Writing Threads - Write Aggregated Output to File
	go aggFuncs.WriteToFile(outFileChan, sessionWrite)
	go aggFuncs.WriteToFile(outFileChan2, sessionWrite2)
	//*/

	/*/ Start UDP Distribution Threads
	go aggFuncs.UdpTransmit(outUDPChan)
	go udpServer.UdpServer(outUDPChan)
	//*/

	//*/  Start Metric Processing Threads - Process Aggregated Packets and Output Metrics (Azure and UDP are options)
	go metricFuncs.MetricFunc(metricChan /*, outUDPChan*/, outAzureChan, outFileChan2)
	//*/

	//*/  Start Azure Output Threads - Send Packets up to Azure
	// Device List Options
	//devList := "Config/DevList/testDevList.txt"

	// Clare & Cork 		- Channel 1
	//devList := "Config/DevList/FenwayClareCorkDevList.txt"

	// Cork 				- Channel 1
	devList := "Config/DevList/FenwayCorkDevList.txt"

	// Limerick & Wexford - Channel 9
	//devList := "Config/DevList/FenwayLimWexDevList.txt"
	
	go azureFuncs.AzureUpload(outAzureChan, devList)
	//*/

	//*/  Start UDP Connection Thread for each beacon connected to the system. Addresses taken from CSV file.
	// Get Beacon Addresses & Details from CSV file at designated location
	apexBeaconInfo := "Config/NetworkInfo/NetworkInfoKingspan.txt"
	beaconAddr := aggFuncs.GetCSV(apexBeaconInfo)

	// Address of Device Running the Code
	localAddr := "194.55.105.52:0"

	aggFuncs.UdpConnect(beaconAddr, inUDPChan, keepAlive, localAddr)
	//*/

	/*/	 Start Reading Packets in from an existing session record
	// Read In Session Options
	sessionName := "Sessions/Alumni Stadium - Cork - 17_11_18/Session- 2018-11-17 18-06-47"

	go aggFuncs.ReadFromFile(inUDPChan, sessionName)
	//*/

	// Run Indefinitely until Break. Add Monitoring for Functions and Routines?	Break Function to Kill code?
	for {
		time.Sleep(time.Second * 10)
		//fmt.Println("ping")
	}
}
