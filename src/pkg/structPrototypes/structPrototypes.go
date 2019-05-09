package structPrototypes

// AzureFuncs Structs

//*/
// Configuration Struct
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
	AMPQOn        bool
}

//*/

/*/ JSON Config Notes

Use 'true' or 'false' for binary values
IP Address format is "xxx.xxx.xxx.xxx",though values less than 100 need no additional values


"SessWrite":"true",
"SessWrite2":"true",
"SessReadIn":"true",
"ReadFile":"NameOfFile",
"LocalAddr":"IPAddress",
"BeaconAddr":"BIPAddress",
"UDPInOn":"true",
"AzureOn":"true"
//*/

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

// Azure Input Struct Containing Raw Metrics from the Processed Packets
type AzureChanStruct struct {
	DevID   string
	Unix    string
	RawData string
}

// Lookup Table for Apex Devices, Teams and Player Names/Numbers
type ApexLookupTable struct {
	DevID  string
	TeamID string
	//PlayerName string
	PlayerID string
	//Flag       string
	//TeamName   string
}

// Azure Fragment Identifiers Struct built from Input Struct and Values in the Lookup Table
type AzureFragID struct {
	DevID string
}

// Azure Fragment Identifiers Struct built from Input Struct and Values in the Lookup Table
type AzureFragIDDevList struct {
	DevID string
	//	TeamID   string `json:"tID"`
	//	PlayerID string `json:"pID"`
}

// Output Structure containing JSON Strings - Values are Concatenated to form a Fragment of the Output Packet
type AzureOutputStruct struct {
	FragIDs string
	Unix    string
	RawData string
}

type UnixStruct struct {
	Unix string `json:"Unix"`
}

type SecArray struct {
	LastSec []string
	CurrSec []string
}
