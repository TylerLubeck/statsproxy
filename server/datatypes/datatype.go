package datatypes

import (
	"strings"
)

type DataType interface {
	ToNewRelic() string
	ToDataDog() string
	ToFormat(string) string
}

func ParseDataPacket(packet string) DataType {

	var p DataType

	if strings.HasPrefix(packet, "_sc") {
		//p = parseServiceCheck(packet)
	} else if strings.HasPrefix(packet, "_e") {
		//p = parseEvent(packet)
	} else {
		p = parseMetric(packet)
	}

	return p
}

func parseEvent(packet string) *DataType {
	return nil

}

func parseServiceCheck(packet string) *DataType {
	return nil

}
