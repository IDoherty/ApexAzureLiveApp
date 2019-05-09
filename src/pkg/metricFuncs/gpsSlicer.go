package metricFuncs

import (
	"encoding/binary"
	"encoding/hex"

	//	"fmt"
	"strconv"
	"time"
)

func Truncate(some float32) float32 {
	return float32(int(some*100)) / 100
}

func gpsSlicer(dataPacket []byte, gpsData *GPSStruct) {

	//*/ DevID - Uint16 or String Formats (pick one)
	//gpsData.devID = binary.LittleEndian.Uint16(dataPacket[0:2])
	gpsData.devID = hex.EncodeToString(dataPacket[0:2])
	//*/

	//*/ GPS Time
	gpsTimeRaw := binary.BigEndian.Uint32(dataPacket[2:6])
	gpsData.gpsTime = gpsTimeRaw
	//fmt.Println(gpsTimeRaw)

	gpsTimeHour := gpsTimeRaw / 10000000
	gpsTimeMin := gpsTimeRaw/100000 - (gpsTimeHour * 100)
	gpsTimeSec := gpsTimeRaw/1000 - (gpsTimeHour * 10000) - (gpsTimeMin * 100)
	gpsData.MilliSec = (gpsTimeRaw - (gpsTimeHour * 10000000) - (gpsTimeMin * 100000) - (gpsTimeSec * 1000)) / 100
	//*/

	/*/
	fmt.Println("Hour - ", gpsTimeHour)
	fmt.Println("Min - ", gpsTimeMin)
	fmt.Println("Sec - ", gpsTimeSec)
	fmt.Println("ms - ", gpsTimeMili)
	fmt.Println()
	//*/

	//*/ GPS Date
	tempData := dataPacket[6:10]
	gpsDateRaw := uint32(tempData[1])<<16 + uint32(tempData[2])<<8 + uint32(tempData[3])<<0
	gpsData.gpsDate = gpsDateRaw

	gpsDateDay := gpsDateRaw / 10000
	gpsDateMon := gpsDateRaw/100 - (gpsDateDay * 100)
	gpsDateYear := gpsDateRaw - (gpsDateDay * 10000) - (gpsDateMon * 100)

	gpsData.GpsHAcc = uint32(tempData[0])
	//*/

	/*/
	fmt.Println("Day - ", gpsDateDay)
	fmt.Println("Month - ", gpsDateMon)
	fmt.Println("Year - ", gpsDateYear)
	fmt.Println()
	//*/

	//*/ Convert To UTC and Unix
	const layout = "2006-1-2T15:4:5.0Z"

	UTCStr := "20" + strconv.Itoa(int(gpsDateYear)) + "-" + strconv.Itoa(int(gpsDateMon)) + "-" + strconv.Itoa(int(gpsDateDay)) + "T" + strconv.Itoa(int(gpsTimeHour)) + ":" + strconv.Itoa(int(gpsTimeMin)) + ":" + strconv.Itoa(int(gpsTimeSec)) + "." + strconv.Itoa(int(gpsData.MilliSec)) + "Z"
	//fmt.Println(UTCStr)

	gpsData.UTCTime, _ = time.Parse(layout, UTCStr)
	//if err != nil {
	//	fmt.Println("error", err)
	//}

	gpsData.UnixTime = gpsData.UTCTime.Unix()
	//*/

	/*/ Print Formatted Time and Date
	fmt.Println("UTC Format - ", gpsData.UTCTime)
	fmt.Println("Unix - ", gpsData.UnixTime)
	fmt.Println("ms - ", gpsData.MilliSec)
	fmt.Println()
	//*/

	//*/ Coded Latitude, Longitude & Altitude
	tempULat := binary.LittleEndian.Uint32(dataPacket[10:14])
	gpsData.codedLat = int32(tempULat)

	tempULong := binary.LittleEndian.Uint32(dataPacket[14:18])
	gpsData.codedLong = int32(tempULong)

	//fmt.Println(tempULat)
	//fmt.Println(tempULong)

	//gpsData.codedAlt = binary.LittleEndian.Uint16(dataPacket[18:20])
	//*/

	//*/ Coded Speed
	gpsData.codedSpeed = binary.LittleEndian.Uint16(dataPacket[20:22])
	//*/

	/*/ Magnetometer Metrics
	gpsData.magX = binary.LittleEndian.Uint16(dataPacket[22:24])
	gpsData.magY = binary.LittleEndian.Uint16(dataPacket[24:26])
	gpsData.magZ = binary.LittleEndian.Uint16(dataPacket[26:28])
	//*/

	gpsData.Flags = binary.LittleEndian.Uint16(dataPacket[28:30])

	return
}
