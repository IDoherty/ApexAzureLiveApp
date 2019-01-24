package aggFuncs

import (
		"log"
		"encoding/csv"
		"io"
		"os"
		"bufio"
)

func getCSV(fileName string)([]Beacon){
	
	textIn, err := os.Open(fileName)
	if err != nil{
		log.Fatal(err)
	}
	
	reader := csv.NewReader(bufio.NewReader(textIn))

	var beaconAddr []Beacon
	for {
		
		line, err := reader.Read()
		if err == io.EOF{
			break
		}else if err != nil{
			log.Fatal(err)
		}
		beaconAddr = append(beaconAddr, Beacon{
			Address	: 	line[0],
			Name	:  	line[1],
			Group	:  	line[2],
		})

	}
		
	return beaconAddr
}
