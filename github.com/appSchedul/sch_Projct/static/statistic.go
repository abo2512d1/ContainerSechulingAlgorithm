package static

import (
	"fmt"
	"math"
	"strconv"
)

//LoadbalanceIndector is function that calculate the load  balancing indecator
func LoadbalanceIndector(sd float64, mean float64, zoneNum int) float64 {
	return 100 - (100*(sd/mean))/math.Sqrt(float64(zoneNum))
}

//CalculateMean is function that calculate the mean
func CalculateMean(appNum int, zoneNum int) float64 {
	x := float64(appNum / zoneNum)
	return x
}

//CalculateSD is function that calculate the Stander Deviation
func CalculateSD(appNuminEacZone map[string]int, appNum int, zoneNum int, mean float64) float64 {
	var sum, x float64
	for _, v := range appNuminEacZone {
		x = float64(v) - mean
		sum = sum + math.Pow(x, 2)
	}
	fmt.Println("sum=", sum)
	sd := float64(math.Sqrt(sum / float64(zoneNum)))
	return sd
}

//CountAppInEachZone is function the return map of AppNumInEacZone as value and zone name as key
func CountAppInEachZone(mapping map[string]string, zoneNum int) map[string]int {
	var zoneName string
	AppNumInEachZone := make(map[string]int, zoneNum)
	count := 0
	for i := 0; i < zoneNum; i++ {
		zoneName = "zone" + strconv.Itoa(i)
		count = 0
		for _, v := range mapping {
			if v == zoneName {
				count++
			}
		}

		AppNumInEachZone[zoneName] = count
	}
	return AppNumInEachZone
}
