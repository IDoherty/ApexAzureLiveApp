// azureFrags
package azureFuncs

import (
	"encoding/json"
	"fmt"
	"strings"

	structs "pkg/structPrototypes"
)

func azureFrags(numDevices int, azureIn structs.AzureChanStruct, devTable []structs.ApexLookupTable, fragChan chan<- structs.AzureOutputStruct) {

	//Process Incoming Metrics into Azure Fragments

	// Build Processing Structs and Flag
	var azureFragIDs structs.AzureFragID
	var azureFragment structs.AzureOutputStruct
	var devFlag int = 0

	// Fill Output Fragment Identifiers
	for x := 0; x < numDevices; x++ {
		if strings.Compare(azureIn.DevID, devTable[x].DevID) == 0 {
			azureFragIDs.DevID = azureIn.DevID
			azureFragIDs.PlayerID = devTable[x].PlayerID
			azureFragIDs.TeamID = devTable[x].TeamID
			devFlag = 1
			break
		}
		//fmt.Println(x) // Index Counter
	}

	if devFlag == 1 {
		azureFragment.RawData = azureIn.RawData

		tempval, err := json.Marshal(azureFragIDs)
		if err != nil {
			fmt.Println("Marshall Error:", err)
		}
		azureFragment.FragIDs = string(tempval)

		//*/ Filled Fragment Printout
		//fmt.Println("Filled Fragment - ", azureFragment)
		//fmt.Println()
		//*/

		fragChan <- azureFragment

		devFlag = 0
	} else {
		//fmt.Println("Device not Found - ", azureIn.DevID)
		//fmt.Println()
	}
	return
}
