package appfillconstraints

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//ReadAppConstraintsCPU read app constraints file application CPU
func ReadAppConstraintsCPU(filenamePath string, appAname string) int {

	var appCPU int
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
					if scanner.Text() == "CPU" {
						for flag1 {
							scanner.Scan()
							if scanner.Text() == appAname {
								scanner.Scan()
								flag1 = false
								num := scanner.Text()

								appCPU, err = strconv.Atoi(num)
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

	flag = false

	return appCPU

}

//ReadAppConstraintsRAM read app constraints file application RAM
func ReadAppConstraintsRAM(filenamePath string, appAname string) int {

	var appRAM int
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
					if scanner.Text() == "RAM" {
						for flag1 {
							scanner.Scan()
							if scanner.Text() == appAname {
								scanner.Scan()
								flag1 = false
								num := scanner.Text()

								appRAM, err = strconv.Atoi(num)
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

	flag = false

	return appRAM

}

//ReadAppConstraintsIO read app constraints file application CPU
func ReadAppConstraintsIO(filenamePath string, appAname string) string {

	var appIO string
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
					if scanner.Text() == "IO" {
						for flag1 {
							scanner.Scan()
							if scanner.Text() == appAname {
								scanner.Scan()
								flag1 = false
								num := scanner.Text()

								appIO = num
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

	flag = false

	return appIO

}
