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

	pkt20 := time.NewTicker(time.Millisecond * 5)

	session, err := os.Open(sessionName)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	scanner := bufio.NewScanner(session)

	for scanner.Scan() {
		select {
		case <-pkt20.C:
			//fmt.Println(scanner.Text())
			inUDPChan <- scanner.Text()
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}
