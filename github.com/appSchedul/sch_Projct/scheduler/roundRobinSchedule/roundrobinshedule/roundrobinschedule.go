package roundrobinshedule

import "strconv"

//RoundRobinSchedul function that schedule grouping application (mapping the applictions to zones)
func RoundRobinSchedul(groups [][]string, zoneNum int, appNum int) [][]map[string]string {
	mapping := make([][]map[string]string, 4, 10)

	for i := 0; i < 4; i++ {
		//mapping[i] = make([]map[string]string, appNum)

		for j := 0; j < 10; j++ {

			mapping[i][j][groups[i][j]] = "zone" + strconv.Itoa(i)

		}
	}

	return mapping
}
