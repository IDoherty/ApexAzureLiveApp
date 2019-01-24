package aggFuncs

import (
        //"fmt"
)

func testValidity(lastVals *LastPacket, testNo byte, testTime uint32)(bool, int){
	
	var flag bool
	var pktType int
	var timFlag bool
	
	if lastVals.gpsTime < testTime{ 
		timFlag = true
	}
	
	switch{
		
		case lastVals.seqNo < testNo && lastVals.seqNo >= 2 && timFlag:
			
			//fmt.Println(lastVals.seqNo," => ",testNo)
			lastVals.seqNo = testNo
			lastVals.gpsTime = testTime
			//fmt.Println()
			flag = true
			pktType = 1
			
		case lastVals.seqNo > 253 && testNo < 2 && timFlag:
	
			//fmt.Println(lastVals.seqNo," => ",testNo)
			lastVals.seqNo = testNo
			lastVals.gpsTime = testTime
			//fmt.Println()
			flag = true
			pktType = 1
			
		case lastVals.seqNo < testNo && testNo <= 253 && timFlag:
			
			//fmt.Println(lastVals.seqNo," => ",testNo)
			lastVals.seqNo = testNo
			lastVals.gpsTime = testTime
			//fmt.Println()
			flag = true
			pktType = 1
		
		case lastVals.seqNo == testNo && lastVals.gpsTime == testTime:
	
			//fmt.Println("Same Packet == ", lastVals.seqNo)
			//fmt.Println()
			flag = false
			pktType = 2
			
		case lastVals.seqNo > testNo && lastVals.gpsTime > testTime:
	
			//fmt.Println("Old Packet < ", lastVals.seqNo)
			//fmt.Println()
			flag = false
			pktType = 3
	
	}
	return flag, pktType
}
