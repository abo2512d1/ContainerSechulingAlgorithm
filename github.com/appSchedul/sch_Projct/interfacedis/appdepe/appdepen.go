package appdepe

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	readappnum "github.com/appSchedul/sch_Projct/interfacedis/readAppNum"
)

//ReadConfiAppDep read conf file to calculte application depencency
func ReadConfiAppDep(filenamePath1 string, filenamePath2 string, oldappNum int) [][]int {

	appNum := readappnum.ReadConfiAppNum(filenamePath1) //read conf file to read applcations number
	if filenamePath1 == "/Users/abdullah/go/src/github.com/appSchedul/souceFile/appConfiNewArrive20.cog" {
		appNum = appNum + oldappNum
	}
	var appDep = make([][]int, appNum, appNum)
	flag, flag1 := true, true
	file, err := os.Open(filenamePath2)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for flag {
		scanner.Scan()
		if scanner.Text() != "END" {
			if scanner.Text() == "Application" {

				for flag1 {
					scanner.Scan()
					if scanner.Text() == "dependency" {
						for i := 0; i < appNum; i++ {
							appDep[i] = make([]int, appNum)
							for j := 0; j < appNum; j++ {
								scanner.Scan()

								num := scanner.Text()

								appDep[i][j], err = strconv.Atoi(num)
							}
						}
						flag1 = false
					} else if scanner.Text() == "END" {
						flag1 = false
						flag = false
					}
				}
			} else if scanner.Text() == "END" {
				flag1 = false
				flag = false
			}
		} else if scanner.Text() == "END" {
			flag1 = false
			flag = false
		}
		flag = false
	}
	return appDep

}

//ReadDependNum  thei sfunction readd the dependency number
func ReadDependNum(depd [][]int, appNum int) int {
	var count int
	count = 0
	for i := 0; i < len(depd); i++ {
		for j := 0; j < len(depd); j++ {
			if depd[i][j] == 1 {
				count++
			}
		}

	}
	fmt.Println("count=", count)
	return count
}
