package udpServer

import (
	//"encoding/binary"
	//"encoding/hex"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

type UdpClientStruct struct {
	clientAddr  *net.UDPAddr
	clientIndex int
	// Channel Elements
	writeUdpChan chan string // Clone of outChan containing outgoing packets
}

type UDPServer struct {
	addr    string
	server  *net.UDPConn
	udpAddr *net.UDPAddr
}

type PortStruct struct {
	portIndex string
	portBound bool
}

// Keep Alive Test Value
var KeepAliveTestVal uint32 = 65795

// Declare Indexing Variable and Slice
var ClientIndex int = 0
var indexSlice []int

//var indexPtr = &indexSlice
var tempIndex int

var MAX_CLIENTS int = 4
var numClients int = 0
var clientPtr = &numClients

// Declare Server Struct
var udp UDPServer

// Declare Client Address Slice
var UdpClient []UdpClientStruct
var udpClientptr = &UdpClient

// Declare Flags and Channels for Removal of GoRoutines
var RemovalRequired bool = false
var removalPtr = &RemovalRequired

var running bool = false

// Declare Mutex for Add/Remove Operations
var AddRemoveMutex = &sync.Mutex{}

// Main Server Function
func UdpServer(outUdpChan <-chan string) {

	// Declare Server Address
	hostAddr := "192.168.187.131:"

	portNum := make([]PortStruct, 4)

	portNum[0].portIndex = "8808"
	portNum[1].portIndex = "8809"
	portNum[2].portIndex = "8810"
	portNum[3].portIndex = "8811"

	// Misc. Variables and Channels
	buf := make([]byte, 1024)
	//cnt := 0

	//testVal := 0

	// Define Channels
	DeadChan := make(chan int, 5)
	RemoveChan := make(chan int, 4)

	go removeFlag(DeadChan, RemoveChan, removalPtr)

	for {

		// Add Client and Start Writing Data to it
		if numClients < MAX_CLIENTS {

			// Test Available Ports for Usage
			for j := 0; j < MAX_CLIENTS; j++ {

				if portNum[j].portBound == false {
					portNum[j].portBound = true
					udp.addr = hostAddr + portNum[j].portIndex
					break
				}
			}
			fmt.Println("portNum: ", portNum)

			// Resolve Local Address
			laddr, err := net.ResolveUDPAddr("udp", udp.addr)
			if err != nil {
				log.Fatal(err)
			}
			// Split here and Form Listener Threads to next Break
			// setup listener for incoming UDP connection
			udp.server, err = net.ListenUDP("udp", laddr)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("UDP server up and listening on: ", laddr)

			// Read Initial Packet and Address of Sender from UDP Connection
			_, udpAddr, err := udp.server.ReadFromUDP(buf)
			if err != nil {
				fmt.Println("Break err")
				log.Fatal(err)
			}
			// If no Address Recieved, Return to Start of Loop
			if udpAddr == nil {
				continue
			}
			fmt.Println("udpAddr: ", udpAddr)

			AddRemoveMutex.Lock()
			UdpClient, tempIndex = udpClientAdd(UdpClient, udpAddr)
			indexSlice = append(indexSlice, tempIndex)
			AddRemoveMutex.Unlock()

			numClients++

			//fmt.Println("index: ", tempIndex)
			//fmt.Println("indexSlice: ", indexSlice)
			//fmt.Println("UdpClient Values: ", len(UdpClient))

			AddRemoveMutex.Lock()
			go udpClientRead(UdpClient[tempIndex], udp.server, DeadChan)
			AddRemoveMutex.Unlock()

			fmt.Println("New Client - ", UdpClient[tempIndex].clientAddr)
		}

		// Add reciever function for deadChan and setup the Client Removal Function
		// Add Clone function for outChan. One channel for each Indexed Client
		if running == false {
			go func(RemoveChan <-chan int, clients *int, indexSlice []int) {
				fmt.Println("Closer Started")
				for {
					select {
					case deadIndex := <-RemoveChan:
						fmt.Println(deadIndex)

						AddRemoveMutex.Lock()
						udpClientRemove(UdpClient, deadIndex)
						for i := range indexSlice {
							if indexSlice[i] == deadIndex {
								fmt.Println("Clearing ", i)
								indexSlice = append(indexSlice[:i], indexSlice[i:]...)
								portNum[i].portBound = false
							}
						}
						AddRemoveMutex.Unlock()

						*clients--

					case <-time.After(5 * time.Second):
						//fmt.Println("Number of Clients - ", *clients)
					}
				}
			}(RemoveChan, clientPtr, indexSlice)

			go func(outUdpChan <-chan string, udpClientptr *[]UdpClientStruct) {
				fmt.Println("Cloner Started")
				for {
					cloneVal := <-outUdpChan
					udpRange := len(*udpClientptr)
					//fmt.Println("Number of Write Channels - ", udpRange, udpClientptr)

					for i := 0; i < udpRange; i++ {
						UdpClient[i].writeUdpChan <- cloneVal
					}
				}
			}(outUdpChan, udpClientptr)

			running = true
		}
	}

}
