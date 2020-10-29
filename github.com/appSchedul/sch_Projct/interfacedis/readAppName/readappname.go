package readappname

import (
	"bufio"
	"fmt"
	"os"

	readappnum "github.com/appSchedul/sch_Projct/interfacedis/readAppNum"
)

//interfacedis.ReadConfiAppNu

//ReadConfiAppname read conf file application name
func ReadConfiAppname(filenamePath string) []string {
	appNum := readappnum.ReadConfiAppNum(filenamePath) //read conf file to read applcations number
	appName := make([]string, appNum)
	flag, flag1 := true, true

	file, err := os.Open(filenamePath)
	//fileSize, err := file.Stat()
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
					if scanner.Text() == "name" {
						for i := 0; i < appNum; i++ {
							scanner.Scan()

							appName[i] = scanner.Text()

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
	return appName

}
