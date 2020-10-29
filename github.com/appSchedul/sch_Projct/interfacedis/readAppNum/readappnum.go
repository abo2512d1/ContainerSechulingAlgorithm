package readappnum

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//ReadConfiAppNum read conf file application number
func ReadConfiAppNum(filenamePath string) int {

	var appNum int
	flag, flag1 := true, true
	file, err := os.Open(filenamePath)
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
					if scanner.Text() == "number" {
						scanner.Scan()
						flag1 = false
						num := scanner.Text()

						appNum, err = strconv.Atoi(num)
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
	return appNum

}
