package datatypes

import (
	"fmt"
	"strconv"
	"strings"
)

type Metric struct {
	MetricName string
	Value      float64
	MetricType string
	SampleRate int
	Tags       map[string]string
}

func (m Metric) ToDataDog() string {
	metricString := fmt.Sprintf(
		"%s:%f|%s|@%d",
		m.MetricName,
		m.Value,
		m.MetricType,
		m.SampleRate)

	tagString := ""
	for k, v := range m.Tags {
		if v != "" {
			tagString = fmt.Sprintf("%s%s:%s", tagString, k, v)
		} else {
			tagString = fmt.Sprintf("%s%s", tagString, k)
		}

		tagString = fmt.Sprintf("%s,", tagString)
	}

	tagString = fmt.Sprintf("|#%s", strings.TrimRight(tagString, ","))

	return fmt.Sprintf("%s%s", metricString, tagString)
}

func (m Metric) ToNewRelic() string {
	return ""
}

func (m Metric) ToFormat(format string) string {
	return ""
}

func parseMetric(packet string) DataType {

	//var err error

	var m Metric
	m.Tags = make(map[string]string)
	m.SampleRate = 1

	pieces := strings.Split(packet, "|")
	metricValuePair := strings.Split(pieces[0], ":")

	m.MetricName = metricValuePair[0]
	Value, _ := strconv.ParseFloat(metricValuePair[1], 64)
	m.Value = Value

	m.MetricType = pieces[1]

	for _, piece := range pieces[2:] {
		if strings.HasPrefix(piece, "@") {
			sampleRate := strings.TrimPrefix(piece, "@")
			m.SampleRate, _ = strconv.Atoi(sampleRate)
		} else if strings.HasPrefix(piece, "#") {
			tags := strings.TrimPrefix(piece, "#")
			pairs := strings.Split(tags, ",")
			for _, pair := range pairs {
				split := strings.Split(pair, ":")
				if len(split) == 1 {
					m.Tags[split[0]] = ""
				} else if len(split) == 2 {
					m.Tags[split[0]] = split[1]
				} else {

				}
			}
		}
	}

	return m
}
