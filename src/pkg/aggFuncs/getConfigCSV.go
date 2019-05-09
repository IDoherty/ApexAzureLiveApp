package aggFuncs

import (
	"bufio"
	"encoding/csv"

	"fmt"
	"log"
	"os"
	"strings"

	structs "pkg/structPrototypes"
)

func GetConfigCSV() structs.ConfigStruct {

	//*/
	textIn, err := os.Open("Config/config.txt")
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(bufio.NewReader(textIn))

	var configSettings structs.ConfigStruct
	configSettings.ReadInOn = false
	configSettings.UDPInOn = false
	configSettings.WriteOn = false
	configSettings.Write2On = false
	configSettings.AzureOn = false
	configSettings.UDPOutOn = false

	config, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	// Configure Write 1
	configSettings.SessionWrite = config[0]
	if strings.Compare(config[1], "true") == 0 {
		configSettings.WriteOn = true
	}

	// Configure Write 2
	configSettings.SessionWrite2 = config[2]
	if strings.Compare(config[3], "true") == 0 {
		configSettings.Write2On = true
	}

	// Configure Read From Session File
	configSettings.SessionName = config[4]
	if strings.Compare(config[5], "true") == 0 {
		configSettings.ReadInOn = true
	}

	// Configure Read From Beacons
	configSettings.LocalAddr = config[6]
	configSettings.BeaconAddr = config[7]
	if strings.Compare(config[8], "true") == 0 {
		configSettings.UDPInOn = true
	}

	// Configure Azure Device List and Settings
	configSettings.DevList = config[9]
	if strings.Compare(config[10], "true") == 0 {
		configSettings.AzureOn = true
	}

	// Configure UDP Output
	if strings.Compare(config[11], "true") == 0 {
		configSettings.UDPOutOn = true
		fmt.Println("UDP Send On")
	}
	//*/
	return configSettings
}
