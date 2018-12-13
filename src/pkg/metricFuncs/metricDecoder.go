package metricFuncs

import (
	"fmt"
)

//"fmt"
//"encoding/binary"
//"encoding/hex"
//"bytes"

func metricDecoder(metricData *MetricStruct, gpsData *GPSStruct, decodedData *DecodedStruct) {

	// Declare Constant Values
	var DistConv float32 = 0.0711111
	var SpeedConv float32 = 0.00277778

	//*/ Session Distance
	tempDist := metricData.sessDist
	var floatDist float32 = float32(tempDist) * DistConv
	decodedData.TotalDistance = Truncate(floatDist)
	//*/

	/*/ Session Distances Z5&6
	tempDist = metricData.sessDistZ5
	floatDist = float32(tempDist) * DistConv
	decodedData.TotalDistanceZ5 = Truncate(floatDist)

	tempDist = metricData.sessDistZ6
	floatDist = float32(tempDist) * DistConv
	decodedData.TotalDistanceZ6 = Truncate(floatDist)
	//*/

	//*/ Max Speed 10s
	tempMaxSpeed := metricData.maxSpeed10s
	var floatMaxSpeed float32 = float32(tempMaxSpeed) * SpeedConv
	decodedData.MaxSpeed = Truncate(floatMaxSpeed)
	//*/

	/*/ Speed Intensity
	decodedData.SpeedIntens = metricData.speedIntens
	//*/

	//*/ Accels and Decels
	decodedData.TotalAccel = metricData.totalNrAccel
	decodedData.TotalDecel = metricData.totalNrDecel
	//*/

	/*/ Step Balance
	decodedData.StepBalLeft = Truncate(40 + ((float32(metricData.stepBal) * 20) / 255))
	decodedData.StepBalRight = 100 - decodedData.StepBalLeft
	//*/

	//*/  Impacts and DSL
	decodedData.Impacts = metricData.impacts
	decodedData.Dsl = metricData.dsl
	//*/

	/*/ Average Speed
	tempAvSpeed := metricData.speed
	var floatAvSpeed float32 = float32(tempAvSpeed) * SpeedConv
	decodedData.AverageSpeed = Truncate(floatAvSpeed)
	//*/

	//*/ Current Speed
	tempCurrentSpeed := gpsData.codedSpeed
	var floatCurSpeed float32 = float32(tempCurrentSpeed) * SpeedConv
	decodedData.CurrentSpeed = Truncate(floatCurSpeed)
	//*/

	fmt.Println(gpsData.codedLong)
	fmt.Println(gpsData.codedLat)
	//*/ GPS Lat and Long
	var tempLatInt int32
	var LatOrientation string

	var tempLongInt int32
	var LongOrientation string

	// Lattitude Decode
	tempLatInt = gpsData.codedLat
	//tempLatDouble := float64(tempLatInt) / 100000

	LatOrientation = "N"
	if tempLatInt < 0 {
		tempLatInt *= -1
		LatOrientation = "S"
	}

	LatIntDeg := tempLatInt / 10000000
	LatIntMin := (tempLatInt - (LatIntDeg * 10000000)) / 100000
	LatIntSec := (tempLatInt - (LatIntMin * 100000) - (LatIntDeg * 10000000))

	fmt.Println("tempLatInt - ", tempLatInt)
	fmt.Println("LatIntDeg - ", LatIntDeg)
	fmt.Println("LatIntMin - ", LatIntMin)
	fmt.Println("LatIntSec - ", LatIntSec)
	fmt.Println()

	//LatIntMinFrac := (LatIntDeg * 6000000) + LatIntMin

	LatDecDeg := (LatIntDeg * 10000000) + (LatIntMin * 10000000 / 60) + (LatIntSec * 10000000 / (3600 * 1000))

	if LatOrientation != "N" {
		LatDecDeg *= -1
	}

	decodedData.decodedLat = LatDecDeg

	// Longitude Decode
	tempLongInt = gpsData.codedLong
	//tempLongDouble := float64(tempLongInt) / 100000

	LongOrientation = "E"
	if tempLongInt < 0 {
		tempLongInt *= -1
		LongOrientation = "W"
	}

	LongIntDeg := tempLongInt / 10000000
	LongIntMin := (tempLongInt - (LongIntDeg * 10000000)) / 100000
	LongIntSec := (tempLongInt - (LongIntMin * 100000) - (LongIntDeg * 10000000))

	fmt.Println("tempLongInt - ", tempLongInt)
	fmt.Println("LongIntDeg - ", LongIntDeg)
	fmt.Println("LongIntMin - ", LongIntMin)
	fmt.Println("LongIntSec - ", LongIntSec)
	fmt.Println()

	//LongIntMinFrac := (LongIntDeg * 6000000) + LongIntMin

	LongDecDeg := (LongIntDeg * 10000000) + (LongIntMin * 10000000 / 60) + (LongIntSec * 10000000 / (3600 * 1000))

	if LongOrientation != "N" {
		LongDecDeg *= -1
	}

	decodedData.decodedLong = LongDecDeg

	fmt.Print(LatDecDeg)
	fmt.Println(",", LongDecDeg)
	fmt.Println()
	//*/

	/*/
	int32_t *p_i32 = (int32_t*)&b2[gps_base + 8];
	int32_t lat_i = *p_i32;
	char lat_c = 'N';
	if (lat_i < 0) {
	lat_i *= -1;
	lat_c = 'S';
	}
	double lat_d = lat_i;

	++p_i32;
	int32_t lon_i = *p_i32;
	char lon_c = 'E';
	if (lon_i < 0) {
	lon_i *= -1;
	lon_c = 'W';
	}
	double lon_d = lon_i;
	lat_d /= 100000;
	lon_d /= 100000;
	//*/
	/*/
	// lon_i and lat_i have a format like DDMMmmmmm
	// lon_i_m and lat_i_m should normally have kept the original sign
	// but we removed it in lon_i and lat_i so we have to add back

	int32_t lon_i_d = lon_i / 10000000;
	int32_t lon_i_m = lon_i - lon_i_d * 10000000;
	int32_t lat_i_d = lat_i / 10000000;
	int32_t lat_i_m = lat_i - lat_i_d * 10000000;

	// _mf is minutes_fraction format and is in 1/10000  minutes
	// !!! actually NOW is in 1/100000 minutes for max internal resolution
	//lon_i_mf = lon_i_d * 60 * 10000 + lon_i_m / 10;

	lon_i_mf = lon_i_d * 60 * 100000 + lon_i_m ;

	//lat_i_mf = lat_i_d * 60 * 10000 + lat_i_m / 10;

	lat_i_mf = lat_i_d * 60 * 100000 + lat_i_m ;

	d_lon = lon_i_m;
	d_lon /= 6000000;
	d_lon += lon_i_d;
	if(lon_c != 'E') {
	d_lon *= -1;
	lon_i_mf = -lon_i_mf;
	}
	d_lat = lat_i_m;
	d_lat /= 6000000;
	d_lat += lat_i_d;
	if(lat_c != 'N') {
	d_lat *= -1;
	lat_i_mf = -lat_i_mf;
	}

	//*/

	return
}
