package firstapproach

import (
	"strconv"

	"github.com/appSchedul/sch_Projct/applicationdis/appCreator/appproperties"
	"github.com/appSchedul/sch_Projct/dependencytree/bulddepetree"
	"github.com/appSchedul/sch_Projct/maxcutop/maxcut"
	"github.com/appSchedul/sch_Projct/zone/zoneproprteties"
)

// AddAppToExZone add new arrival applications to the existed groping applications into existed zones group by round robin algorithm
func AddAppToExZone(app []appproperties.ApplProperty, zone []zoneproprteties.ZoneProperity, oldgroping [][]string, zoneNum int, cellNum int, appNum int) [][]string {
	var newgrouping = make([][]string, zoneNum)
	//var newgrouping1 []string
	//newgrouping1[0] = "Assss"
	//newgrouping1[1] = "fff"
	var appName string
	for i := 0; i < zoneNum; i++ {
		newgrouping[i] = make([]string, appNum/zoneNum)
		for j := 0; j < appNum/zoneNum; j++ {
			newgrouping[i][j] = oldgroping[i][j]
		}
	}

	for appIndex := 0; appIndex < len(app); appIndex++ {
		appName = app[appIndex].GetAppAname()
		for i := 0; i < zoneNum; i++ {

			if maxcut.CompairApplictionWithZon(app, zone, appName, "zone"+strconv.Itoa(i)) {

				for j := 0; j < appNum/zoneNum; j++ {
					if newgrouping[i][j] == "" {
						newgrouping[i][j] = appName
						i = zoneNum
						break
					}
				}

			}
		}
	}
	return newgrouping
}

//CreateGroupFroLastStageOfinhanceLB creat group array from last mappint
func CreateGroupFroLastStageOfinhanceLB(mapping map[string]string, zoneNum int, appNum int) [][]string {
	appNname := ""
	zoneNum1 := 0
	newGrouping := make([][]string, zoneNum)
	for i := 0; i < zoneNum; i++ {
		newGrouping[i] = make([]string, appNum)
	}
	for i := 0; i < appNum; i++ {
		appNname = "A" + strconv.Itoa(i)
		zoneNum1 = bulddepetree.GetZoneNUmber(appNname, mapping)
		for j := 0; j < appNum; j++ {
			if newGrouping[zoneNum1][j] == "" {
				newGrouping[zoneNum1][j] = appNname
				break
			}
		}

	}
	return newGrouping
}
