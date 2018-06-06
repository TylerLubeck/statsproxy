package datatypes

import (
	"strings"
)

type Type string

const (
	METRIC       Type = "metric"
	STATUS_CHECK Type = "status_check"
	EVENT        Type = "event"
)

type DataType interface {
	ToString() string
	GetType() Type
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
