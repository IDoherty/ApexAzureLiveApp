package azureFuncs

/* 	Functions Relating to the connection between this program and Microsoft Azure. */

import (
	"fmt"

	structs "pkg/structPrototypes"

	"github.com/Azure/azure-service-bus-go"
)

type User struct {
	Id   string
	Name string
	Dist int
}

type SimpleObject struct {
	FieldA string
	FieldB int
}

//*/ AzureUpload
func AzureUpload(outAzureChan <-chan structs.AzureChanStruct, devList string) {

	go azureUpload(outAzureChan, devList)
	fmt.Println("Start Data Output to Azure")
} //*/

//*/ buildFenwayMessage
func buildFenwayMessage(tId, pId, strMess string) *servicebus.Message {

	tMsg := servicebus.NewMessageFromString(strMess)
	tMsg.Set("tId", tId) //Set team Id property
	tMsg.Set("pId", pId) // Set player Id Property

	return tMsg
} //*/

//*/ GetApexCSV
func GetApexCSV(apexList string) ([]structs.ApexLookupTable, int) {

	return getApexCSV(apexList)
} //*/
