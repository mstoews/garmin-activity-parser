package utils

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"os"
) 

type Welcome4 struct {
	TrainingCenterDatabase TrainingCenterDatabase `json:"TrainingCenterDatabase"`
}

type TrainingCenterDatabase struct {
	Activities        Activities `json:"Activities"`         
	XmlnsNs5          string     `json:"_xmlns:ns5"`         
	XmlnsNs3          string     `json:"_xmlns:ns3"`         
	XmlnsNs2          string     `json:"_xmlns:ns2"`         
	Xmlns             string     `json:"_xmlns"`             
	XmlnsXsi          string     `json:"_xmlns:xsi"`         
	XmlnsNs4          string     `json:"_xmlns:ns4"`         
	XsiSchemaLocation string     `json:"_xsi:schemaLocation"`
}

type Activities struct {
	Activity Activity `json:"Activity"`
}

type Activity struct {
	ID      string  `json:"Id"`     
	Lap     []Lap   `json:"Lap"`    
	Creator Creator `json:"Creator"`
	Sport   string  `json:"_Sport"` 
}

type Creator struct {
	Name      string  `json:"Name"`     
	UnitID    string  `json:"UnitId"`   
	ProductID string  `json:"ProductID"`
	Version   Version `json:"Version"`  
	XsiType   string  `json:"_xsi:type"`
}

type Version struct {
	VersionMajor string `json:"VersionMajor"`
	VersionMinor string `json:"VersionMinor"`
	BuildMajor   string `json:"BuildMajor"`  
	BuildMinor   string `json:"BuildMinor"`  
}

type Lap struct {
	TotalTimeSeconds    string        `json:"TotalTimeSeconds"`   
	DistanceMeters      string        `json:"DistanceMeters"`     
	MaximumSpeed        string        `json:"MaximumSpeed"`       
	Calories            string        `json:"Calories"`           
	AverageHeartRateBPM HeartRateBPM  `json:"AverageHeartRateBpm"`
	MaximumHeartRateBPM HeartRateBPM  `json:"MaximumHeartRateBpm"`
	Intensity           string        `json:"Intensity"`          
	TriggerMethod       string        `json:"TriggerMethod"`      
	Track               Track         `json:"Track"`              
	Extensions          LapExtensions `json:"Extensions"`         
	StartTime           string        `json:"_StartTime"`         
}

type HeartRateBPM struct {
	Value string `json:"Value"`
}

type LapExtensions struct {
	Lx Lx `json:"LX"`
}

type Lx struct {
	AvgSpeed      AvgRunCadence `json:"AvgSpeed"`     
	AvgRunCadence AvgRunCadence `json:"AvgRunCadence"`
	MaxRunCadence AvgRunCadence `json:"MaxRunCadence"`
	Prefix        Prefix        `json:"__prefix"`     
}

type AvgRunCadence struct {
	Prefix Prefix `json:"__prefix"`
	Text   string `json:"__text"`  
}

type Track struct {
	Trackpoint []Trackpoint `json:"Trackpoint"`
}

type Trackpoint struct {
	Time           string               `json:"Time"`          
	DistanceMeters string               `json:"DistanceMeters"`
	HeartRateBPM   HeartRateBPM         `json:"HeartRateBpm"`  
	Extensions     TrackpointExtensions `json:"Extensions"`    
}

type TrackpointExtensions struct {
	Tpx Tpx `json:"TPX"`
}

type Tpx struct {
	Speed      AvgRunCadence `json:"Speed"`     
	RunCadence AvgRunCadence `json:"RunCadence"`
	Prefix     Prefix        `json:"__prefix"`  
}

type Prefix string
const (
	Ns3 Prefix = "ns3"
)

func ProcessActivity(filename string, db *sql.DB) {
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Successfully Opened %s\n", filename)
	
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	var act Activities

	xml.Unmarshal(byteValue, &act)
	insertActivities(&act, db)
	
}

func insertActivities(act *Activities, db *sql.DB) {
	for index := 0; index < len(act.Activity.Lap ); index++ {
		// lap(act.Activity.Lap, db, index)
	}
}

func lap(lap *Lap, db *sql.DB, index int) {

}
