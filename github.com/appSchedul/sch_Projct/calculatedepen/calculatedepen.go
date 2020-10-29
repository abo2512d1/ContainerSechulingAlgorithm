package calculatedepen

import (
	"fmt"
	"strconv"

	"github.com/appSchedul/sch_Projct/dependencytree/treeclass"

	"github.com/appSchedul/sch_Projct/dependencytree/bulddepetree"
)

//CountLevelthreeEdges to count the edges of path from zone to zone return two number count of level 3 edges and count of level 2 edges respectivly
func CountLevelthreeEdges(zoneNumR int, zoneNumC int) (int, int) {
	return 2, 4
}

//CountLevelTwoandOneEdges to count the edges of path in the same  zone  return  number count of level 2 nad 1 edges
func CountLevelTwoandOneEdges(zoneNumR int, zoneNumC int, appNamC string, appNamR string, appMapToZone map[string]string) int {
	x := 0
	i := 0
	j := 0
	for k := range appMapToZone {
		i++
		if k == appNamC {
			break
		}
	}
	for k := range appMapToZone {
		j++
		if k == appNamR {
			break
		}
	}
	if i == j+1 || i+1 == j {
		x = 2
	} else {
		x = 4
	}
	return x
}

//CountLevelTwoandOneEdges1 to count the edges of path in the same  zone  return  number count of level 2 nad 1 edges for max cut
func CountLevelTwoandOneEdges1(zoneNumR int, zoneNumC int) int {
	return 4
}

//CalculateDepFunc is function that calculte the dependencey between application
func CalculateDepFunc(depeMat [][]int, appMapToZone map[string]string, trafCost [][]int, appNum int) (int, int, int) {
	//zoneName := ""
	appNamC := ""
	appNamR := ""
	zonrNumC := 0
	zonrNumR := 0
	countDeBetwZone := 0
	countDeInZone := 0
	countDeLevelThree := 0
	countDeLevelTwoandOne := 0
	deBz := 0
	deIz := 0
	//movingAppCost := 0
	calDepen := 0

	for i := 0; i < appNum; i++ {
		for j := 0; j < appNum; j++ {
			if depeMat[i][j] > 0 {
				//zoneName = "zone" + strconv.Itoa(i)
				appNamR = "A" + strconv.Itoa(i)
				appNamC = "A" + strconv.Itoa(j)
				if appNamR != appNamC {
					if appMapToZone[appNamC] != "" && appMapToZone[appNamR] != "" {
						zonrNumC = bulddepetree.GetZoneNUmber(appNamC, appMapToZone)
						zonrNumR = bulddepetree.GetZoneNUmber(appNamR, appMapToZone)
						if appMapToZone[appNamC] != appMapToZone[appNamR] {
							//fmt.Println("appNameR=", appNamR, appMapToZone[appNamR], "==>", "zome", zonrNumR, appMapToZone[appNamC], "   ", "appNameC=", appNamC, "==>", "zome", zonrNumC)
							deBz++
							//calDepen = calDepen + trafCost[0][1]
							countDeLevelThree, countDeLevelTwoandOne = CountLevelthreeEdges(zonrNumC, zonrNumR)
							countDeBetwZone = countDeBetwZone + countDeLevelThree
							countDeInZone = countDeInZone + countDeLevelTwoandOne
						} else {
							deIz++
							//calDepen = calDepen + trafCost[0][0]
							countDeLevelTwoandOne = CountLevelTwoandOneEdges(zonrNumC, zonrNumR, appNamR, appNamC, appMapToZone)
							countDeInZone = countDeInZone + countDeLevelTwoandOne

						}
					}
				}
			}

		}
	}
	//fmt.Println("count dep be zome=", deBz, "===", countDeBetwZone, "count dep In zome=", deIz, "===", countDeInZone, "total dep numb=", deBz+deIz)
	calDepen = (countDeInZone * trafCost[0][0]) + (countDeBetwZone * trafCost[0][1])
	return calDepen, countDeBetwZone, countDeInZone
}

//CaculateProprotionOFtraffic calcalute the proprotion of traffic
func CaculateProprotionOFtraffic(trafficCost int, movingAppCost int, movingAppNum int, independ int, outdepend int) float64 {
	fmt.Println("moving application number=", movingAppNum)
	fmt.Println("bast=", trafficCost+movingAppCost, "Maqam=", float64((independ*1000+outdepend*1000)+(2*movingAppNum*1000+4*movingAppNum*1000)))
	return (100 * float64(trafficCost+movingAppCost) / float64((independ*1000+outdepend*1000)+(2*movingAppNum*1000+4*movingAppNum*1000)))
}

//ColudPerformance caculate cloud performance
func ColudPerformance(loadBalance float64, traffic float64) float64 {
	fmt.Println("CV=", loadBalance, "traffic=", traffic)
	return (loadBalance + traffic) / 2
}

//ColudPerformance1 caculate cloud performance
func ColudPerformance1(loadBalance float64, traffic float64) float64 {
	fmt.Println("CV=", loadBalance, "traffic=", traffic)
	return (1 / (loadBalance + traffic))
}

//CalculateDeNumFromClass
func CalculateDeNumFromClass(treeDep []treeclass.DepTree) {
	count := 0
	var zoneF, zoneC int
	for i := 0; i < len(treeDep); i++ {
		_, zoneF = treeDep[i].PrintFname()
		_, zoneC = treeDep[i].PrintCname()
		if zoneF != zoneC {
			count++
		}

	}
	fmt.Println("count number of dep from class ", count)
}
