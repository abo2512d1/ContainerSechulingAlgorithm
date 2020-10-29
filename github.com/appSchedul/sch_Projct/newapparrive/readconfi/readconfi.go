package readconfi

import (
	"strconv"

	"github.com/appSchedul/sch_Projct/applicationdis/appCreator/appproperties"
	appfillconstraints "github.com/appSchedul/sch_Projct/applicationdis/appCreator/appproperties/appFillConstraints"
	"github.com/appSchedul/sch_Projct/interfacedis/appdepe"
	readappname "github.com/appSchedul/sch_Projct/interfacedis/readAppName"
	readappnum "github.com/appSchedul/sch_Projct/interfacedis/readAppNum"
)

//ReadAppAndDepnConfi    read applications, and dependencies configration for new arraival applicationz
func ReadAppAndDepnConfi(dependold [][]int, appNumOld int) ([]appproperties.ApplProperty, [][]int, int, int) {
	var appfile, appfileConstraint, depndfile string
	if appNumOld < 200 {
		//fmt.Println("/Users/abdullah/go/src/github.com/appSchedul/souceFile/depndConfiNewArrive20.cog")
		appfile = "/Users/abdullah/go/src/github.com/appSchedul/souceFile/appConfiNewArrive20.cog"
		appfileConstraint = "/Users/abdullah/go/src/github.com/appSchedul/souceFile/appConstNewArrive20.cog"
		depndfile = "/Users/abdullah/go/src/github.com/appSchedul/souceFile/depndConfiNewArrive20.cog"
	} else {
		//fmt.Println("/Users/abdullah/go/src/github.com/appSchedul/souceFile/depndConfiNewArrive200.cog")
		appfile = "/Users/abdullah/go/src/github.com/appSchedul/souceFile/appConfiNewArrive20.cog"
		appfileConstraint = "/Users/abdullah/go/src/github.com/appSchedul/souceFile/appConstNewArrive20.cog"
		depndfile = "/Users/abdullah/go/src/github.com/appSchedul/souceFile/depndConfiNewArrive200.cog"
	}
	var ap string   // application name fill the struct
	var apCPU int   // application cpu fill the struct
	var apRAM int   // application ram fill the struct
	var apIO string // application io fill the struct
	j := 0
	appNum := readappnum.ReadConfiAppNum(appfile) //read application number form the file
	appName := make([]string, appNum)
	appName = readappname.ReadConfiAppname(appfile)         //read application name form the file
	var applic = make([]appproperties.ApplProperty, appNum) //build applications infromation calss
	for i := appNumOld; i < appNumOld+appNum; i++ {         //fill application from config file

		ap = "A" + strconv.Itoa(j)
		apCPU = appfillconstraints.ReadAppConstraintsCPU(appfileConstraint, ap)
		apRAM = appfillconstraints.ReadAppConstraintsRAM(appfileConstraint, ap)
		apIO = appfillconstraints.ReadAppConstraintsIO(appfileConstraint, ap)
		ap = "A" + strconv.Itoa(i)
		applic[j].ApplPropertyFu(appName[j], ap, 1, apCPU, apRAM, apIO) //fill the application spicification array
		j++
	}
	//finsh reading application informations

	//start reading applications dependencies
	var depen = make([][]int, appNum)
	var newAppDepen = make([][]int, appNum)
	depen = appdepe.ReadConfiAppDep(appfile, depndfile, appNumOld)
	//m := 0
	for i := 0; i < appNum; i++ {
		newAppDepen[i] = depen[i] //this the dependencies of the new arrival applications
	}

	// comaining the dependencies matrix old application dependencies and new arrival applications dependencies
	compAppdepen := make([][]int, appNum+appNumOld)

	for i := 0; i < appNum+appNumOld; i++ {
		compAppdepen[i] = make([]int, appNum+appNumOld)

		for j := 0; j < appNum+appNumOld; j++ {

			if i < appNumOld {
				if j < appNumOld {
					compAppdepen[i][j] = dependold[i][j]
				} else {
					compAppdepen[i][j] = 0
				}
			} else {

				compAppdepen[i][j] = newAppDepen[i-appNumOld][j]

			}
		}

	}

	depndNum := appdepe.ReadDependNum(compAppdepen, appNum+appNumOld) //count the dpendentst number
	return applic, compAppdepen, depndNum, appNum + appNumOld
}
