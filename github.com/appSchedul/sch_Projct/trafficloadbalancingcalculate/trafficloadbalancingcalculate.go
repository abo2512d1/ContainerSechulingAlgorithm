package trafficloadbalancingcalculate

import (
	"fmt"

	"github.com/appSchedul/sch_Projct/calculatedepen"
	"github.com/appSchedul/sch_Projct/dependencytree/treeclass"
	"github.com/appSchedul/sch_Projct/maxcutop/maxcut"
	"github.com/appSchedul/sch_Projct/static"
)

//CalculateTrafficCostandLoadbalancing calculate traffic cost and ,oad balancing
func CalculateTrafficCostandLoadbalancing(treeDep []treeclass.DepTree, newAppDepne [][]int, newAppMapping map[string]string, appTraf [][]int, zoneNum int, newdepNum int, newAppNum int) {
	appNuminEacZone := static.CountAppInEachZone(newAppMapping, zoneNum)
	fmt.Println("Application number in each zones=", appNuminEacZone)

	newAppEdgesInZones, newAppEdgesBeZones := maxcut.CalcuDepdFromTreeClass(treeDep) //calculate the total traffic cost and the number of edges
	//fmt.Println("newAppDepCost=", newAppDepCost)
	//calculatedepen.CalculateDeNumFromClass(treeDep)
	//fmt.Println("newAppEdgesBeZones*100+newAppEdgesInZones*10", newAppEdgesBeZones*100+newAppEdgesInZones*10)
	newAppTrafficCost := calculatedepen.CaculateProprotionOFtraffic((newAppEdgesBeZones)*100+newAppEdgesInZones*10, 0, 0, newAppEdgesInZones, newAppEdgesBeZones) //trafic cost proportion
	fmt.Println("Total Treffic Cost =", newAppTrafficCost, "\nTotal Treffic Cost 10X=", newAppTrafficCost*10, "\nTotal edges Between zones =", newAppEdgesBeZones, "\nTotal edges In the same zones =", newAppEdgesInZones)
	newappNuminEacZone := static.CountAppInEachZone(newAppMapping, zoneNum) //the nummber of appliction in each zoneusing for calculate SD
	mean := static.CalculateMean(newAppNum, zoneNum)
	sd := static.CalculateSD(newappNuminEacZone, newAppNum, zoneNum, mean)
	fmt.Println("mean=", mean, "\nSD= ", sd)
	loadbalance := static.LoadbalanceIndector(sd, mean, zoneNum)
	fmt.Println("load balancing=", loadbalance)
}
