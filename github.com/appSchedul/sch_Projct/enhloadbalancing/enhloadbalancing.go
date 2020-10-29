package enhloadbalancing

import (
	"fmt"
	"strconv"

	"github.com/appSchedul/sch_Projct/calculatedepen"
	"github.com/appSchedul/sch_Projct/maxcutop/maxcut"
	"github.com/appSchedul/sch_Projct/static"

	"github.com/appSchedul/sch_Projct/applicationdis/appCreator/appproperties"
	"github.com/appSchedul/sch_Projct/dependencytree/treeclass"
	"github.com/appSchedul/sch_Projct/zone/zoneproprteties"
)

//EnhanceLoadBalancingFun last stage of the algorithm and it enahce the load balancing
func EnhanceLoadBalancingFun(depnTree []treeclass.DepTree, appCon []appproperties.ApplProperty, zoneCon []zoneproprteties.ZoneProperity, appdepMax [][]int, mapping map[string]string, zoneNum int, appNum int, depndNum int) map[string]string {
	idealZoneNumberOfApp := int(appNum / zoneNum)
	indepnedApp := make([]string, appNum)
	depnedApp := make([]string, appNum*4)
	count := 0
	count1 := 0
	appNname := ""
	for i := 0; i < appNum; i++ {
		for j := 0; j < appNum; j++ {
			if appdepMax[i][j] == 1 {
				depnedApp[count1] = "A" + strconv.Itoa(i)
				count1++
				depnedApp[count1] = "A" + strconv.Itoa(j)
				count1++
			}

		}
	}
	for i := 0; i < appNum; i++ {
		appNname = "A" + strconv.Itoa(i)
		for j := 0; j < appNum; j++ {
			if appdepMax[i][j] == 0 && j == (appNum-1) {
				if IsnotInDependentApp(depnedApp, appNname, appNum) {
					indepnedApp[count] = "A" + strconv.Itoa(i)
					count++
				}
			}
		}
	}
	fmt.Println("indepnedApp", indepnedApp)
	newMap := RemapIndpendentApp(appCon, zoneCon, depnTree, mapping, indepnedApp, idealZoneNumberOfApp, appNum, zoneNum)
	return newMap
}

//IsnotInDependentApp()
func IsnotInDependentApp(depnedApp []string, appName string, appNum int) bool {
	b := true
	for i := 0; i < appNum; i++ {
		if depnedApp[i] == appName {
			b = false
		}
	}
	return b
}

//RemapIndpendentApp remap the indpendent applications
func RemapIndpendentApp(appCon []appproperties.ApplProperty, zoneCon []zoneproprteties.ZoneProperity, depnTree []treeclass.DepTree, mapping map[string]string, indepnedApp []string, idealZoneNumberOfApp int, appNum int, zoneNum int) map[string]string {
	fmt.Println("mapping after load balancing", mapping)
	fmt.Println(maxcut.CalcuDepdFromTreeClass(depnTree))
	maxcut.ShowZoneAferMacCut(mapping, appNum, zoneNum)
	zoneAppNum := static.CountAppInEachZone(mapping, zoneNum)
	overLoadZones := make(map[string]int, zoneNum)
	numOfappMov := 0
	appZoneName := ""
	percenNu := len(indepnedApp)
	overNumOfApp := 0
	zoneHostIndAppName := ""
	amountOfAppInZines := static.CountAppInEachZone(mapping, zoneNum) //count appication in each zone
	test := false
	underLoazoneIdeNumberOfApp := 0
	underloadZone := 0
	laodBa := 0.0000000001
	if percenNu%2 > 0 {
		percenNu++
	}
	for k, v := range zoneAppNum { // make an map of zone that overload and the overload amount
		if v-idealZoneNumberOfApp > 0 {
			overLoadZones[k] = v - idealZoneNumberOfApp

		}
	}
	for _, v := range overLoadZones {
		overNumOfApp = overNumOfApp + v
	}
	fmt.Println(overNumOfApp)
	for i := 0; i < len(indepnedApp); i++ { // if the independent application in overload zone application in
		if laodBa == 0.0 {
			break
		}
		appZoneName = mapping[indepnedApp[i]]
		if int(float64(percenNu)*.2) == i || int(float64(percenNu)*.6) == i || int(float64(percenNu)*.8) == i && laodBa != 0.0 {
			laodBa = PrinFinalStageStaticAndZoneShow(mapping, appNum, zoneNum, numOfappMov)
		}
		if overNumOfApp > 0 && amountOfAppInZines[appZoneName] > idealZoneNumberOfApp {
			if IsappInOverloadZone(appZoneName, overLoadZones) { //if true will move tha application to unovleoad zone
				for j := 0; j < zoneNum; j++ {

					zoneHostIndAppName = "zone" + strconv.Itoa(j)

					if zoneHostIndAppName != appZoneName { // if the host zone not the orginal zone of the appl

						if !IsappInOverloadZone(zoneHostIndAppName, overLoadZones) { //the new zone not in the list of overload zones
							amountOfAppInZines = static.CountAppInEachZone(mapping, zoneNum) //update the zone number of application
							if amountOfAppInZines[zoneHostIndAppName] < idealZoneNumberOfApp {
								if maxcut.CompairApplictionWithZon(appCon, zoneCon, indepnedApp[i], zoneHostIndAppName) {
									mapping[indepnedApp[i]] = zoneHostIndAppName
									fmt.Println("moving app :", indepnedApp[i], "to zone", zoneHostIndAppName)
									overNumOfApp--
									numOfappMov++
									//laodBa = PrinFinalStageStaticAndZoneShow(mapping, appNum, zoneNum, numOfappMov)
									amountOfAppInZines = static.CountAppInEachZone(mapping, zoneNum)

									break

								}
							}

						}

					}
				}
			}
		}

	}
	fmt.Println("final result")
	laodBa = PrinFinalStageStaticAndZoneShow(mapping, appNum, zoneNum, numOfappMov)
	//if load balance != 100
	amountOfAppInZines = static.CountAppInEachZone(mapping, zoneNum)
	if laodBa > 0.0 {

		for _, v := range amountOfAppInZines {
			if v <= idealZoneNumberOfApp {

				underLoazoneIdeNumberOfApp = underLoazoneIdeNumberOfApp + v
				underloadZone++

			}
			if v < idealZoneNumberOfApp {
				test = true
			}

		}
		underLoazoneIdeNumberOfApp = int(underLoazoneIdeNumberOfApp / underloadZone)
		numOfappMov = 0
		if test {

			for i := 0; i < zoneNum; i++ { // if the independent application in underload zone application
				if amountOfAppInZines["zone"+strconv.Itoa(i)] > underLoazoneIdeNumberOfApp && amountOfAppInZines["zone"+strconv.Itoa(i)] <= idealZoneNumberOfApp { //if the zone has number more then the underidealNumber and less than idea number 18<zoneAppnumber <20
					for k, v := range mapping {
						if v == "zone"+strconv.Itoa(i) && CheckAppOnIndeApp(k, indepnedApp) { //get zone applicatio name that independnet
							for j := 0; j < zoneNum; j++ { //check zone again
								if amountOfAppInZines["zone"+strconv.Itoa(j)] < underLoazoneIdeNumberOfApp { //if zone less then undernew ideal number less the 18
									if maxcut.CompairApplictionWithZon(appCon, zoneCon, k, "zone"+strconv.Itoa(i)) {

										mapping[k] = "zone" + strconv.Itoa(j)
										numOfappMov++
										amountOfAppInZines = static.CountAppInEachZone(mapping, zoneNum)

									}
								}
							}

						}
					}

				}
				if i == 2 || i == 5 || i == 9 {
					PrinFinalStageStaticAndZoneShow(mapping, appNum, zoneNum, numOfappMov)
				}
			}
		}
	}
	fmt.Println("mapping after load balancing", mapping)
	maxcut.ShowZoneAferMacCut(mapping, appNum, zoneNum)
	fmt.Println(maxcut.CalcuDepdFromTreeClass(depnTree))
	PrinFinalStageStaticAndZoneShow(mapping, appNum, zoneNum, numOfappMov)
	return mapping
}

//CheckAppOnIndeApp chech if the appliction is endpndent
func CheckAppOnIndeApp(appNam string, indeApp []string) bool {
	b := false
	for i := 0; i < len(indeApp); i++ {
		if appNam == indeApp[i] {
			b = true
		}
	}
	return b
}

//PrinFinalStageStaticAndZoneShow function print fial stage statsic and shows the zone
func PrinFinalStageStaticAndZoneShow(mapping map[string]string, appNum int, zoneNum int, numOfappMov int) float64 {
	maxcut.ShowZoneAferMacCut(mapping, appNum, zoneNum)
	amountOfAppInZines := static.CountAppInEachZone(mapping, zoneNum)
	fmt.Println(amountOfAppInZines, "aap movied=", numOfappMov)
	fmt.Println("Cost of Moving Application", calculatedepen.CaculateProprotionOFtraffic(0, CalMovingAppCost(numOfappMov), numOfappMov, 0, 0))
	fmt.Println("movAppCost", CalMovingAppCost(numOfappMov))
	mean := static.CalculateMean(appNum, zoneNum)
	sd := static.CalculateSD(amountOfAppInZines, appNum, zoneNum, mean)
	fmt.Println("mean= ", mean)
	fmt.Println("SD= ", sd)
	fmt.Println("load balane =", static.LoadbalanceIndector(sd, mean, zoneNum))
	return sd
}

//IsappInOverloadZone function returen true if application in overload zone
func IsappInOverloadZone(appZoneName string, overLoadZones map[string]int) bool {
	b := false
	for k := range overLoadZones {
		if appZoneName == k {
			b = true
		}
	}
	return b
}

//CalMovingAppCost  calculate moving application while enhance the load balancing
func CalMovingAppCost(countMovedApp int) int {

	return ((countMovedApp * 4 * 10) + (countMovedApp * 2 * 100)) * 10

}
