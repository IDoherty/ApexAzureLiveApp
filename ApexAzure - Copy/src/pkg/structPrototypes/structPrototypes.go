package structPrototypes

// AzureFuncs Structs

// Azure Input Struct Containing Raw Metrics from the Processed Packets
type AzureChanStruct struct {
	DevID   string
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
	DevID    string
	TeamID   string `json:"tID"`
	PlayerID string `json:"pID"`
}

// Output Structure containing JSON Strings - Values are Concatenated to form a Fragment of the Output Packet
type AzureOutputStruct struct {
	FragIDs string
	RawData string
}
