// getApexCSV - Pull Apex Lookup Table from CSV file.
package azureFuncs

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"time"

	structs "pkg/structPrototypes"
)

func getApexCSV(apexBeaconInfo string) ([]structs.ApexLookupTable, int) {

	fmt.Println("Retrieving Apex Device List")

	textIn, err := os.Open(apexBeaconInfo)
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(bufio.NewReader(textIn))

	var DeviceList []structs.ApexLookupTable
	var numDevices int

	for {

		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		DeviceList = append(DeviceList, structs.ApexLookupTable{
			DevID:  line[0],
			TeamID: line[1],
			//PlayerName: line[2],
			PlayerID: line[3],
			//Flag:       line[4],
			//TeamName:   line[5],
		})

	}

	numDevices = len(DeviceList)
	//fmt.Println(numDevices)
	//fmt.Println(DeviceList)

	time.Sleep(time.Second * 2)

	return DeviceList, numDevices
}
