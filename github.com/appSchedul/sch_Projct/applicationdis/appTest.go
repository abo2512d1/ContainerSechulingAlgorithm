package applicationdis

//AppTest apptest pakage name  the struct
type AppTest struct {
	//for test
	testResult string
}

// PrintResult print result
func (a AppTest) PrintResult() string {
	return a.testResult
}

//Result fill calss intity
func (a *AppTest) Result(a1 string) {
	a.testResult = a1
}

// S va
var S string = "jhfvhg"

//App d
var App AppTest
