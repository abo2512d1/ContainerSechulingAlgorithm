package maxcut

import (
	"fmt"
	"strconv"

	"github.com/appSchedul/sch_Projct/applicationdis/appCreator/appproperties"
	"github.com/appSchedul/sch_Projct/calculatedepen"
	"github.com/appSchedul/sch_Projct/static"
	"github.com/appSchedul/sch_Projct/zone/zoneproprteties"

	"github.com/appSchedul/sch_Projct/dependencytree/treeclass"
)

//CalcuDepdFromTreeClass calacuate the dependencies from tree class array in-dependncdey, out-dependency
func CalcuDepdFromTreeClass(depnTree []treeclass.DepTree) (int, int) {
	var inDepnd, outDepnd, fDepnd, cDpend int
	inDepnd = 0
	outDepnd = 0
	fname := ""
	cname := ""

	countDepLevelThree := 0
	countDepLevelTwoandone := 0
	for i := 0; i < len(depnTree); i++ {
		fname, fDepnd = depnTree[i].PrintFname()
		cname, cDpend = depnTree[i].PrintCname()
		if fname != "" && cname != "" {
			if fDepnd == cDpend {
				countDepLevelTwoandone = calculatedepen.CountLevelTwoandOneEdges1(fDepnd, cDpend)
				inDepnd = inDepnd + countDepLevelTwoandone
			} else {
				countDepLevelThree, countDepLevelTwoandone = calculatedepen.CountLevelthreeEdges(fDepnd, cDpend)
				outDepnd = outDepnd + countDepLevelThree
				inDepnd = inDepnd + countDepLevelTwoandone

			}
		}
	}

	//fmt.Println("out dep cut from terr calss=", countMaxCt)
	return inDepnd, outDepnd
}

//MaxCut function that find neigbours for the pplication
func MaxCut(depnTree []treeclass.DepTree, appCon []appproperties.ApplProperty, zoneCon []zoneproprteties.ZoneProperity, mapping map[string]string, depndNum int, zoneNum int, appNum int) ([]treeclass.DepTree, int, int, int, int, int, int) {
	neighFName := ""
	nieghFZoneNum := 0
	neighCName := ""
	nieghCZoneNum := 0
	zoneNmaeForText := ""
	var countMovedApp, movingAppCost int
	countMovedApp = 0

	//fmt.Println("start")
	//------------------------------ save zone state
	neighFNameState := ""
	nieghFZoneNumState := 0
	neighCNameState := ""
	nieghCZoneNumState := 0
	//countAppInZone := make(map[string]int)
	fmt.Println(neighFNameState,
		nieghFZoneNumState)
	//------------------------ calcuate dependency
	var inDependBeforMovind, outDependBeforMovind, inDependAfterMovind, outDependAfterMovind int
	inDependBeforMovind, outDependBeforMovind = CalcuDepdFromTreeClass(depnTree)
	//---------------------appF class Info
	appFname := ""
	//appFZoneNum := 0
	//-----------------------appC class Info
	appCname := ""
	//appCZoneNum := 0

	for k := 0; k < len(depnTree); k++ {
		neighFName, nieghFZoneNum = depnTree[k].PrintFname()
		neighCName, nieghCZoneNum = depnTree[k].PrintCname()
		nubOfDep := 0.0
		if nieghFZoneNum != nieghCZoneNum {
			neighFNameState = neighFName       //save state for roolback
			nieghFZoneNumState = nieghFZoneNum //save state for roolback
			neighCNameState = neighCName       //save state for roolback
			nieghCZoneNumState = nieghCZoneNum //save state for roolback
			//-----------------------------------move Cnieghbour to Fnieghbour zone
			depnTree[k].SetappCnamFun(neighCName, nieghFZoneNum) //change the zone number for appcC in appF class ---need roolbak if not cut
			countMovedApp++

			//fmt.Println(countMovedApp, "here")
			//-----------------------------------change the appC Info
			appCname, _ = depnTree[k].PrintCname()
			for i := 0; i < len(depnTree); i++ { //change one number in AppC in orginal appF ---need roolbak if not cut
				appFname, _ = depnTree[i].PrintFname()
				//	countAppInZone = static.CountAppInEachZone(mapping, zoneNum)
				if appFname == appCname { //&& countAppInZone["zone"+strconv.Itoa(appFZoneNum)] < appNum/zoneNum {
					//fmt.Println("moving app:", appCname, "to zone :", nieghFZoneNum)
					//fmt.Println("here1 move", countMovedApp)
					depnTree[i].SetappFnamFun(appCname, nieghFZoneNum)
					depnTree[i].SetMovingStituos(true)
				}
			}
			inDependAfterMovind, outDependAfterMovind = CalcuDepdFromTreeClass(depnTree) //calsulate Dependency after change appc zone number
			zoneNmaeForText = "zone" + strconv.Itoa(nieghCZoneNum)
			if outDependAfterMovind > outDependBeforMovind && CompairApplictionWithZon(appCon, zoneCon, neighFName, zoneNmaeForText) == false { // if ture must roolback
				depnTree[k].SetappCnamFun(neighCNameState, nieghCZoneNumState)
				fmt.Println("return app:", neighCNameState)
				countMovedApp = countMovedApp - 1
				//fmt.Println("here2 return", countMovedApp)
				for i := 0; i < len(depnTree); i++ {
					appFname, _ = depnTree[i].PrintFname()
					if appFname == appCname {

						depnTree[i].SetappFnamFun(neighCNameState, nieghCZoneNumState)
						depnTree[k].SetMovingStituos(false)
					}
				}

			}
			//UpdteMappingApp(mapping, depnTree, depndNum)

			movingAppCost = ((countMovedApp * 4 * 10) + (countMovedApp * 2 * 100)) * 10
			appnamem := ""
			for i := 0; i < len(depnTree); i++ {
				appnamem, _ = depnTree[i].PrintFname()
				if appnamem != "" {
					nubOfDep++
				}
			}
			if len(depnTree)%2 != 0 {
				nubOfDep++
			}
			inDependAfterMovind, outDependAfterMovind = CalcuDepdFromTreeClass(depnTree)
			if countMovedApp == int(nubOfDep*.2) || countMovedApp == int(nubOfDep*.4) || countMovedApp == int(nubOfDep*.6) || countMovedApp == int(nubOfDep*.8) {
				fmt.Println("out dep cut=", outDependAfterMovind/6)
				CutMaxProgress(inDependAfterMovind, outDependAfterMovind, countMovedApp, movingAppCost, depnTree, mapping, depndNum, zoneNum, appNum)
			}
		}
	}
	//maxCutPromapping := UpdteMappingApp(mapping, depnTree, depndNum)
	return depnTree, inDependBeforMovind, outDependBeforMovind, inDependAfterMovind, outDependAfterMovind, countMovedApp, movingAppCost
}

//UpdteMappingApp this update mapping appiction after Max_cut
func UpdteMappingApp(mappin map[string]string, depnTree []treeclass.DepTree, dpendNum int) map[string]string {
	for i := 0; i < dpendNum; i++ {
		appName, newZone := depnTree[i].PrintFname()

		if newZone > 0 {

			appZonename := "zone" + strconv.Itoa(newZone)

			mappin[appName] = appZonename
		}

	}
	return mappin
}

//UpdteMappingApp1 this update mapping appiction after Max_cut
//func UpdteMappingApp1(mappin map[string]string, depnTree []treeclass.DepTree, dpendNum int) map[string]string {
//	var appName, appName1, appNameC, appZonename, appZonenameC string //,
//	var newZone, newZone1, newZoneC int
//b := false

//	for i := 0; i < len(depnTree); i++ {
//	appName, newZone = depnTree[i].PrintFname()
//
//	if appName != "" {

//			appZonename := "zone" + strconv.Itoa(newZone)

//		mappin[appName] = appZonename
//		}
//	}
//	for i := 0; i < len(depnTree); i++ {
//		appNameC, newZoneC = depnTree[i].PrintCname()
//	for j := 0; j < len(depnTree); j++ {
//		appName1, newZone1 = depnTree[i].PrintFname()
//		if appNameC != "" && appNameC == appName1 && newZone1 == newZoneC {
//			appZonenameC := "zone" + strconv.Itoa(newZoneC)
//			mappin[appNameC] = appZonenameC
//		} else {
//			appZonenameC := "zone" + strconv.Itoa(newZoneC)
//			mappin[appNameC] = appZonenameC
//		}
//	}
//	}

//	fmt.Println(appZonename, appZonenameC)
//	return mappin
//}

//ShowZoneAferMacCut function shows the balance of zines after Mac-Cut
func ShowZoneAferMacCut(newmappingAfterMacCut map[string]string, appNum int, zoneNum int) {
	countZone := make([]int, zoneNum)
	zoneNme := ""
	for i := 0; i < len(countZone); i++ {
		zoneNme = "zone" + strconv.Itoa(i)
		for _, v := range newmappingAfterMacCut {
			if v == zoneNme {
				countZone[i]++
			}
		}
	}
	for i := 0; i < len(countZone); i++ {
		fmt.Print("zone", i, "  |")
		for j := 0; j < countZone[i]; j++ {
			fmt.Print("+")
		}
		fmt.Println()
	}

}

//CompairApplictionWithZon function comapier the appliction requirments with Zone capacity
func CompairApplictionWithZon(appCon []appproperties.ApplProperty, zoneCon []zoneproprteties.ZoneProperity, appName string, zoneName string) bool {
	var zoneCPU, zoneRAM int
	var b bool
	for i := 0; i < len(zoneCon); i++ {
		if zoneCon[i].GetZoneName() == zoneName {
			zoneCPU = zoneCon[i].GetZoneCPU()
			zoneRAM = zoneCon[i].GetZoneRAM()

		}
	}
	for j := 0; j < len(appCon); j++ {
		if appCon[j].GetAppAname() == appName {
			if appCon[j].GetAppCPU() <= zoneCPU && appCon[j].GetAppRAM() <= zoneRAM {
				b = true
			} else {
				b = false
			}
		}
	}
	return b
}

//CutMaxProgress this function keep cut the dependencies in 20%,40%,60%,80% and show the resute
func CutMaxProgress(inDependAfterMovind int, outDependAfterMovind int, countMovedApp int, movingAppCost int, depnTree []treeclass.DepTree, mapping map[string]string, depndNum int, zoneNum int, appNum int) {
	fmt.Println("MaX cut Progress\n")
	maxCutPromapping := UpdteMappingApp(mapping, depnTree, depndNum)
	fmt.Println(movingAppCost)
	ShowZoneAferMacCut(mapping, appNum, zoneNum)
	appNuminEacZoneAftermaxCut := static.CountAppInEachZone(maxCutPromapping, zoneNum)
	fmt.Println(appNuminEacZoneAftermaxCut)
	fmt.Println("traffic Cost =", calculatedepen.CaculateProprotionOFtraffic(inDependAfterMovind*10+outDependAfterMovind*100, movingAppCost, countMovedApp, inDependAfterMovind, outDependAfterMovind))

	fmt.Println("Cost of Moving Application", calculatedepen.CaculateProprotionOFtraffic(0, movingAppCost, countMovedApp, 0, 0))
	meanAfterMaxCut := static.CalculateMean(appNum, zoneNum)
	sdAfterMaxCut := static.CalculateSD(appNuminEacZoneAftermaxCut, appNum, zoneNum, meanAfterMaxCut)
	fmt.Println("mean= ", meanAfterMaxCut)
	fmt.Println("SD= ", sdAfterMaxCut)
	fmt.Println("load balane =", static.LoadbalanceIndector(sdAfterMaxCut, meanAfterMaxCut, zoneNum), "\n")
}
