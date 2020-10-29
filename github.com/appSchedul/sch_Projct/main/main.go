package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/appSchedul/sch_Projct/dependencytree/bulddepetree"
	"github.com/appSchedul/sch_Projct/dependencytree/treeclass"
	"github.com/appSchedul/sch_Projct/enhloadbalancing"
	"github.com/appSchedul/sch_Projct/maxcutop/maxcut"
	"github.com/appSchedul/sch_Projct/newapparrive/firstapproach"
	"github.com/appSchedul/sch_Projct/newapparrive/readconfi"
	"github.com/appSchedul/sch_Projct/static"
	"github.com/appSchedul/sch_Projct/trafficloadbalancingcalculate"

	"github.com/appSchedul/sch_Projct/applicationdis/appCreator/appproperties"
	appfillconstraints "github.com/appSchedul/sch_Projct/applicationdis/appCreator/appproperties/appFillConstraints"
	"github.com/appSchedul/sch_Projct/calculatedepen"
	"github.com/appSchedul/sch_Projct/interfacedis/appdepe"
	readappname "github.com/appSchedul/sch_Projct/interfacedis/readAppName"
	readappnum "github.com/appSchedul/sch_Projct/interfacedis/readAppNum"
	"github.com/appSchedul/sch_Projct/interfacedis/trafficcost"
	"github.com/appSchedul/sch_Projct/scheduler/roundRobinSchedule/roundrobingroup"
	"github.com/appSchedul/sch_Projct/scheduler/roundRobinSchedule/roundrobingroup/mappingapptozone"
	"github.com/appSchedul/sch_Projct/zone/zonefillconstraints"
	"github.com/appSchedul/sch_Projct/zone/zoneproprteties"
)

func main() {

	appfile := "/Users/abdullah/go/src/github.com/appSchedul/souceFile/appConfi200.cog"
	appfileConstraint := "/Users/abdullah/go/src/github.com/appSchedul/souceFile/appConst200.cog"
	zonefile := "/Users/abdullah/go/src/github.com/appSchedul/souceFile/zoneConfi200.cog"
	depndfile := "/Users/abdullah/go/src/github.com/appSchedul/souceFile/depndConfi200.cog"

	var ap string        // application name fill the struct
	var apCPU int        // application cpu fill the struct
	var apRAM int        // application ram fill the struct
	var apIO, zIO string // application io fill the struct
	var zoneNum, zCPU, zRAM, dendCost, numDeBetZone, numDeInZone int
	var appFTrname, appCTrname string                         // to save the name of the application father and cilde name
	var appCterzoneNum, appFterzoneNum, movingAppCost int     // to save the zone number of the application father and cilde
	zoneNum = trafficcost.ReadConfiZoneNum(appfile, zonefile) //appdepe.ReadConfiZoneNum(zonefile)                                 // number of zone
	appNum := readappnum.ReadConfiAppNum(appfile)
	m := float64(float64(appNum) / float64(zoneNum)) //call function to read conf file to read applcations number
	cellNum := int(math.Round(float64(m)))
	fmt.Println(appNum, zoneNum, cellNum) // number of cells per zone
	fmt.Println("here2")
	appName := make([]string, appNum)       //array of application name
	appDep := make([][]int, appNum, appNum) //array of dependency

	appTraf := make([][]int, 2)                   //array of dependency cost
	depdTree := make([]treeclass.DepTree, appNum) //array of dependency class

	appName = readappname.ReadConfiAppname(appfile) //call function to read conf file to read applcations name
	//fmt.Println(appName)
	appDep = appdepe.ReadConfiAppDep(appfile, depndfile, 0) //call function to read conf file to read applcations dependncy
	//fmt.Println("app dependency here")
	//fmt.Println(appDep)
	depndNum := appdepe.ReadDependNum(appDep, appNum)
	fmt.Print("dependen Num")
	fmt.Println(depndNum)

	appTraf = trafficcost.ReadConfiTrafficCost(depndfile) //call function to read conf file to read applcations traffic cost
	//fmt.Println(appDep, appTraf)
	var applic = make([]appproperties.ApplProperty, appNum)   // create array of applications empty properteties
	var zone = make([]zoneproprteties.ZoneProperity, zoneNum) // create array of zone empty properteties
	var groupApp = make([][]string, zoneNum, cellNum)         //create two dimintion of empty group to singe appliction to it for round-robing schedule
	mapping := make(map[string]string, appNum)
	maxCutmapping := make(map[string]string, appNum)
	for i := 0; i < appNum; i++ { //fill application from config file

		ap = "A" + strconv.Itoa(i)
		apCPU = appfillconstraints.ReadAppConstraintsCPU(appfileConstraint, ap)
		apRAM = appfillconstraints.ReadAppConstraintsRAM(appfileConstraint, ap)
		apIO = appfillconstraints.ReadAppConstraintsIO(appfileConstraint, ap)
		//fmt.Println(ap)
		applic[i].ApplPropertyFu(appName[i], ap, 1, apCPU, apRAM, apIO) //لازم تعديل
	}
	//fmt.Println(applic)
	for i := 0; i < zoneNum; i++ { //intalise zone info
		zn := "zone" + strconv.Itoa(i)
		zCPU = zonefillconstraints.ZoneConstraintsCPU(zonefile, zn)
		zRAM = zonefillconstraints.ZoneConstraintsRAM(zonefile, zn)
		zIO = zonefillconstraints.ZoneConstraintsIO(zonefile, zn)
		zone[i].ZoneInitFunc(zn, zCPU, zRAM, zIO, cellNum)
	}
	//fmt.Println("zone av =")

	//zoneCPUAvg := roundrobingroup.GetzoneCatCPUAvr(zone, zoneNum)
	//zoneRAMAvg := roundrobingroup.GetzoneCatRAMAvr(zone, zoneNum)
	//fmt.Println(zoneCPUAvg)
	//fmt.Println(zoneRAMAvg)
	//fmt.Println(zone)
	groupApp = roundrobingroup.GroupingApp(applic, zone, zoneNum, cellNum, appNum) // to gropu the application in two dimi array perparing to schedule them by Round-Robin algorithm
	fmt.Println("\n\nintlize the application configration ========================================>")
	fmt.Println("\n\n application number ========================================>")
	fmt.Println(appNum)
	fmt.Println("zome number", zoneNum)
	fmt.Println("cell number", cellNum)
	fmt.Println("\n\n application name ========================================>")
	fmt.Println(appName)
	fmt.Println("\n\n application dependencies ========================================>")
	//fmt.Println(appDep)
	//fmt.Print("  ")
	//for i := 0; i < appNum; i++ {
	//	fmt.Print("A", i, "  ")
	//}
	//fmt.Println()
	//for i := 0; i < appNum; i++ {
	//fmt.Print("a", i, "  ")
	//for j := 0; j < appNum; j++ {
	//	fmt.Print(appDep[i][j], "    ")
	//}
	fmt.Println("the number of depedences bwtween appications is ", depndNum)
	//}

	fmt.Println("\n\n application trafic cost matix ========================================>")
	fmt.Println(appTraf)
	fmt.Println("\n\n application app properities ========================================>")
	fmt.Println(applic)
	//fmt.Println(appNum)
	fmt.Println("\n\nintlize the zone configration ========================================>")
	fmt.Println(zone)
	fmt.Println("\n\nintlize the grouping  ========================================>.......")
	fmt.Println(groupApp)

	fmt.Println("\n\nintlize the mapping and scheduling   ========================================>.......")
	mapping = mappingapptozone.MappingApptoZone(groupApp, zoneNum, cellNum, appNum)
	fmt.Println("depndNum=", depndNum)
	mappingapptozone.PrintingZoneLoad(groupApp, mapping, zoneNum, cellNum) //print zone load balancing
	//calculate dependency
	fmt.Println("\n\n application calculate depen ========================================>")
	dendCost, numDeBetZone, numDeInZone = calculatedepen.CalculateDepFunc(appDep, mapping, appTraf, appNum)
	traffic := calculatedepen.CaculateProprotionOFtraffic(dendCost, 0, 0, numDeBetZone, numDeInZone) //trafic cost proportion
	trafficTime := 10.0
	fmt.Println("Traffic cost= ", dendCost, "\n dependency between application in same zone =", numDeInZone,
		"\n dependency between application betwwen zones =", numDeBetZone, "\n Moving Application cost", movingAppCost,
		"\ntotal properation traffic used=", traffic, "\ntotal properation traffic used 10X=", traffic*trafficTime, "\nthe traffic performance=",
	)
	appNuminEacZone := static.CountAppInEachZone(mapping, zoneNum) //the nummber of appliction in each zoneusing for calculate SD
	mean := static.CalculateMean(appNum, zoneNum)
	sd := static.CalculateSD(appNuminEacZone, appNum, zoneNum, mean)
	fmt.Println("SD= ", sd)
	loadbalance := static.LoadbalanceIndector(sd, mean, zoneNum) //load balancing proportion
	fmt.Println("load balaning= ", loadbalance)
	//fmt.Println("cloud Performanc for one time=", calculatedepen.ColudPerformance(loadbalance, 100-(traffic))) //performance propostion
	//fmt.Println("cloud Performanc for one time=", calculatedepen.ColudPerformance(loadbalance, 100-(traffic*trafficTime)))
	fmt.Println("cloud performance after max cut function1  =", calculatedepen.ColudPerformance1(sd, traffic*trafficTime))
	//var app *applicationdis.AppTest
	//result("done")
	//fmt.Println(printResult)
	fmt.Println("\n\n application tree ========================================>")
	fmt.Print("dependen Num")
	fmt.Println(depndNum)
	depdTree = bulddepetree.GetArryofDepndentApp(appDep, mapping, appNum, depndNum) //fill build array of all aplication that have dependency in tree depeendency class from mapping
	fmt.Println(depdTree)
	fmt.Println("application Name\tapplication zone #\t|| Depedent application name\tDepedent application zone number ")
	for i := 0; i < appNum; i++ {

		appFTrname, appFterzoneNum = depdTree[i].PrintFname()
		appCTrname, appCterzoneNum = depdTree[i].PrintCname()
		fmt.Println(appFTrname, "\t\t\t", appFterzoneNum, "\t\t---------------------->\t\t", appCTrname, "\t\t", appCterzoneNum)
		if appFTrname == "" {
			break
		}
	}
	fmt.Println("\n\n application dependencies count ========================================>")
	fmt.Println(maxcut.CalcuDepdFromTreeClass(depdTree)) // "in   --  out"
	fmt.Println("\n\n MAX-CUT ========================================>")
	depCostAfterMaxcut, inDepnBeforMaxcut, otDepnBeforMaxcut, inDependAfterMovind, outDependAfterMovind, appMoviedNum, movingAppCostafterMax := maxcut.MaxCut(depdTree, applic, zone, mapping, depndNum, zoneNum, appNum)
	fmt.Println("This is the tree of the application that has dependencies begin ==================>")
	fmt.Println(depCostAfterMaxcut)
	fmt.Println("This is the tree of the application that has dependencies end ==================>")
	fmt.Println("\n\n maping appliction to the zones after max cut ==================================>")
	maxCutmapping = maxcut.UpdteMappingApp(mapping, depCostAfterMaxcut, depndNum)
	fmt.Println(maxCutmapping)
	maxcut.ShowZoneAferMacCut(maxCutmapping, appNum, zoneNum)
	fmt.Println("maping appliction to the zones after max cut end ==================================>")
	//maxcut.UpdteMappingApp(mapping, depCostAfterMaxcut, depndNum)
	fmt.Println("This is the load balancing after finshing Max Cut begin ================>")
	//lastMacCutMap := mapping
	//maxcut.ShowZoneAferMacCut(lastMacCutMap, appNum, zoneNum)
	//fmt.Println("======")
	//maxcut.ShowZoneAferMacCut(maxCutmapping, appNum, zoneNum)
	fmt.Println("This is the load balancing after finshing Max Cut end ================>")
	appNuminEacZoneAftermaxCut := static.CountAppInEachZone(maxCutmapping, zoneNum) //the nummber of appliction in each zoneusing for calculate SD
	fmt.Println(appNuminEacZoneAftermaxCut)

	fmt.Println("out dep cut=", outDependAfterMovind/6)

	//fmt.Println("cut 20% of Dependincey", calculatedepen.CaculateProprotionOFtraffic((inDependAfterMovind)*10+int(float64(outDependAfterMovind)*.2*100), int(float64(movingAppCostafterMax)*.2), int(float64(appMoviedNum)*.2), inDependAfterMovind, int(float64(outDependAfterMovind)*.2)), "the number of depedncies=", int(float64(outDependAfterMovind)*.2))

	//fmt.Println("cut 40% of Dependincey", calculatedepen.CaculateProprotionOFtraffic((inDependAfterMovind)*10+int(float64(outDependAfterMovind)*.4*100), int(float64(movingAppCostafterMax)*.4), int(float64(appMoviedNum)*.4), inDependAfterMovind, int(float64(outDependAfterMovind)*.4)), "the number of depedncies=", int(float64(outDependAfterMovind)*.4))

	//fmt.Println("cut 60% of Dependincey", calculatedepen.CaculateProprotionOFtraffic((inDependAfterMovind)*10+int(float64(outDependAfterMovind)*.6*100), int(float64(movingAppCostafterMax)*.6), int(float64(appMoviedNum)*.6), inDependAfterMovind, int(float64(outDependAfterMovind)*.6)), "the number of depedncies=", int(float64(outDependAfterMovind)*.6))

	//fmt.Println("cut 80% of Dependincey", calculatedepen.CaculateProprotionOFtraffic((inDependAfterMovind)*10+int(float64(outDependAfterMovind)*.8*100), int(float64(movingAppCostafterMax)*.8), int(float64(appMoviedNum)*.8), inDependAfterMovind, int(float64(outDependAfterMovind)*.8)), "the number of depedncies=", int(float64(outDependAfterMovind)*.8))

	trafficAfterMacCut := calculatedepen.CaculateProprotionOFtraffic(inDependAfterMovind*10+outDependAfterMovind*100, movingAppCostafterMax, appMoviedNum, inDependAfterMovind, outDependAfterMovind)
	fmt.Println("depCostAfterMaxcut", "\nin Depn Befor Maxcut", inDepnBeforMaxcut, "\nout Depn Befor Maxcut", otDepnBeforMaxcut, "\nin Depend After Movind", inDependAfterMovind, "\nout Depend After Movind", outDependAfterMovind, "\nNumber of app movied", appMoviedNum, "\n mpving app cost", movingAppCostafterMax)
	fmt.Println("\n dependency between application in same zone =", inDependAfterMovind, "\n dependency between application betwwen zones =", outDependAfterMovind)
	proprotOfTrafficwithoutMovingApp := calculatedepen.CaculateProprotionOFtraffic(inDependAfterMovind*10+outDependAfterMovind*100, 0, 0, inDependAfterMovind, outDependAfterMovind)
	fmt.Println("the cost of using traffic after moving appliction=", proprotOfTrafficwithoutMovingApp)
	fmt.Println("Trafic cost without moving application 10 time ", proprotOfTrafficwithoutMovingApp*10)
	fmt.Println("traffic for moving app=", calculatedepen.CaculateProprotionOFtraffic(0, movingAppCostafterMax, appMoviedNum, 0, 0))
	fmt.Println("Total traffic used  proprotion=", trafficAfterMacCut)
	fmt.Println("taltal not used of traffic=", 100-trafficAfterMacCut)
	//calculatedepen.CaculateProprotionOFtraffic(dendCost, movingAppCost, 0, numDeBetZone, numDeInZone)
	meanAfterMaxCut := static.CalculateMean(appNum, zoneNum)
	sdAfterMaxCut := static.CalculateSD(appNuminEacZoneAftermaxCut, appNum, zoneNum, meanAfterMaxCut)
	fmt.Println("mean= ", meanAfterMaxCut)
	fmt.Println("SD= ", sdAfterMaxCut)
	loadbalaneAftermaxcut := static.LoadbalanceIndector(sdAfterMaxCut, meanAfterMaxCut, zoneNum) //load balancing propotion
	fmt.Println("load balaning= ", loadbalaneAftermaxcut)                                        //the nummber of appliction in each zoneusing for calculate SD
	//fmt.Println("cloud performance after max cut =", calculatedepen.ColudPerformance(loadbalaneAftermaxcut, 100-trafficAfterMacCut))
	fmt.Println("cloud performance after max cut function1  =", calculatedepen.ColudPerformance1(sdAfterMaxCut, trafficAfterMacCut))
	fmt.Println(maxCutmapping)
	fmt.Println("last stage of the Algorithm =================> ")
	maxcut.ShowZoneAferMacCut(mapping, appNum, zoneNum)
	//laststageMap :=
	enhloadbalancing.EnhanceLoadBalancingFun(depdTree, applic, zone, appDep, mapping, zoneNum, appNum, depndNum)
	//maxcut.ShowZoneAferMacCut(lastMacCutMap, appNum, zoneNum)
	fmt.Println("=================================================\nThe second satage applications arrive for deploy\n=================================================\n")
	fmt.Println("=================================================\nFrist approach\n=================================================\n")
	//maxcut.ShowZoneAferMacCut(laststageMap, appNum, zoneNum)
	//=================================================
	//=================================================
	//=================================================
	//=================================================
	//newAppli, newAppDepne, newdepNum, newAppNum
	//mapping1 := make(map[string]string, appNum)
	//mapping1 = maxCutmapping
	//var zoneF, zoneC int
	//var apnf, apnc string
	//for i := 0; i < len(depdTree); i++ {
	//apnf, zoneF = depdTree[i].PrintFname()
	//apnc, zoneC = depdTree[i].PrintCname()
	//if mapping1[apnf] != "zone"+strconv.Itoa(zoneF) && mapping1[apnc] != "zone"+strconv.Itoa(zoneC) {
	//	mapping1[apnf] = "zone" + strconv.Itoa(zoneF)
	//	mapping1[apnc] = "zone" + strconv.Itoa(zoneC)
	//}
	//}
	//mappingapptozone.PrintingZoneLoad(groupApp, laststageMap, zoneNum, cellNum)
	//for i := 0; i < 40; i++ {
	//for k, v := range maxCutmapping {
	//if v == "zone"+strconv.Itoa(i) {
	//	fmt.Print(" ", k)

	//}
	//}
	//fmt.Println(maxCutmapping)
	//}

	//maxCutmapping = maxcut.UpdteMappingApp1(mapping, depCostAfterMaxcut, depndNum)
	//var zoneF, zoneC int
	//var apnf, apnc string
	//for i := 0; i < len(depdTree); i++ {
	//apnf, zoneF = depdTree[i].PrintFname()
	//apnc, zoneC = depdTree[i].PrintCname()
	//if apnf != "" && apnc != "" {
	//	if maxCutmapping[apnf] != "zone"+strconv.Itoa(zoneF) && maxCutmapping[apnc] != "zone"+strconv.Itoa(zoneC) {
	//	fmt.Println(apnf, "==", maxCutmapping[apnf], apnc, "==", maxCutmapping[apnc])
	//	} else {
	//		fmt.Println("match", apnf, zoneF, "==", maxCutmapping[apnf], apnc, zoneC, "==", maxCutmapping[apnc])
	//	}
	//}
	//	}
	//fmt.Println(maxcut.CalcuDepdFromTreeClass(depdTree))
	//fmt.Println(maxCutmapping)
	loadBalanceLastStageGroup := firstapproach.CreateGroupFroLastStageOfinhanceLB(maxCutmapping, zoneNum, appNum)
	//maxcut.ShowZoneAferMacCut(lastMacCutMap, appNum, zoneNum)
	//for i := 0; i < len(loadBalanceLastStageGroup); i++ {
	//fmt.Println(loadBalanceLastStageGroup[i])
	//}

	newAppli, newAppDepne, newdepNum, newAppNum := readconfi.ReadAppAndDepnConfi(appDep, appNum) //read applications information(name , constraints,dependences )returne compining dependencies matrix and application infrmation array
	//c1 := 0
	//c := 0
	//fmt.Println(appNum, newAppNum)
	//for i := 0; i < newAppNum; i++ {
	//	c1 = 0
	//	c = 0
	//	if i < appNum {
	//	for j := 0; j < appNum; j++ {
	//		if appDep[i][j] == 1 {
	//			c1++
	//		}

	//	}
	//} else {
	//		c1 = 9
	//}
	//	for j := 0; j < newAppNum; j++ {
	//	if newAppDepne[i][j] == 1 {
	//		c++
	//	}
	//	}
	//fmt.Println(i, "=", "c1", c1, "c", c)

	//	}
	newCellNum := newAppNum / zoneNum
	newGroupApp := firstapproach.AddAppToExZone(newAppli, zone, loadBalanceLastStageGroup, zoneNum, cellNum, newAppNum) //round robin grouping
	//for i := 0; i < len(newGroupApp); i++ {
	//fmt.Println(i, newGroupApp[i])
	//}
	newAppMapping := mappingapptozone.MappingApptoZone(newGroupApp, zoneNum, newCellNum, newAppNum) //mapping application to zone
	//fmt.Println(newAppMapping, maxCutmapping)
	mappingapptozone.PrintingZoneLoad(newGroupApp, newAppMapping, zoneNum, newCellNum)

	fmt.Println("=================================================\nSecond approach\n=================================================\n")
	fmt.Println("newdepNum", newdepNum, "newAppNum", newAppNum)
	//maxCutmapping = maxcut.UpdteMappingApp1(newAppMapping, depdTree, newAppNum)
	newAppdepdTree := bulddepetree.GetArryofDepndentAppNewAppArive(newAppDepne, maxCutmapping, newAppNum, appNum, newdepNum, depCostAfterMaxcut) //fill build array of all aplication that have dependency in tree depeendency class from mapping

	fmt.Println("===", len(newAppdepdTree))

	fmt.Println(maxcut.CalcuDepdFromTreeClass(newAppdepdTree))
	trafficloadbalancingcalculate.CalculateTrafficCostandLoadbalancing(newAppdepdTree, newAppDepne, newAppMapping, appTraf, zoneNum, newdepNum, newAppNum) //here need to calcuate the dependencies  very accutite                                            //print zone load balancing//calculate traffic and load balancing
	fmt.Println("dep from tree class")
	//fmt.Println(maxcut.CalcuDepdFromTreeClass(newAppdepdTree))
	//fmt.Println(newAppdepdTree)
	depCostAfterMaxcut, inDepnBeforMaxcut, otDepnBeforMaxcut, inDependAfterMovind, outDependAfterMovind, appMoviedNum, movingAppCostafterMax = maxcut.MaxCut(newAppdepdTree, newAppli, zone, newAppMapping, newdepNum, zoneNum, newAppNum)

	fmt.Println("=================================================\nCalaculation After Finsh Max Cut \n=================================================\n")
	fmt.Println("This is the tree of the application that has dependencies begin ==================>")
	//fmt.Println(depCostAfterMaxcut)
	fmt.Println("This is the tree of the application that has dependencies end ==================>")
	fmt.Println("\n\n maping appliction to the zones after max cut ==================================>")
	maxCutmapping = maxcut.UpdteMappingApp(newAppMapping, newAppdepdTree, newAppNum)
	fmt.Println("maping appliction to the zones after max cut end ==================================>")
	fmt.Println("This is the load balancing after finshing Max Cut begin ================>")
	maxcut.ShowZoneAferMacCut(maxCutmapping, newAppNum, zoneNum)
	fmt.Println("This is the load balancing after finshing Max Cut end ================>")
	appNuminEacZoneAftermaxCut = static.CountAppInEachZone(maxCutmapping, zoneNum) //the nummber of appliction in each zoneusing for calculate SD
	fmt.Println(appNuminEacZoneAftermaxCut)

	fmt.Println("Dependencies Between Zones=", outDependAfterMovind/6)

	trafficAfterMacCut = calculatedepen.CaculateProprotionOFtraffic(inDependAfterMovind*10+outDependAfterMovind*100, movingAppCostafterMax, appMoviedNum, inDependAfterMovind, outDependAfterMovind)
	fmt.Println("depCostAfterMaxcut", "\nin Depn Befor Maxcut", inDepnBeforMaxcut, "\nout Depn Befor Maxcut", otDepnBeforMaxcut, "\nin Depend After Movind", inDependAfterMovind, "\nout Depend After Movind", outDependAfterMovind, "\nNumber of app movied", appMoviedNum, "\n mpving app cost", movingAppCostafterMax)
	fmt.Println("\n dependency between application in same zone =", inDependAfterMovind, "\n dependency between application betwwen zones =", outDependAfterMovind)
	proprotOfTrafficwithoutMovingApp = calculatedepen.CaculateProprotionOFtraffic(inDependAfterMovind*10+outDependAfterMovind*100, 0, 0, inDependAfterMovind, outDependAfterMovind)
	fmt.Println("the cost of using traffic after moving appliction=", proprotOfTrafficwithoutMovingApp)
	fmt.Println("Trafic cost without moving application 10 time ", proprotOfTrafficwithoutMovingApp*10)
	fmt.Println("traffic for moving app=", calculatedepen.CaculateProprotionOFtraffic(0, movingAppCostafterMax, appMoviedNum, 0, 0))
	fmt.Println("Total traffic used  proprotion=", trafficAfterMacCut)
	fmt.Println("taltal not used of traffic=", 100-trafficAfterMacCut)
	//calculatedepen.CaculateProprotionOFtraffic(dendCost, movingAppCost, 0, numDeBetZone, numDeInZone)
	meanAfterMaxCut = static.CalculateMean(appNum, zoneNum)
	sdAfterMaxCut = static.CalculateSD(appNuminEacZoneAftermaxCut, appNum, zoneNum, meanAfterMaxCut)
	fmt.Println("mean= ", meanAfterMaxCut)
	fmt.Println("SD= ", sdAfterMaxCut)
	loadbalaneAftermaxcut = static.LoadbalanceIndector(sdAfterMaxCut, meanAfterMaxCut, zoneNum) //load balancing propotion
	fmt.Println("load balaning= ", loadbalaneAftermaxcut)                                       //the nummber of appliction in each zoneusing for calculate SD
	//========================================
	//=======================================

	//======================================
	enhloadbalancing.EnhanceLoadBalancingFun(newAppdepdTree, newAppli, zone, newAppDepne, maxCutmapping, zoneNum, newAppNum, newdepNum)
}
