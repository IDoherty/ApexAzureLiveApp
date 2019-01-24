// readFromFile
package aggFuncs

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func readFromFile(inUDPChan chan<- string, sessionName string) {

	//pktTicker := time.NewTicker(time.Millisecond * 50)
	pktTicker := time.NewTicker(time.Millisecond * 5)
	//pktTicker := time.NewTicker(time.Millisecond)
	//pktTicker := time.NewTicker(time.Microsecond * 10)

	session, err := os.Open(sessionName)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	scanner := bufio.NewScanner(session)

	for scanner.Scan() {
		select {
		case <-pktTicker.C:
			//fmt.Println(scanner.Text())
			inUDPChan <- scanner.Text()
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}
