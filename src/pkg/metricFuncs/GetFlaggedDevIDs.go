// getApexCSV - Pull Flagged Apex Devices from CSV file.
package metricFuncs

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	//	"time"
	//	structs "pkg/structPrototypes"
)

func getFlaggedDevIDsCSV(apexBeaconInfo string) ([]string, int) {

	fmt.Println("Retrieving Flagged Apex Devices")

	textIn, err := os.Open(apexBeaconInfo)
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(bufio.NewReader(textIn))

	var DeviceList []string
	var numDevices int

	for {

		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		if line[4] == "01" {
			DeviceList = append(DeviceList, line[0])
		}

	}

	numDevices = len(DeviceList)
	fmt.Println(numDevices)
	fmt.Println(DeviceList)

	//time.Sleep(time.Second * 2)

	return DeviceList, numDevices
}
