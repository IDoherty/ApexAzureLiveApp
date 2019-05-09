package aggFuncs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	structs "pkg/structPrototypes"
)

func GetConfigJSON() structs.ConfigStruct {

	var JSONconfig structs.JsonConfigStruct
	var configSettings structs.ConfigStruct

	raw, err := ioutil.ReadFile("Config/config.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &JSONconfig)

	fmt.Println(JSONconfig)

	// Configure Write 1
	configSettings.SessionWrite = JSONconfig.WriteAddr
	if strings.Compare(JSONconfig.SessWrite, "true") == 0 {
		configSettings.WriteOn = true
	}

	// Configure Write 2
	configSettings.SessionWrite2 = JSONconfig.WriteAddr2
	if strings.Compare(JSONconfig.SessWrite2, "true") == 0 {
		configSettings.Write2On = true
	}

	// Configure Read From Session File
	configSettings.SessionName = JSONconfig.ReadFile
	if strings.Compare(JSONconfig.SessReadIn, "true") == 0 {
		configSettings.ReadInOn = true
	}

	// Configure Read From Beacons
	configSettings.LocalAddr = JSONconfig.LocalAddr
	configSettings.BeaconAddr = JSONconfig.BeaconAddr
	if strings.Compare(JSONconfig.UDPInOn, "true") == 0 {
		configSettings.UDPInOn = true
	}

	// Configure Azure Output
	// configSettings.DevList = config[9] // CSV Dev List (OLD)
	if strings.Compare(JSONconfig.AzureOn, "true") == 0 {
		configSettings.AzureOn = true
	}

	//*/ Configure AMPQ Output
	if strings.Compare(JSONconfig.AMPQOn, "true") == 0 {
		configSettings.AMPQOn = true
	}

	/*/ Configure UDP Output - (OLD) UDP Output Flag
	if strings.Compare(config[11], "true") == 0 {
		configSettings.UDPOutOn = true
		fmt.Println("UDP Send On")
	}
	/*/

	return configSettings

}

/*/ Test Function
func main() {
    config := getConfigJSON()
    fmt.Println(config)
    for _, test := range config {
        fmt.Println(test.toString())
    }
//*/

/*/ JSON Format
type JsonConfigStruct struct {
	SessWrite  string `json:"SessWrite"`
	WriteAddr  string `json:"WriteAddr"`
	SessWrite2 string `json:"SessWrite2"`
	WriteAddr2 string `json:"WriteAddr2"`
	SessReadIn string `json:"SessReadIn"`
	ReadFile   string `json:"ReadFile"`
	LocalAddr  string `json:"LocalAddr"`
	BeaconAddr string `json:"BeaconAddr"`
	UDPInOn    string `json:"UDPInOn"`
	AzureOn    string `json:"AzureOn"`
	AMPQOn     string `json:"AMPQOn`
}
//*/

/*/ Config Struct
type ConfigStruct struct {
	SessionWrite  string
	WriteOn       bool
	SessionWrite2 string
	Write2On      bool
	SessionName   string
	ReadInOn      bool
	LocalAddr     string
	BeaconAddr    string
	UDPInOn       bool
	DevList       string
	AzureOn       bool
	UDPOutOn      bool
	AMPQOn		 bool
}
//*/
