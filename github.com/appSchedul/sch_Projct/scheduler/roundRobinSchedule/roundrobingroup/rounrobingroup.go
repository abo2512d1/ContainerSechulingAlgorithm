package roundrobingroup

import (
	"github.com/appSchedul/sch_Projct/applicationdis/appCreator/appproperties"
	"github.com/appSchedul/sch_Projct/zone/zoneproprteties"
)

//GetzoneCatCPUAvr function calculate zone average of CPU
func GetzoneCatCPUAvr(zone1 []zoneproprteties.ZoneProperity, zoneNum int) int {
	sum := 0
	for i := 0; i < zoneNum; i++ {
		sum = sum + zone1[i].GetZoneCPU()
	}
	return sum / zoneNum
}

//GetzoneCatRAMAvr function calculate zone average of CPU
func GetzoneCatRAMAvr(zone1 []zoneproprteties.ZoneProperity, zoneNum int) int {
	sum := 0
	for i := 0; i < zoneNum; i++ {
		sum = sum + zone1[i].GetZoneRAM()
	}
	return sum / zoneNum
}

//GroupingApp fuction the grouping the applicatons based on zone and cells numbers and CPU constairn
func GroupingApp(app []appproperties.ApplProperty, zone []zoneproprteties.ZoneProperity, zoneNum int, cellNum int, appNum int) [][]string {
	var groups = make([][]string, zoneNum)
	var appNotingroups = make([]string, zoneNum*cellNum)
	var appCount int
	appCount = 0
	//cellNum = 100
	z := 0
	c := 0
	ce := 0

	for i := 0; i < zoneNum; i++ {
		groups[i] = make([]string, cellNum)

	}
	for i := 0; i < zoneNum; i++ {
		for j := 0; j < cellNum; j++ {
			z = 0
			c = 0
			ce = 0
			for z < zoneNum && c < cellNum && appNum-appCount > 0 {

				if app[appCount].GetAppCPU() <= zone[z].GetZoneCPU() && app[appCount].GetAppRAM() <= zone[z].GetZoneRAM() {
					if groups[z][c] == "" {
						groups[z][c] = app[appCount].GetAppAname()
						if appNum-appCount != 0 {
							appCount++
						}
						break
					} else {
						if c >= cellNum-1 {
							if z == zoneNum {
								ce = 0
								for ce < len(appNotingroups) {
									if appNotingroups[ce] == "" {
										appNotingroups[ce] = app[appCount].GetAppAname()

										if appNum-appCount != 0 {
											appCount++
										}

										break
									} else {
										ce++
									}
								}
							} else {
								z++
								c = 0

							}
						} else {
							c++

						}
					}

				} else {
					if z == zoneNum-1 {
						ce = 0
						for ce < len(appNotingroups) {
							if appNotingroups[ce] == "" {
								appNotingroups[ce] = app[appCount].GetAppAname()
								if appNum-appCount != 0 {
									appCount++
								}
								break
							} else {
								ce++
							}
						}
					} else {
						z++
						c = 0
					}
				}
			}

		}

	}
	//}
	return groups
}
