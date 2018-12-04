package aggFuncs

import (
        //"fmt"
)

func testValidity(lastVals *LastPacket, testNo byte, testTime uint32)(bool, int){
	
	var flag bool
	var pktType int
	
	if lastVals.seqNo < testNo && lastVals.gpsTime < testTime{
		if lastVals.seqNo >= 5{ 
			//fmt.Println(lastVals.seqNo," => ",testNo)
			lastVals.seqNo = testNo
			lastVals.gpsTime = testTime
			//fmt.Println()
			flag = true
			pktType = 1
			
		}else if testNo <= 250{
			//fmt.Println(lastVals.seqNo," => ",testNo)
			lastVals.seqNo = testNo
			lastVals.gpsTime = testTime
			//fmt.Println()
			flag = true
			pktType = 1
			
		}
	}else if lastVals.seqNo > 250 && testNo < 5 && lastVals.gpsTime < testTime{
	
		//fmt.Println(lastVals.seqNo," => ",testNo)
		lastVals.seqNo = testNo
		lastVals.gpsTime = testTime
		//fmt.Println()
		flag = true
		pktType = 1
		
	}else if lastVals.seqNo == testNo && lastVals.gpsTime == testTime{
	
		//fmt.Println("Same Packet == ", lastVals.seqNo)
		//fmt.Println()
		flag = false
		pktType = 2
		
	}else if lastVals.seqNo > testNo && lastVals.gpsTime == testTime{
	
		//fmt.Println("Old Packet < ", lastVals.seqNo)
		//fmt.Println()
		flag = false
		pktType = 3
	}
	
	return flag, pktType
}