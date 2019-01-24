package aggFuncs

import (
	//"fmt"
	"log"

	//"time"
	"encoding/binary"
	"encoding/hex"
)

func procPackets(inUDPChan <-chan string, outFileChan chan<- string, metricChan chan<- string, keepAlive []byte, writeVals bool) {

	// Build Variables
	var arrayLastPackets [50]LastPacket
	//valPkt:= 0;

	var fileData string
	var metricData string

	keepAlivePacketID := make([]byte, 2)

	keepAlivePacketID[0] = 0x55
	keepAlivePacketID[1] = 0xdd

	packetTest := binary.LittleEndian.Uint16(keepAlivePacketID)

	//fmt.Println(packetTest)
	//fmt.Println()

	for {
		returnedData := <-inUDPChan

		// Revert Data to []byte
		destringifiedData, err := hex.DecodeString(returnedData)
		if err != nil {
			log.Fatal(err)
		}

		// Filter out KA responses and test Validity for incoming packets
		if binary.LittleEndian.Uint16(destringifiedData[0:2]) == packetTest {

			/*/ Hex Dump of Packet
			//fmt.Println("Keep Alive Packet: Processing")
			//fmt.Printf("%s", hex.Dump(destringifiedData))
			//fmt.Println()
			//*/

		} else {
			seqNo := destringifiedData[4:5]
			slotNo := destringifiedData[5:6]
			gpsTime := binary.BigEndian.Uint32(destringifiedData[8:12])
			//devID := hex.EncodeToString(destringifiedData[6:8])

			/*/ Packet Identifier Slices
			fmt.Printf("%s", hex.Dump(destringifiedData))
			fmt.Println()

			//fmt.Println(slotNo[0])
			//fmt.Println(seqNo[0])
			fmt.Println(devID)
			fmt.Println()
			//*/

			valid, pktType := TestValidity(&arrayLastPackets[slotNo[0]], seqNo[0], gpsTime)
			//fmt.Println(valid)
			//countChan <- pktType
			pktType++

			if valid {
				// Valid Packet Counter
				//valPkt++

				/*/ Channel to UDP Functions
				outData := returnedData
				outUDPChan <- outData
				//*/

				//*/ Channel to Metric Functions
				metricData = returnedData
				metricChan <- metricData
				//fmt.Println(metricData)
				//*/

				//*/ Channel to Output File Writer
				if writeVals {
					fileData = returnedData
					outFileChan <- fileData
				}
				//*/
			}

		}

	}

}
