package metricFuncs

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

	//*/ GPS Lat and Long
	var LatIntDeg int32
	var tempLatFloat float32
	var LatOrientation string

	var LongIntDeg int32
	var tempLongFloat float32
	var LongOrientation string

	//fmt.Println(gpsData.codedLong)
	//fmt.Println(gpsData.codedLat)

	// Latitude Decode
	LatIntDeg = gpsData.codedLat / 10000000
	tempLatFloat = float32(gpsData.codedLat) / 10000000

	LatOrientation = "N"
	if LatIntDeg < 0 {
		LatIntDeg *= -1
		tempLatFloat *= -1
		LatOrientation = "S"
	}

	LatIntMin := (tempLatFloat - float32(LatIntDeg)) * 100

	/*/
	fmt.Println("LatIntDeg - ", LatIntDeg)
	fmt.Println("LatIntMin - ", LatIntMin)
	fmt.Println()
	//*/

	LatDecDeg := float32(LatIntDeg) + (LatIntMin / 60)

	if LatOrientation != "N" {
		LatDecDeg *= -1
	}

	decodedData.decodedLat = LatDecDeg

	// Longitude Decode
	LongIntDeg = gpsData.codedLong / 10000000
	tempLongFloat = float32(gpsData.codedLong) / 10000000

	LatOrientation = "N"
	if LongIntDeg < 0 {
		LongIntDeg *= -1
		tempLongFloat *= -1
		LongOrientation = "S"
	}

	LongIntMin := (tempLongFloat - float32(LongIntDeg)) * 100

	/*/
	fmt.Println("tempLongFloat - ", tempLongFloat)
	fmt.Println("LongIntDeg - ", LongIntDeg)
	fmt.Println("LongIntMin - ", LongIntMin)
	fmt.Println()
	//*/

	LongDecDeg := float32(LongIntDeg) + (LongIntMin / 60)

	if LongOrientation != "N" {
		LongDecDeg *= -1
	}

	decodedData.decodedLong = LongDecDeg

	/*/
	fmt.Print(LatDecDeg)
	fmt.Println(",", LongDecDeg)
	fmt.Println()
	//*/
	//*/

	return
}
