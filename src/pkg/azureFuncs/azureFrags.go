// azureFrags
package azureFuncs

import (
	"encoding/json"
	"fmt"

	//	"strings"

	structs "pkg/structPrototypes"
)

func azureFrags(azureIn structs.AzureChanStruct, fragChan chan<- structs.AzureOutputStruct) {

	//Process Incoming Metrics into Azure Fragments

	// Build Processing Structs and Flag
	var azureFragIDs structs.AzureFragID
	var azureFragment structs.AzureOutputStruct
	//	var devFlag int = 0

	// Fill Output Fragment Identifiers
	azureFragIDs.DevID = azureIn.DevID

	azureFragment.RawData = azureIn.RawData

	tempval, err := json.Marshal(azureFragIDs)
	if err != nil {
		fmt.Println("Marshall Error:", err)
	}
	azureFragment.FragIDs = string(tempval)

	/*/ Filled Fragment Printout
	fmt.Println("Filled Fragment - ", azureFragment)
	fmt.Println()
	//*/
	fragChan <- azureFragment

	return
}
