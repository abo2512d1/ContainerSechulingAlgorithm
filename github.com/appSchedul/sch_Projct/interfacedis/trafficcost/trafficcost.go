package trafficcost

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//ReadConfiTrafficCost function calculate traffic cost beween zone
func ReadConfiTrafficCost(filenamePath string) [][]int {

	appNum := 2 //number of array dimaintion of traffic matrix
	flag, flag1 := true, true
	file, err := os.Open(filenamePath)
	var appTraCo = make([][]int, appNum)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for flag {
		scanner.Scan()
		if scanner.Text() == "Application" {

			for flag1 {
				scanner.Scan()
				if scanner.Text() == "traffic" {
					for i := 0; i < appNum; i++ {
						appTraCo[i] = make([]int, appNum)
						for j := 0; j < appNum; j++ {
							scanner.Scan()

							num := scanner.Text()

							appTraCo[i][j], err = strconv.Atoi(num)
						}
					}
					flag1 = false
				}
			}
		}

		flag = false
	}
	return appTraCo

}

//ReadConfiZoneNum read zone number
func ReadConfiZoneNum(filenamePath1 string, filenamePath string) int {

	//appNum := readappnum.ReadConfiAppNum(filenamePath1) //read conf file to read applcations number
	flag, flag1 := true, true
	file, err := os.Open(filenamePath)
	num := ""
	//var appTraCo = make([][]int, appNum)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for flag {
		scanner.Scan()
		if scanner.Text() == "ZONE" {

			for flag1 {
				scanner.Scan()
				if scanner.Text() == "number" {
					//for i := 0; i < appNum; i++ {
					//appTraCo[i] = make([]int, appNum)
					//for j := 0; j < appNum; j++ {
					scanner.Scan()

					num = scanner.Text()

					//appTraCo[i][j], err = strconv.Atoi(num)
					//	}
					//}
					flag1 = false
				}
			}
		}

		flag = false
	}
	num1, err := strconv.Atoi(num)
	return num1

}
