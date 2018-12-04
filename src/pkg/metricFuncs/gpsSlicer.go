package metricFuncs

import (
	"encoding/binary"
	"encoding/hex"
)

func gpsSlicer(dataPacket []byte, gpsData *GPSStruct) (devID string) {

	//*// DevID - Uint16 or String Formats (pick one)
	//gpsData.devID = binary.LittleEndian.Uint16(dataPacket[0:2])
	gpsData.devID = hex.EncodeToString(dataPacket[0:2])
	//*/

	//*// GPS Time & Date
	gpsData.gpsTime = binary.BigEndian.Uint32(dataPacket[2:6])
	gpsData.gpsDate = binary.BigEndian.Uint32(dataPacket[6:10])
	//*/

	/*// Coded Latitude, Longitude & Altitude
	gpsData.codedLatDeg = binary.LittleEndian.Uint32(dataPacket[10:14])
	gpsData.codedLatMin =
	gpsData.codedLongDeg = binary.LittleEndian.Uint32(dataPacket[14:18])
	gpsData.codedLongMin =
	gpsData.codedAlt = binary.LittleEndian.Uint16(dataPacket[18:20])
	//*/

	//*// Coded Speed
	gpsData.codedSpeed = binary.LittleEndian.Uint16(dataPacket[20:22])
	//*/

	/*// Magnetometer Metrics
	gpsData.magX = binary.LittleEndian.Uint16(dataPacket[22:24])
	gpsData.magY = binary.LittleEndian.Uint16(dataPacket[24:26])
	gpsData.magZ = binary.LittleEndian.Uint16(dataPacket[26:28])
	//*/

	gpsData.Flags = binary.LittleEndian.Uint16(dataPacket[28:30])

	return gpsData.devID
}
