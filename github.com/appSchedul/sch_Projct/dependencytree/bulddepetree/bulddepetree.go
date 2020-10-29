package bulddepetree

import (
	"fmt"
	"strconv"

	"github.com/appSchedul/sch_Projct/dependencytree/treeclass"
)

//GetZoneNUmber function take map of application as key and zone number as value and returne zone number
func GetZoneNUmber(appName string, appMapToZone map[string]string) int {
	zoneName := appMapToZone[appName]
	var zomeNum int
	zomeNum1 := ""

	for i := range zoneName {
		if string(zoneName[i]) != "z" && string(zoneName[i]) != "o" && string(zoneName[i]) != "n" && string(zoneName[i]) != "e" && string(zoneName[i]) != "" {

			zomeNum1 = string(zoneName[i])

		}

	}
	if zomeNum1 == "" {
		zomeNum = 999
	} else {
		zomeNum, _ = strconv.Atoi(zomeNum1)
	}
	return zomeNum
}

//GetArryofDepndentApp function that buld array of dependent application
func GetArryofDepndentApp(depeMat [][]int, appMapToZone map[string]string, appNum int, depndNum int) []treeclass.DepTree {
	fmt.Println(depndNum, appNum)
	if depndNum < appNum {
		depndNum = appNum
	}
	adpeApp := make([]treeclass.DepTree, depndNum+100)

	appCount := 0
	appANameRcon := ""
	appANamelcon := ""
	var zomeNmF, zomeNmC int
	for i := 0; i < appNum; i++ {
		appANamelcon = "A" + strconv.Itoa(i)

		for j := 0; j < appNum; j++ {
			appANameRcon = "A" + strconv.Itoa(j)
			if depeMat[i][j] == 1 {
				if appANameRcon != appANamelcon {
					zomeNmF = GetZoneNUmber(appANamelcon, appMapToZone)
					zomeNmC = GetZoneNUmber(appANameRcon, appMapToZone)
					adpeApp[appCount].SetappFnamFun(appANamelcon, zomeNmF)
					adpeApp[appCount].SetappCnamFun(appANameRcon, zomeNmC)
					appCount++
				}
			}
		}
	}
	return adpeApp
}

//GetArryofDepndentAppNewAppArive
func GetArryofDepndentAppNewAppArive(newAppDepne [][]int, maxCutmapping map[string]string, newAppNum int, oldAppNum int, newdepNum int, oldDepTree []treeclass.DepTree) []treeclass.DepTree {
	if newdepNum < newAppNum {
		newdepNum = newAppNum
	}
	fmt.Println("here")
	adpeApp := make([]treeclass.DepTree, newdepNum+1000)
	nameF := ""
	for i := 0; i < len(oldDepTree); i++ {
		nameF, _ = oldDepTree[i].PrintCname()
		if nameF != "" {
			adpeApp[i] = oldDepTree[i]
		}
	}
	appCount := 0
	appANameRcon := ""
	appANamelcon := ""
	var zomeNmF, zomeNmC int
	for i := oldAppNum; i < newAppNum; i++ {
		appANamelcon = "A" + strconv.Itoa(i)

		for j := 0; j < newAppNum; j++ {
			appANameRcon = "A" + strconv.Itoa(j)
			if newAppDepne[i][j] == 1 {
				if appANameRcon != appANamelcon {
					for k := 0; k < len(adpeApp); k++ {
						nameF, _ = adpeApp[k].PrintCname()
						if nameF == "" {
							zomeNmF = GetZoneNUmber(appANamelcon, maxCutmapping)
							zomeNmC = GetZoneNUmber(appANameRcon, maxCutmapping)
							adpeApp[appCount].SetappFnamFun(appANamelcon, zomeNmF)
							adpeApp[appCount].SetappCnamFun(appANameRcon, zomeNmC)
							appCount++
							break
						}
					}
				}
			}
		}
	}
	return adpeApp

}
