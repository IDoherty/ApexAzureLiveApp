package metricFuncs

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	structs "pkg/structPrototypes"
)

type MetricPack0 struct {
	heartRate    byte
	speed        uint16
	maxSpeed10s  uint16
	mp           uint16
	speedIntens  uint16
	hrExert      uint16
	avgHeartrate byte
}

type MetricPack1 struct {
	sessDist   uint32
	sessDistZ5 uint32
	sessDistZ6 uint32
	sessHMLD   uint32
	sessEMD    uint32
}

type MetricPack2 struct {
	stepBal      byte
	totalNrAccel uint16
	totalNrDecel uint16
	totalTimRedZ uint16
	amp          uint16
	impacts      uint16
	dsl          uint16
}

type MetricAccel struct {
	lByteNrAccel    byte
	durationAccel   uint16
	startTimAccel   uint16
	startSpeedAccel uint16
	endSpeedAccel   uint16
	maxAccel        uint16
	distanceAccel   byte
}

type MetricDecel struct {
	lByteNrDecel    byte
	durationDecel   uint16
	startTimDecel   uint16
	startSpeedDecel uint16
	endSpeedDecel   uint16
	maxDecel        uint16
	distanceDecel   byte
}

type MetricStruct struct { // Contains All Fields Used in Various Metric Packets

	MetricPack0
	MetricPack1
	MetricPack2
	MetricAccel
	MetricDecel
}

type GPSStruct struct {
	devID     string
	gpsTime   uint32
	gpsDate   uint32
	codedLat  int32
	codedLong int32
	//codedAlt   uint16
	codedSpeed uint16
	//	magX         uint16
	//	magY         uint16
	//	magZ         uint16
	Flags uint16 //GPS Lock -> 0x8000, HR -> 0x4000, Low Bat -> 0x2000, Ext Flag -> 0x1000
}

type DecodedStruct struct {
	//	TotalDistanceZ5 float32
	//	TotalDistanceZ6 float32
	TotalDistance float32
	MaxSpeed      float32
	CurrentSpeed  float32
	//	SpeedIntens     uint16
	//	AverageSpeed float32
	TotalAccel uint16
	TotalDecel uint16
	//	StepBalLeft     float32
	//	StepBalRight    float32
	Impacts     uint16
	Dsl         uint16
	decodedLat  int32
	decodedLong int32
}

// Temporary Struct for generation of Metrics at Fenway.
// Data sored as Name:Value `JSON Designator for Name` which outputs "Designator":Value when converted to JSON
type FenwayMetricsStruct struct {
	//DevID   	uint16
	//PktNo 		uint16

	TotalDistance float32 `json:"TDist"`
	MaxSpeed      float32 `json:"MaxSp"`
	CurrentSpeed  float32 `json:"CurSp"`

	Impacts uint16 `json:"Impacts"`
	Sprints uint16 `json:"Sprints`

	//DecodedLat  uint32 `json:"Latitude"`
	//DecodedLong uint32 `json:"Longitude"`
	GPSTime uint32 `json:"GPS_Time"`
	GPSDate uint32 `json:"GPS_Date"`
}

/*/ FenwayGPSStruct
type FenwayGPSStruct struct {

}
//*/

type maxMetricLatch struct {
	DevID         string
	MaxSpeed      float32
	TotalDistance float32
	Impacts       uint16
	Sprints       uint16
	GpsTime       uint32
	GpsDate       uint32
}

func truncate(some float32) float32 {
	return float32(int(some*10) / 10)
}

func MetricFunc(metricChan <-chan string, outAzureChan chan<- structs.AzureChanStruct /*, outFileChan2 chan<- string, write2 bool, devInfo string*/) {

	// Set number of Fragments in each Packet
	nrPkts := 3 //Number of individual Metric Packets in each Datagram
	pktTypes := make([]int, 3)

	// Build Arrays for passing data into Functions
	var packetIn = make([]byte, 80)
	var gpsIn = make([]byte, 30)
	var metricPack = make([]byte, 48)

	// Declare GPS Data Struct
	var gpsData GPSStruct

	// Declare Raw Metrics Struct
	var metricRaw MetricStruct

	// Declare Decoded Metrics Struct + Max Speed Table
	var decodedMetrics DecodedStruct
	var maxMetricVals []maxMetricLatch
	var newDevFlag bool = true

	// Declare Fenway Metrics Struct - now including GPS Metrics
	var fenwayMetrics FenwayMetricsStruct
	//var fenwayGPS FenwayGPSStruct
	//var pktNo uint16

	// Declare Azure Output Structs
	var azureOut structs.AzureChanStruct
	// var gpsOut structs.AzureChanStruct

	/*/ Declare Write2 Variables
	var devTable []string
	var numDevs int
	//*/

	/*/ Write 2
	if write2 == true {
		devTable, numDevs = getFlaggedDevIDsCSV(devInfo)
	}
	//*/

	for {
		// Read in String from channel and convert to []byte
		readPacketIn := <-metricChan

		/*/ Packet Counter
		fenwayMetrics.PktNo = pktNo
		pktNo++
		//*/

		decodedHex, err := hex.DecodeString(readPacketIn)
		if err != nil {
			panic(err)
		}
		/*/ hex.Dump
		fmt.Printf("%s", hex.Dump(decodedHex))
		fmt.Println()
		//*/

		headerSlice := decodedHex[4:]
		packetIn = headerSlice[:80]

		/*/ Packet Identifiers
		SeqNo := packetIn[0:1]
		SlotID := packetIn[1:2]
		//*/

		// Slice GPS Data - Slice all 28 bytes (+ 2 devID Bytes) & pass to function
		gpsIn = packetIn[2:32]

		azureOut.DevID = gpsSlicer(gpsIn, &gpsData)

		/*/ Write Selected Packets to File
		if write2 == true {
			for x := 0; x < numDevs; x++ {
				if strings.Compare(azureOut.DevID, devTable[x]) == 0 {
					outFileChan2 <- readPacketIn
					Print Filtered Packets
					fmt.Println("Filtered Packet")
					fmt.Printf("%s", hex.Dump(decodedHex))
					fmt.Println()
				}
			}
		}
		//*/

		//fmt.Println(azureOut.DevID)

		// Slice Metric Packs
		for x := 0; x < nrPkts; x++ {
			start := (x * 16) + 32
			end := start + 17
			metricPack = packetIn[start:end]
			//fmt.Println("metric Pack ", x+1, " - ", metricPack)

			tempCnt := metricSlicer(metricPack, &metricRaw)
			pktTypes[x] = tempCnt

			/*/
			fmt.Println(tempVar)
			fmt.Println()
			//*/
		}

		metricDecoder(&metricRaw, &gpsData, &decodedMetrics)

		// Test New Max Speed Val
		for x := 0; x < len(maxMetricVals); x++ {
			if strings.Compare(maxMetricVals[x].DevID, azureOut.DevID) == 0 {

				/*/ Max Speed Test Prints
				fmt.Println("Comparing Speeds - ", azureOut.DevID)
				fmt.Println("Current Max - ", maxSpeedVals[x].MaxSpeed)
				fmt.Println("New Test - ", decodedMetrics.MaxSpeed)
				//*/

				if decodedMetrics.MaxSpeed > maxMetricVals[x].MaxSpeed {
					maxMetricVals[x].MaxSpeed = decodedMetrics.MaxSpeed
				} else {
					decodedMetrics.MaxSpeed = maxMetricVals[x].MaxSpeed
				}

				if decodedMetrics.TotalDistance > maxMetricVals[x].TotalDistance {
					maxMetricVals[x].TotalDistance = decodedMetrics.TotalDistance // truncate((decodedMetrics.TotalDistance + maxMetricVals[x].TotalDistance) / 2)
				} else {
					decodedMetrics.TotalDistance = maxMetricVals[x].TotalDistance
				}

				if decodedMetrics.Impacts > maxMetricVals[x].Impacts {
					maxMetricVals[x].Impacts = decodedMetrics.Impacts
				} else {
					decodedMetrics.Impacts = maxMetricVals[x].Impacts
				}

				if decodedMetrics.TotalAccel > maxMetricVals[x].Sprints {
					maxMetricVals[x].Sprints = decodedMetrics.TotalAccel
				} else {
					decodedMetrics.TotalAccel = maxMetricVals[x].Sprints
				}

				if gpsData.gpsTime > maxMetricVals[x].GpsTime {
					maxMetricVals[x].GpsTime = gpsData.gpsTime
				} else {
					gpsData.gpsTime = maxMetricVals[x].GpsTime
				}

				if gpsData.gpsDate > maxMetricVals[x].GpsDate {
					maxMetricVals[x].GpsDate = gpsData.gpsDate
				} else {
					gpsData.gpsDate = maxMetricVals[x].GpsDate
				}

				newDevFlag = false
				break
			}
		}

		if newDevFlag == true {
			//fmt.Println("New Packet")
			maxMetricVals = append(maxMetricVals, maxMetricLatch{
				DevID:         azureOut.DevID,
				MaxSpeed:      decodedMetrics.MaxSpeed,
				TotalDistance: decodedMetrics.TotalDistance,
				Impacts:       decodedMetrics.Impacts,
				Sprints:       decodedMetrics.TotalAccel,
			})
		} else {
			newDevFlag = true
		}

		fenwayMetrics.MaxSpeed = decodedMetrics.MaxSpeed
		fenwayMetrics.CurrentSpeed = decodedMetrics.CurrentSpeed
		fenwayMetrics.TotalDistance = decodedMetrics.TotalDistance
		fenwayMetrics.Impacts = decodedMetrics.Impacts
		fenwayMetrics.Sprints = decodedMetrics.TotalAccel

		//fmt.Printf("Current Speed %v m/s - ", fenwayMetrics.CurrentSpeed)

		/*/ GPS Pack
		fenwayGPS.DecodedLat = decodedMetrics.decodedLat
		fenwayGPS.DecodedLong = decodedMetrics.decodedLong
		fenwayGPS.GPSDate = gpsData.gpsDate
		fenwayGPS.GPSTime = gpsData.gpsTime
		//*/

		metricJSON, err := json.Marshal(fenwayMetrics)
		if err != nil {
			fmt.Println("error:", err)
		}

		//*/ outAzureChan
		azureOut.RawData = string(metricJSON)
		outAzureChan <- azureOut
		//*/

		/*/ outUDPChan
		fmt.Println(string(jsonOut))
		fmt.Println()

		outUDPChan <- string(jsonOut)
		//*/

		/*/ Output tests
		fmt.Println("SeqNo 		- ", SeqNo)
		fmt.Println("SlotID 		- ", SlotID)
		fmt.Println("devID 		- ", devID)
		fmt.Println()
		//*/

		/*/ GPS Data
		fmt.Println("devID		- ", gpsData.devID)
		fmt.Println("gpsTime		- ", gpsData.gpsTime)
		fmt.Println("gpsDate		- ", gpsData.gpsDate)
		fmt.Println("codedLat	- ", gpsData.codedLat)
		fmt.Println("codedLong	- ", gpsData.codedLong)
		fmt.Println("codedAlt	- ", gpsData.codedAlt)
		fmt.Println("codedSpeed	- ", gpsData.codedSpeed)
		//fmt.Println("magX		- ", gpsData.magX)
		//fmt.Println("magY		- ", gpsData.magY)
		//fmt.Println("magZ		- ", gpsData.magZ)
		fmt.Println("Flags		- ", gpsData.Flags)
		fmt.Println()
		//*/

		/*/ Metric Pack 0
		//fmt.Println("heartRate	- ", metricRaw.heartRate)
		//fmt.Println("speed		- ", metricRaw.speed)
		fmt.Println("maxSpeed10s	- ", metricRaw.maxSpeed10s)
		//fmt.Println("mp		- ", metricRaw.mp)
		fmt.Println("speedIntens	- ", metricRaw.speedIntens)
		//fmt.Println("hrExert		- ", metricRaw.hrExert)
		//fmt.Println("avgHeartrate	- ", metricRaw.avgHeartrate)
		fmt.Println()
		//*/

		/*/ Metric Pack 1
		//
		fmt.Println("sessDist	- ", metricRaw.sessDist)
		fmt.Println("sessDistZ5	- ", metricRaw.sessDistZ5)
		fmt.Println("sessDistZ6	- ", metricRaw.sessDistZ6)
		fmt.Println("sessHMLD	- ", metricRaw.sessHMLD)
		fmt.Println("sessEMD	- ", metricRaw.sessEMD)
		fmt.Println()
		//*/

		/*/ Metric Pack 2
		fmt.Println("stepBal 	- ", metricRaw.stepBal)
		fmt.Println("totalNrAccel 	- ", metricRaw.totalNrAccel)
		fmt.Println("totalNrDecel 	- ", metricRaw.totalNrDecel)
		//fmt.Println("totalTimRedZ 	- ", metricRaw.totalTimRedZ)
		//fmt.Println("amp 		- ", metricRaw.amp)
		fmt.Println("impacts 	- ", metricRaw.impacts)
		fmt.Println("dsl 		- ", metricRaw.dsl)
		fmt.Println()
		//*/

		/*/ Accel Pack
		fmt.Println("lByteNrAccel 		- ", metricRaw.lByteNrAccel)
		//fmt.Println("durationAccel 		- ", metricRaw.durationAccel)
		//fmt.Println("startTimAccel 		- ", metricRaw.startTimAccel)
		//fmt.Println("startSpeedAccel 	- ", metricRaw.startSpeedAccel)
		//fmt.Println("endSpeedAccel 	- ", metricRaw.endSpeedAccel)
		//fmt.Println("maxAccel 		- ", metricRaw.maxAccel)
		fmt.Println("distanceAccel 		- ", metricRaw.distanceAccel)
		fmt.Println()
		//*/

		/*/ Decel Pack
		fmt.Println("lByteNrDecel 	- ", metricRaw.lByteNrDecel)
		//fmt.Println("durationDecel 	- ", metricRaw.durationDecel)
		//fmt.Println("startTimDecel 	- ", metricRaw.startTimDecel)
		//fmt.Println("startSpeedDecel - ", metricRaw.startSpeedDecel)
		//fmt.Println("endSpeedDecel 	- ", metricRaw.endSpeedDecel)
		//fmt.Println("maxDecel	- ", metricRaw.maxDecel)
		fmt.Println("distanceDecel 	- ", metricRaw.distanceDecel)
		fmt.Println()
		//*/

		/*/ Metric Decoder Prototype
		fmt.Printf("Distance - %.2fm \r\n", decodedMetrics.totalDistance)
		fmt.Printf("Distance - %.2fKm \r\n", decodedMetrics.totalDistance/1000)
		fmt.Println()

		fmt.Printf("Zone 5 Distance - %.2fm \r\n", decodedMetrics.totalDistanceZ5)
		fmt.Printf("Zone 5 Distance - %.2fKm \r\n", decodedMetrics.totalDistanceZ5/1000)
		fmt.Println()

		fmt.Printf("Zone 6 Distance - %.2fm \r\n", decodedMetrics.totalDistanceZ6)
		fmt.Printf("Zone 6 Distance - %.2fKm \r\n", decodedMetrics.totalDistanceZ6/1000)
		fmt.Println()

		fmt.Printf("Max Speed (10s) - %.2fm/s \r\n", decodedMetrics.maxSpeed)
		fmt.Printf("Speed Intensity - %d \r\n", decodedMetrics.speedIntens)
		fmt.Println()

		fmt.Printf("Accelerations - %d \r\n", decodedMetrics.totalAccel)
		fmt.Printf("Decelerations - %d \r\n", decodedMetrics.totalDecel)
		fmt.Println()

		fmt.Printf("Step Balance Left - %.2f \r\n", decodedMetrics.stepBalLeft)
		fmt.Printf("Step Balance Left - %.2f \r\n", decodedMetrics.stepBalRight)
		fmt.Println()

		fmt.Printf("impacts - %d \r\n", decodedMetrics.impacts)
		fmt.Printf("dsl - %d \r\n", decodedMetrics.dsl)
		fmt.Println()
		//*/

		//fmt.Println()
	}
}
