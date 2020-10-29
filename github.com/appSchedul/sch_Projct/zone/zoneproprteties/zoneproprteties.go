package zoneproprteties

//ZoneProperity is struct to build zone propreteites
type ZoneProperity struct {
	zoneName     string
	zoneCPU      int
	zoneRAM      int
	zoneIO       string
	clusterNum   int
	cellAdaptApp []string
}

// ZoneInitFunc to fell the zone information
func (zone *ZoneProperity) ZoneInitFunc(znam string, zcpu int, zram int, zio string, claunum int) {
	zone.zoneName = znam
	zone.zoneCPU = zcpu
	zone.zoneRAM = zram
	zone.zoneIO = zio
	zone.clusterNum = claunum
}

//GetZoneName return zone cpu
func (zone *ZoneProperity) GetZoneName() string {
	return zone.zoneName
}

//GetZoneCPU return zone cpu
func (zone *ZoneProperity) GetZoneCPU() int {
	return zone.zoneCPU
}

//GetZoneRAM return zone RAM
func (zone *ZoneProperity) GetZoneRAM() int {
	return zone.zoneRAM
}

//GetZoneIO return zone IO
func (zone *ZoneProperity) GetZoneIO() string {
	return zone.zoneIO
}
