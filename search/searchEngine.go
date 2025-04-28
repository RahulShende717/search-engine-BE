package search

import (
	"fmt"
	"strings"
	"time"
)

type Record struct {
	MsgId          string
	PartitionId    string
	Timestamp      string
	Hostname       string
	Priority       int64
	Facility       string
	FacilityString string
	Severity       string
	SeverityString string
	AppName        string
	ProcId         string
	Message        string `parquet:"name=Message, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	MessageRaw     string `parquet:"name=MessageRaw, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	StructuredData string `parquet:"name=StructuredData, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Tag            string `parquet:"name=Tag, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Sender         string `parquet:"name=Sender, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Groupings      string `parquet:"name=Groupings, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Event          string `parquet:"name=Event, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	EventId        string `parquet:"name=EventId, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	NanoTimeStamp  string `parquet:"name=NanoTimeStamp, type=INT64"`
	// Namespace      string `parquet:"name=Namespace, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
}

var records []Record

func LoadData(data []Record) {
	records = data
}

func Search(query string) ([]Record, time.Duration) {
	start := time.Now()
	query = strings.ToLower(query)

	fmt.Println("inside search")
	var results []Record
	for _, record := range records {
		if strings.Contains(strings.ToLower(record.Message), query) ||
			strings.Contains(strings.ToLower(record.MessageRaw), query) ||
			strings.Contains(strings.ToLower(record.Tag), query) ||
			strings.Contains(strings.ToLower(record.Sender), query) {
			results = append(results, record)
		}
	}

	fmt.Println("inside search results===>>>", results)
	return results, time.Since(start)
}
