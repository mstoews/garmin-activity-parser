package utils

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"
	_ "strconv"
)

const minutes = 60
const seconds = 60
const meters = 1000

type TrainingCenterDatabase struct {
	XMLName        xml.Name `xml:"TrainingCenterDatabase"`
	Text           string   `xml:",chardata"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Ns5            string   `xml:"ns5,attr"`
	Ns3            string   `xml:"ns3,attr"`
	Ns2            string   `xml:"ns2,attr"`
	Xmlns          string   `xml:"xmlns,attr"`
	Xsi            string   `xml:"xsi,attr"`
	Ns4            string   `xml:"ns4,attr"`
	Activities     struct {
		Text     string `xml:",chardata"`
		Activity struct {
			Text  string `xml:",chardata"`
			Sport string `xml:"Sport,attr"`
			ID    string `xml:"Id"`
			Lap   []struct {
				Text                string `xml:",chardata"`
				StartTime           string `xml:"StartTime,attr"`
				TotalTimeSeconds    string `xml:"TotalTimeSeconds"`
				DistanceMeters      string `xml:"DistanceMeters"`
				MaximumSpeed        string `xml:"MaximumSpeed"`
				Calories            string `xml:"Calories"`
				AverageHeartRateBpm struct {
					Text  string `xml:",chardata"`
					Value string `xml:"Value"`
				} `xml:"AverageHeartRateBpm"`
				MaximumHeartRateBpm struct {
					Text  string `xml:",chardata"`
					Value string `xml:"Value"`
				} `xml:"MaximumHeartRateBpm"`
				Intensity     string `xml:"Intensity"`
				TriggerMethod string `xml:"TriggerMethod"`
				Track         struct {
					Text       string `xml:",chardata"`
					Trackpoint []struct {
						Text           string `xml:",chardata"`
						Time           string `xml:"Time"`
						DistanceMeters string `xml:"DistanceMeters"`
						HeartRateBpm   struct {
							Text  string `xml:",chardata"`
							Value string `xml:"Value"`
						} `xml:"HeartRateBpm"`
						Extensions struct {
							Text string `xml:",chardata"`
							TPX  struct {
								Text       string `xml:",chardata"`
								Speed      string `xml:"Speed"`
								RunCadence string `xml:"RunCadence"`
							} `xml:"TPX"`
						} `xml:"Extensions"`
					} `xml:"Trackpoint"`
				} `xml:"Track"`
				Extensions struct {
					Text string `xml:",chardata"`
					LX   struct {
						Text          string `xml:",chardata"`
						AvgSpeed      string `xml:"AvgSpeed"`
						AvgRunCadence string `xml:"AvgRunCadence"`
						MaxRunCadence string `xml:"MaxRunCadence"`
					} `xml:"LX"`
				} `xml:"Extensions"`
			} `xml:"Lap"`
			Creator struct {
				Text      string `xml:",chardata"`
				Type      string `xml:"type,attr"`
				Name      string `xml:"Name"`
				UnitId    string `xml:"UnitId"`
				ProductID string `xml:"ProductID"`
				Version   struct {
					Text         string `xml:",chardata"`
					VersionMajor string `xml:"VersionMajor"`
					VersionMinor string `xml:"VersionMinor"`
					BuildMajor   string `xml:"BuildMajor"`
					BuildMinor   string `xml:"BuildMinor"`
				} `xml:"Version"`
			} `xml:"Creator"`
		} `xml:"Activity"`
	} `xml:"Activities"`
}

func ProcessActivity(filename string, db *sql.DB) {
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Printf("Successfully Opened %s\n", filename)

	defer xmlFile.Close()

	byteValue, _ := io.ReadAll(xmlFile)

	// fmt.Printf("%s", byteValue)
	var trn TrainingCenterDatabase
	xml.Unmarshal(byteValue, &trn)
	insertActivities(&trn)
}

func ProcessActivityNoDB(filename string) {
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Printf("Successfully Opened %s\n", filename)

	defer xmlFile.Close()

	byteValue, _ := io.ReadAll(xmlFile)

	// fmt.Printf("%s", byteValue)
	var trn TrainingCenterDatabase
	xml.Unmarshal(byteValue, &trn)
	insertActivities(&trn)
}

func insertActivities(trn *TrainingCenterDatabase) {
	fmt.Print("Time, Distance in Meters, HearRate, Velocity, Cadence, StrideLength ")
	for index := 0; index < len(trn.Activities.Activity.Lap); index++ {
		for i := 0; i < len(trn.Activities.Activity.Lap[index].Track.Trackpoint); i++ {
			var speed, _ = strconv.ParseFloat(trn.Activities.Activity.Lap[index].Track.Trackpoint[i].Extensions.TPX.Speed, 8)
			var kph = speed * seconds * minutes / meters
			var cadence, _ = strconv.ParseFloat(trn.Activities.Activity.Lap[index].Track.Trackpoint[i].Extensions.TPX.RunCadence, 8)
			var strideLength = speed * seconds * .5 / cadence
			fmt.Print(
				trn.Activities.Activity.Lap[index].Track.Trackpoint[i].Time, ",",
				trn.Activities.Activity.Lap[index].Track.Trackpoint[i].DistanceMeters, ",",
				trn.Activities.Activity.Lap[index].Track.Trackpoint[i].HeartRateBpm.Value, ",",
				kph, ",",
				speed, ",",
				cadence, ",",
				strideLength, "")
			fmt.Println()
		}
		//fmt.Printf(trn.Activities.Activity.Lap[0].Calories)
		//fmt.Printf(",")
		//fmt.Printf(trn.Activities.Activity.Lap[index].Extensions.LX.AvgSpeed)
		fmt.Println("")
	}
}

/*
Stride length is equivalent to velocity divided by 0.5 cadence (speed/0.5 cadence).
Our athlete with a gait velocity of 70 m/min and a cadence of 100 steps per minute would have a
calculated stride length of 1.4 m: (70 m/min)/(0.5) × 100 steps per minute = stride length of 1.4.

*/
