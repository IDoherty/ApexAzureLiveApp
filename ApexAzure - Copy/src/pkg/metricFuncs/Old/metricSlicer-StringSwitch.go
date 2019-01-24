package metricFuncs

import (
		"fmt"
		"encoding/binary"
		//"encoding/hex"
		"bytes"
)

func metricSlicer(dataPacket []byte, MetricData *MetricStruct)(pktType int){ 
	
	// Remain using one generic Struct for all Packages?
	// Assumed All Conversions to be LSB First
	
	// Case Slice
	metricType := make([]byte, 5)
	
	metricType[0] = 224 
	metricType[1] = 225
	metricType[2] = 226
	metricType[3] = 240
	metricType[4] = 241
	
	packetType := string(dataPacket[0:1])	
				   
	switch packetType {
		
		case string(metricType[0:1]):								 
			
			// Make Struct of Packet Type Metric_Pack_0
	
			MetricData.heartRate 	= dataPacket[1:2]
			MetricData.speed 		= binary.LittleEndian.Uint16(dataPacket[2:4])
			MetricData.maxSpeed10s 	= binary.LittleEndian.Uint16(dataPacket[4:6])
			MetricData.mp			= binary.LittleEndian.Uint16(dataPacket[6:8])
			MetricData.speedIntens	= binary.LittleEndian.Uint16(dataPacket[8:10])
			MetricData.hrExert		= binary.LittleEndian.Uint16(dataPacket[10:12])
			MetricData.avgHeartrate	= dataPacket[12:13]
			
			// return Struct + Type Identifier? Add ID Field to Struct?
			pktType = 1
			
		case string(metricType[1:2]):									 
			
			//Make Struct of Packet Type Metric_Pack_1			
			
			tempDist 				:= string(dataPacket[1:4]) + "0"
			MetricData.sessDist 	= binary.LittleEndian.Uint32([]byte(tempDist))
			tempDistZ5 				:= string(dataPacket[4:7]) + "0"
			MetricData.sessDistZ5	= binary.LittleEndian.Uint32([]byte(tempDistZ5))
			tempDistZ6 				:= string(dataPacket[7:10]) + "0"
			MetricData.sessDistZ6	= binary.LittleEndian.Uint32([]byte(tempDistZ6))
			tempHMLD 				:= string(dataPacket[10:13]) + "0"
			MetricData.sessHMLD		= binary.LittleEndian.Uint32([]byte(tempHMLD))
			tempEMD		 			:= string(dataPacket[13:16]) + "0"
			MetricData.sessEMD		= binary.LittleEndian.Uint32([]byte(tempEMD))
							
			// return Struct + Type Identifier? Add ID Field to Struct?
			pktType = 2
			
		case string(metricType[2:3]):									 
			
			// Make Struct of Packet Type Metric_Pack_2			
			
			MetricData.stepBal	 	= dataPacket[1:2]
			MetricData.totalNrAccel	= binary.LittleEndian.Uint16(dataPacket[2:4])
			MetricData.totalNrDecel	= binary.LittleEndian.Uint16(dataPacket[4:6])
			MetricData.totalTimRedZ	= binary.LittleEndian.Uint16(dataPacket[6:8])
			MetricData.amp			= binary.LittleEndian.Uint16(dataPacket[8:10])
			MetricData.impacts		= binary.LittleEndian.Uint16(dataPacket[10:12])
			MetricData.dsl			= binary.LittleEndian.Uint16(dataPacket[12:14])
			
			// return Struct + Type Identifier? Add ID Field to Struct?
			pktType = 3
			
		case string(metricType[3:4]):									
			
			// Make Struct of Packet Type Accel			
			
			MetricData.lByteNrAccel		= dataPacket[1:2] 
			MetricData.durationAccel   	= binary.LittleEndian.Uint16(dataPacket[2:4])  
			MetricData.startTimAccel	= binary.LittleEndian.Uint16(dataPacket[4:6])  
			MetricData.startSpeedAccel	= binary.LittleEndian.Uint16(dataPacket[6:8])  
			MetricData.endSpeedAccel	= binary.LittleEndian.Uint16(dataPacket[8:10]) 
			MetricData.maxAccel	    	= binary.LittleEndian.Uint16(dataPacket[10:12])
			MetricData.distanceAccel	= dataPacket[12:13]
			
			// return Struct + Type Identifier? Add ID Field to Struct?
			pktType = 5
			
		case string(metricType[4:5]):									
			
			// Make Struct of Packet Type Decel								
																
			MetricData.lByteNrDecel		= dataPacket[1:2] 
			MetricData.durationDecel   	= binary.LittleEndian.Uint16(dataPacket[2:4])               	
			MetricData.startTimDecel	= binary.LittleEndian.Uint16(dataPacket[4:6])               	
			MetricData.startSpeedDecel 	= binary.LittleEndian.Uint16(dataPacket[6:8])               	
			MetricData.endSpeedDecel	= binary.LittleEndian.Uint16(dataPacket[8:10])              	
			MetricData.maxDecel	    	= binary.LittleEndian.Uint16(dataPacket[10:12])             	
			MetricData.distanceDecel	= dataPacket[12:13]             	
																  
			// return Struct + Type Identifier? Add ID Field to Struct?
			pktType = 6	
		default:
			fmt.Println("Unknown Packet Type. Identifier - ", dataPacket[0:1])
			fmt.Println()
	}
	return pktType
}	  