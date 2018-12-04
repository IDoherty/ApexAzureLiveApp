package azureFuncs

import (
	"fmt"
	"time"

	/*/ Basic Libraries
	"context"
	"strings"
	//*/
	/*/ Encoding Libraries
	"encoding/hex"
	"encoding/json"
	//*/
	/*/ os/io Libraries
	"os"
	"os/signal"
	"io"
	//*/
	/*/ Azure Libraries
	"github.com/Azure/azure-amqp-common-go/persist"
	"github.com/Azure/azure-event-hubs-go"
	"github.com/Azure/azure-service-bus-go"
	//*/

	structs "pkg/structPrototypes"
)

func azureUpload(outAzureChan <-chan structs.AzureChanStruct, devList string) {

	var azureIn structs.AzureChanStruct
	/*/ Unused Variables
	var tMsg *servicebus.Message
	var strMsg string

	var azureFragIDs structs.AzureFragID
	var azureFragment structs.AzureOutputStruct
	var devFlag int = 0
	//*/

	fmt.Printf("Start Fenway Project\n")

	// Build Lookup Table for Devices

	devTable, numDevices := getApexCSV(devList)

	fmt.Println()
	fmt.Println("Total Devices Loaded In - ", numDevices)
	fmt.Println(devTable)

	/*/ Azure Connection Prototype - used in sendAzurePacket
	connStr := "Endpoint=sb://fenwaybus.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=njO049+bG5AY6QhmXIwu+mU6InX8dybImistGVobcR0="
	if connStr == "" {
		fmt.Println("FATAL: expected environment variable SERVICEBUS_CONNECTION_STRING not set")
		return
	}

	// Create a client to communicate with a Service Bus Namespace.
	ns, err := servicebus.NewNamespace(servicebus.NamespaceWithConnectionString(connStr))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a client to communicate with the queue. (The queue must have already been created, see `QueueManager`)
	topic, err := ns.NewTopic("playerFrank")
	if err != nil {
		fmt.Println("FATAL: ", err)
		return
	}
	fmt.Println("Topic Name", topic.Name)
	//*/

	fragChan := make(chan structs.AzureOutputStruct, 64)
	strMsgChan := make(chan string, 256)

	// Start Output Thread
	go sendAzurePacket(strMsgChan)

	// Start Ticker for Packet Building
	buildPacket := time.NewTicker(time.Second)

	for {

		select {
		case azureIn = <-outAzureChan:
			/*/ Test Printouts
			//fmt.Println(azureIn.DevID)
			//fmt.Println(devTable[0].DevID)
			//fmt.Println(devTable[1].DevID)
			//fmt.Println("Raw data - ", azureIn.RawData)
			//fmt.Println("GPS data - ", azureIn.GPSData)
			//fmt.Println()

			//fmt.Println("azureIn value - ", azureIn.DevID)
			//*/

			// Format Fragments
			go azureFrags(numDevices, azureIn, devTable, fragChan)

			/*/ AzureFrags Prototype Code
				// Fill Output Fragment
				//go func(azureIn structs.AzureChanStruct, devTable []structs.ApexLookupTable, fragChan chan<- structs.AzureChanStruct) {
				cnt := 0
				for x := 0; x < numDevices; x++ {
					if strings.Compare(azureIn.DevID, devTable[cnt].DevID) == 0 {
						azureFragIDs.DevID = azureIn.DevID
						azureFragIDs.PlayerID = devTable[cnt].PlayerID
						azureFragIDs.TeamID = devTable[cnt].TeamID
						//azureIn.PlayerName = devTable[cnt].PlayerName
						devFlag = 1
						break
					}
					cnt++
					//fmt.Println(x)
				}

				if devFlag == 1 {
					azureFragment.RawData = azureIn.RawData

					tempval, err := json.Marshal(azureFragIDs)
					if err != nil {
						fmt.Println("Marshall Error:", err)
					}
					azureFragment.FragIDs = string(tempval)

					fmt.Println("Filled Fragment - ", azureFragment)
					fmt.Println()

					fragChan <- azureFragment

					devFlag = 0
				} else {
					fmt.Println("Device not Found")
					fmt.Println()
				}
				//}(azureIn, devTable, fragChan)
			//*/

		case <-buildPacket.C:
			/*/ Build Packet from Fragments
			fmt.Println("New Build Thread")
			fmt.Println()
			//*/
			go azurePacketBuilder(fragChan, strMsgChan)
		}
	}

	/*/ Send Azure Packet Prototype Code
	////////////////   Limerick
	messageTimeout := time.Now().Add(3000 * time.Millisecond)

	strMsg = <- strMsgChan
	tMsg = buildFenwayMessage("Limerick", "01", strMsg)

	go func(tMsg *servicebus.Message) {
		ctx, cancel := context.WithDeadline(context.Background(), messageTimeout)
		defer cancel()
			err := topic.Send(ctx, strMsgChan)
		if err != nil {
			fmt.Println("Error Sending to Topic: ", err)
		}
	}(tMsg)
	}


	////
	//fmt.Println("Context err", ctx.Err())
	//*/

	/*/ Event Hub Prototype Code
		if err != nil {
			fmt.Printf("Error: Connecting to Event Hub : %s\n", err)
			os.Exit(1)
		} else {
			fmt.Printf("Event Hub Has been Created\n")
		}
		//hub.Close(context.Background())

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		for i := 0; i < 5; i++ {
			now := time.Now()

			user := &User{Name: "Frank", Id: now.Format("2006-01-02T15:04:05.999999"), Dist: i}
			bMessage, err := json.Marshal(user)
			if err != nil {
				fmt.Printf("Error: %s", err)
				return
			}
			strMessage := string(bMessage)
			fmt.Println(strMessage)

			err = hub.Send(ctx, eventhub.NewEventFromString(strMessage))

			if err != nil {
				fmt.Printf("Error: Sending data to Event Hub : %s\n", err)
				hub.Close(context.Background())
				os.Exit(2)
			}

			time.Sleep(1000 * time.Millisecond) //every second
		}
	//*/

	fmt.Printf("Finished Do I need to clean up anything?\n")

}
