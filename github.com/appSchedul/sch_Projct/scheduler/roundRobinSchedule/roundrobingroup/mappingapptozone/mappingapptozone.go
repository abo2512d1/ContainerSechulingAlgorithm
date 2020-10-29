package mappingapptozone

import (
	"fmt"
	"strconv"
)

//MappingApptoZone fuction that maaping the application to zone after roud-robin finish
func MappingApptoZone(gropedApp [][]string, zoneNum int, cellNum int, appNum int) map[string]string {
	mapping := make(map[string]string, appNum)
	mapping[""] = ""
	countApp := 0
	for i := 0; i < zoneNum; i++ {
		for j := 0; j < cellNum; j++ {
			mapping[gropedApp[i][j]] = "zone" + strconv.Itoa(i)
			countApp++
		}
	}
	return mapping
}

//PrintingZoneLoad print zone mapping
func PrintingZoneLoad(groupApp [][]string, mapping map[string]string, zoneNum int, cellNum int) {

	//fmt.Println(mapping)
	for i := 0; i < zoneNum; i++ {
		fmt.Print("zone", i, "   |")
		for j := 0; j < cellNum; j++ {
			if groupApp[i][j] != "" {
				fmt.Print("+")
			} else {
				fmt.Print("..")
			}

		}
		fmt.Println()
	}
}
