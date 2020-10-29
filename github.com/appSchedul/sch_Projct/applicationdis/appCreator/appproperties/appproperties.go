//appproperties cearte and fill the application property sturct

package appproperties

import "fmt"

// ApplProperty is struct for application properties
type ApplProperty struct {
	appfname     string
	appAname     string
	appPrioprity int
	appCPU       int
	appRAM       int
	appIo        string
}

//ApplPropertyFu constractor for class ApplProperty fill the name of app and cpu, ram capicity and io avalaiblity
func (app *ApplProperty) ApplPropertyFu(af string, aa string, priority int, acpu int, aram int, aio string) {
	app.appfname = af
	app.appAname = aa
	app.appPrioprity = priority
	app.appCPU = acpu
	app.appRAM = aram
	app.appIo = aio

}

//GetAppAname return application name
func (app *ApplProperty) GetAppAname() string {
	return app.appAname
}

//GetAppCPU return application CPU
func (app *ApplProperty) GetAppCPU() int {
	return app.appCPU
}

//GetAppPri return application priority
func (app *ApplProperty) GetAppPri() int {
	return app.appPrioprity
}

//GetAppIO return application priority
func (app *ApplProperty) GetAppIO() string {
	return app.appIo
}

//GetAppRAM return application RAM
func (app *ApplProperty) GetAppRAM() int {
	return app.appRAM
}

// ApplProprinting print appliction properities
func (app ApplProperty) ApplProprinting() {
	fmt.Println("application full name: ", app.appfname)
	fmt.Println("application full apperfration name: ", app.appAname)
	fmt.Println("application requriment of CPU: ", app.appCPU)
	fmt.Println("application requriment of RAM: ", app.appRAM)
	fmt.Println("is the application needs IO: ", app.appIo)
}
