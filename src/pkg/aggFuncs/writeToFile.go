package aggFuncs

import (
	//	"encoding/hex"
	"fmt"
	"os"
	"time"
)

func writeToFile(outJsonChan <-chan string, session string) {

	newline := []byte("\n")
	//var pktcnt int

	// Take Current Time + Date
	sessionTime := time.Now()
	formattedTime := sessionTime.Format("2006-01-02 15-04-05")

	SessionName := session + formattedTime
	fmt.Println()
	fmt.Println("Recording Session Data to - ", SessionName)

	//Build File
	file, err := os.Create(SessionName)
	if err != nil {
		fmt.Printf("Bad File")
	}

	defer file.Close()

	for {

		packetIn := <-outJsonChan

		/*/ Incoming Packet Printout
		fmt.Println(packetIn)
		destringifiedData, _ := hex.DecodeString(packetIn)

		fmt.Println(pktcnt)
		pktcnt++
		//*/

		file.Write([]byte(packetIn))
		file.Write(newline)
	}
}
