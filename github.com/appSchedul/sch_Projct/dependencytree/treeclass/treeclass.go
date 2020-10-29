package treeclass

//DepTree is the class for dependency tree
type DepTree struct {
	appFname string
	appCname string
	zomeNumF int
	zomeNumC int
	moved    bool
}

//{appFname:make(map[string]string)}

//SetappFnamFun function that enter class father name
func (sefn *DepTree) SetappFnamFun(appfNme string, zomeNumF int) {

	sefn.appFname = appfNme
	sefn.zomeNumF = zomeNumF

}

//SetappCnamFun function that enter class father name
func (sefn *DepTree) SetappCnamFun(appfNme string, zomeNumC int) {

	sefn.appCname = appfNme
	sefn.zomeNumC = zomeNumC
}

//PrintFname father application name and zone number
func (sefn *DepTree) PrintFname() (string, int) {
	return sefn.appFname, sefn.zomeNumF
}

//PrintCname children application name and zone number
func (sefn *DepTree) PrintCname() (string, int) {
	return sefn.appCname, sefn.zomeNumC
}

//SetMovingStituos this change the moving statuos for application
func (sefn *DepTree) SetMovingStituos(b bool) {
	sefn.moved = b
}

//GetMovingStituos this change the moving statuos for application
func (sefn *DepTree) GetMovingStituos() bool {
	return sefn.moved
}

//ReturnMovingStituos this change the moving statuos for application
func (sefn *DepTree) ReturnMovingStituos() bool {
	return sefn.moved

}
