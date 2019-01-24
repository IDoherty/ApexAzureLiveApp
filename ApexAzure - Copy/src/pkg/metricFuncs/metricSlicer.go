package metricFuncs

import (
	"bytes"
	"encoding/binary"
	//"encoding/hex"
	//"fmt"
)

func metricSlicer(dataPacket []byte, MetricData *MetricStruct) (pktType int) {

	// Assumed All Conversions to be LSB First unless otherwise stated in documentation

	// Declare Mask for Bitwise Operations

	mask3Bytes := make([]byte, 4)
	mask3Bytes[0] = 0x00
	mask3Bytes[1] = 0x00
	mask3Bytes[2] = 0xff
	mask3Bytes[3] = 0x00

	// Case Slice - Packet Identifiers for Sorting Incoming Packets
	metricType := make([]byte, 5)

	metricType[0] = 224
	metricType[1] = 225
	metricType[2] = 226
	metricType[3] = 240
	metricType[4] = 241

	if bytes.Compare(metricType[0:1], dataPacket[0:1]) == 0 {

		// Make Struct of Packet Type Metric_Pack_0

		MetricData.heartRate = dataPacket[1]
		MetricData.speed = binary.LittleEndian.Uint16(dataPacket[2:4])
		MetricData.maxSpeed10s = binary.LittleEndian.Uint16(dataPacket[4:6])
		MetricData.mp = binary.LittleEndian.Uint16(dataPacket[6:8])
		MetricData.speedIntens = binary.LittleEndian.Uint16(dataPacket[8:10])
		MetricData.hrExert = binary.LittleEndian.Uint16(dataPacket[10:12])
		MetricData.avgHeartrate = dataPacket[12]

		pktType = 1
	} else if bytes.Compare(metricType[1:2], dataPacket[0:1]) == 0 {

		//Make Struct of Packet Type Metric_Pack_1

		tempDist := dataPacket[1:4]
		MetricData.sessDist = uint32(tempDist[0])<<0 + uint32(tempDist[1])<<8 + uint32(tempDist[2])<<16
		tempDistZ5 := dataPacket[4:7]
		MetricData.sessDistZ5 = uint32(tempDistZ5[0])<<0 + uint32(tempDistZ5[1])<<8 + uint32(tempDistZ5[2])<<16
		tempDistZ6 := dataPacket[7:10]
		MetricData.sessDistZ6 = uint32(tempDistZ6[0])<<0 + uint32(tempDistZ6[1])<<8 + uint32(tempDistZ6[2])<<16
		tempHMLD := dataPacket[10:13]
		MetricData.sessHMLD = uint32(tempHMLD[0])<<0 + uint32(tempHMLD[1])<<8 + uint32(tempHMLD[2])<<16
		tempEMD := dataPacket[13:16]
		MetricData.sessEMD = uint32(tempEMD[0])<<0 + uint32(tempEMD[1])<<8 + uint32(tempEMD[2])<<16

		pktType = 2

	} else if bytes.Compare(metricType[2:3], dataPacket[0:1]) == 0 {

		// Make Struct of Packet Type Metric_Pack_2

		MetricData.stepBal = dataPacket[1]
		MetricData.totalNrAccel = binary.LittleEndian.Uint16(dataPacket[2:4])
		MetricData.totalNrDecel = binary.LittleEndian.Uint16(dataPacket[4:6])
		MetricData.totalTimRedZ = binary.LittleEndian.Uint16(dataPacket[6:8])
		MetricData.amp = binary.LittleEndian.Uint16(dataPacket[8:10])
		MetricData.impacts = binary.LittleEndian.Uint16(dataPacket[10:12])
		MetricData.dsl = binary.LittleEndian.Uint16(dataPacket[12:14])

		pktType = 3

	} else if bytes.Compare(metricType[3:4], dataPacket[0:1]) == 0 {

		// Make Struct of Packet Type Accel

		MetricData.lByteNrAccel = dataPacket[1]
		MetricData.durationAccel = binary.LittleEndian.Uint16(dataPacket[2:4])
		MetricData.startTimAccel = binary.LittleEndian.Uint16(dataPacket[4:6])
		MetricData.startSpeedAccel = binary.LittleEndian.Uint16(dataPacket[6:8])
		MetricData.endSpeedAccel = binary.LittleEndian.Uint16(dataPacket[8:10])
		MetricData.maxAccel = binary.LittleEndian.Uint16(dataPacket[10:12])
		MetricData.distanceAccel = dataPacket[12]

		pktType = 5

	} else if bytes.Compare(metricType[4:5], dataPacket[0:1]) == 0 {

		// Make Struct of Packet Type Decel

		MetricData.lByteNrDecel = dataPacket[1]
		MetricData.durationDecel = binary.LittleEndian.Uint16(dataPacket[2:4])
		MetricData.startTimDecel = binary.LittleEndian.Uint16(dataPacket[4:6])
		MetricData.startSpeedDecel = binary.LittleEndian.Uint16(dataPacket[6:8])
		MetricData.endSpeedDecel = binary.LittleEndian.Uint16(dataPacket[8:10])
		MetricData.maxDecel = binary.LittleEndian.Uint16(dataPacket[10:12])
		MetricData.distanceDecel = dataPacket[12]

		pktType = 6
	} else {
		//fmt.Println("Unknown Packet Type. Identifier - ", dataPacket[0:1])
		//fmt.Println(dataPacket)
		//fmt.Println()
	}
	return pktType
}
