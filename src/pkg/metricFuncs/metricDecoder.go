package metricFuncs

//"fmt"
//"encoding/binary"
//"encoding/hex"
//"bytes"

func Truncate(some float32) float32 {
	return float32(int(some*100)) / 100
}

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

	/*/ GPS Lat and Long
	decodedData.decodedLat =
	decodedData.decodedLong =
	//*/

	return
}
