package zonefillconstraints

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//ZoneConstraintsCPU read app constraints file zone CPU
func ZoneConstraintsCPU(filenamePath string, zoneName string) int {

	var zoneCPU int
	zoneCPU = 44
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
			if scanner.Text() == "ZONE" {

				for flag1 {
					scanner.Scan()
					if scanner.Text() == "cpu" {

						for flag1 {
							scanner.Scan()
							if scanner.Text() == zoneName {
								scanner.Scan()
								flag1 = false
								num := scanner.Text()

								zoneCPU, err = strconv.Atoi(num)
							}
						}
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
	return zoneCPU

}

//ZoneConstraintsRAM read app constraints file zone CPU
func ZoneConstraintsRAM(filenamePath string, zoneName string) int {

	var zoneRAM int

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
			if scanner.Text() == "ZONE" {

				for flag1 {
					scanner.Scan()
					if scanner.Text() == "ram" {

						for flag1 {
							scanner.Scan()
							if scanner.Text() == zoneName {
								scanner.Scan()
								flag1 = false
								num := scanner.Text()

								zoneRAM, err = strconv.Atoi(num)
							}
						}
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
	return zoneRAM

}

//ZoneConstraintsIO read app constraints file zone CPU
func ZoneConstraintsIO(filenamePath string, zoneName string) string {

	var zoneIO string

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
			if scanner.Text() == "ZONE" {

				for flag1 {
					scanner.Scan()
					if scanner.Text() == "io" {

						for flag1 {
							scanner.Scan()
							if scanner.Text() == zoneName {
								scanner.Scan()
								flag1 = false
								num := scanner.Text()

								zoneIO = num
							}
						}
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
	return zoneIO

}
