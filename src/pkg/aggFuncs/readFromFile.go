// readFromFile
package aggFuncs

import (
	"bufio"
	"encoding/binary"
	"encoding/hex"

	//	"fmt"
	"log"
	"os"
	"time"

	structs "pkg/structPrototypes"
)

func readFromFile(inUDPChan chan<- string, sessionName string) {

	// Declare Variable Array
	//
	// Start Ticker
	//pkt5ms := time.NewTicker(time.Millisecond * 5)
	pkt1sec := time.NewTicker(time.Second * 1)

	// Declare Ticker Channel (Package Buffer)
	tickChan := make(chan []string, 5)

	go func(tickChan chan<- []string, sessionName string) {

		var inArrays structs.SecArray

		var timSec uint32 = 0
		var curTim uint32 = 0

		session, err := os.Open(sessionName)
		if err != nil {
			log.Fatal(err)
		}
		defer session.Close()

		scanner := bufio.NewScanner(session)

		for scanner.Scan() {
			// Read In Packet from File
			IntakeVal := scanner.Text()

			// Decode Values
			decodedVal, err := hex.DecodeString(IntakeVal)
			if err != nil {
				panic(err)
			}

			// Calculate TimeStamp (Second)
			gpsTimeRaw := binary.BigEndian.Uint32(decodedVal[8:14])
			timSec = gpsTimeRaw / 1000
			//fmt.Println(timSec)

			if timSec == curTim {
				inArrays.CurrSec = append(inArrays.CurrSec, IntakeVal)
			} else if timSec == (curTim - 1) {
				inArrays.LastSec = append(inArrays.LastSec, IntakeVal)
			} else if timSec > curTim {
				// Print and Send LastSec
				//fmt.Println(curTim, inArrays.LastSec)
				tickChan <- inArrays.LastSec
				// Clear LastSec Array
				inArrays.LastSec = []string{}
				// Shift CurrSec Vals to LastSec
				inArrays.LastSec = inArrays.CurrSec
				// Clear Old CurrSec
				inArrays.CurrSec = []string{}
				// Print and Update CurTim Value
				//fmt.Println(curTim)
				curTim = timSec
			} else {
				// If invalid Value for CurTim
				IntakeVal = "clear"
			}
		}
		/*/
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
		//*/
	}(tickChan, sessionName)

	for {
		select {
		case <-pkt1sec.C:
			// Recieve Packets with given Second's Timestamp from tickChan
			secArray := <-tickChan

			x := 0
			for x < len(secArray) {
				inUDPChan <- secArray[x]
				x++
			}
			secArray = []string{}
		}
	}
}
