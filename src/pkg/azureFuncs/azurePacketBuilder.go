// azurePacketBuilder
package azureFuncs

import (
	//	"encoding/json"
	"fmt"
	"strings"
	"time"

	structs "pkg/structPrototypes"
)

func azurePacketBuilder(fragChan <-chan structs.AzureOutputStruct, strMsgChan chan<- string) {
	builderTimeout := time.NewTimer(1000 * time.Millisecond)

	var fragIn structs.AzureOutputStruct
	var azurePacket []structs.AzureOutputStruct
	var newDevFlag bool = true
	var strMsg string
	//var strMsgTemp string

	for {
		select {
		case <-builderTimeout.C:

			/*/ String Message Printout
			strMsgTemp = azurePacket[0].FragIDs
			fmt.Println("input frag - ", azurePacket[0].FragIDs)
			fmt.Println("strMsgTemp - ", strMsgTemp)
			//*/

			//*/
			for y := 0; y < len(azurePacket); y++ {
				//*/
				tempbyte := []byte(azurePacket[y].FragIDs)
				azurePacket[y].FragIDs = string(tempbyte[:len(tempbyte)-1])

				tempdata := []byte(azurePacket[y].RawData)
				tempdata[0] = 44
				azurePacket[y].RawData = string(tempdata)
				//*/

				strMsg = strMsg + azurePacket[y].FragIDs + azurePacket[y].RawData

				//*/
				if y != len(azurePacket)-1 {
					strMsg += ","
				}
				//*/
			}
			strMsg = "[" + strMsg + "]"
			fmt.Println("Packet Payload - ", strMsg)
			fmt.Println(len(azurePacket))
			fmt.Println()
			//*/

			/*/ Packet Formatting Prototype
			tempbyte := []byte(strMsgTemp)
			comma := []byte(",")
			fmt.Println("tempbyte - ", tempbyte)
			fmt.Println("comma - ", comma)
			templen := len(tempbyte)
			fmt.Println("templen - ", templen)
			fmt.Println(tempbyte[templen-1])
			tempbyte[templen-1] = 44
			strMsgTemp = string(tempbyte)
			//*/

			strMsgChan <- strMsg
			return

		case fragIn = <-fragChan:

			for x := 0; x < len(azurePacket); x++ {
				if strings.Compare(azurePacket[x].FragIDs, fragIn.FragIDs) == 0 {
					//fmt.Println("New Values - ", fragIn.FragIDs)
					azurePacket[x] = fragIn
					newDevFlag = false
					break
				}
			}

			if newDevFlag == true {
				//fmt.Println("New Device - ", fragIn.FragIDs)
				azurePacket = append(azurePacket, structs.AzureOutputStruct{
					FragIDs: fragIn.FragIDs,
					RawData: fragIn.RawData,
				})
				//fmt.Println("Overall Struct - ", azurePacket)
				//fmt.Println()
			} else {
				newDevFlag = true
			}

		}
	}
}
