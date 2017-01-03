package main

import (
	"flag"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/rorymckinley/timetracker/data"
	"github.com/rorymckinley/timetracker/trackingdata"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var Category = flag.String("cat", "", "Category")
var StartTime = flag.String("start", "", "Start Time")
var EndTime = flag.String("end", "", "End Time")
var Subs = flag.String("subs", "", "Subcategories")

type ActionData struct {
	Action string
	data.Data
}

func read(location string) trackingdata.TrackingData {
	trackingData := trackingdata.TrackingData{}

	fileData, _ := ioutil.ReadFile(location)
	err := yaml.Unmarshal(fileData, &trackingData)

	if err != nil {
		spew.Dump(err)
	}

	return trackingData
}

func persist(data *trackingdata.TrackingData, location string) {
	export, _ := yaml.Marshal(data)
	f, _ := os.Create(location)
	f.Write(export)
	f.Close()
}

func determineConfig() ActionData {
	spew.Dump(flag.Args())
	switch {
	case len(flag.Args()) == 0:
		return ActionData{Action: "toggle-last"}
	case len(flag.Args()) == 1:
		data := data.Data{Description: flag.Arg(0)}
		SetCategory(&data)
		SetStartTime(&data)
		SetEndTime(&data)
		SetSubcategories(&data)
		return ActionData{Action: "create-new", Data: data}
	default:
		return ActionData{}
	}
}

func SetCategory(data *data.Data) {
	if *Category != "" {
		data.Category = *Category
	} else {
		fmt.Println("Please supply a category")
		os.Exit(1)
	}
}

func SetStartTime(data *data.Data) {
	const dateTemplate = "2006 January 2"
	const timeTemplate = "2006 January 2 15:04 -0700"
	if *StartTime != "" {
		input := time.Now().Format(dateTemplate) + " " + *StartTime + " +0200"
		data.StartTime, _ = time.Parse(timeTemplate, input)
	} else {
		data.StartTime = time.Now()
	}
}

func SetEndTime(data *data.Data) {
	const dateTemplate = "2006 January 2"
	const timeTemplate = "2006 January 2 15:04 -0700"
	if *EndTime != "" {
		input := time.Now().Format(dateTemplate) + " " + *EndTime + " +0200"
		data.EndTime, _ = time.Parse(timeTemplate, input)
	}
}

func SetSubcategories(data *data.Data) {
	data.Subcategories = strings.Split(*Subs, ",")
}

func main() {
	flag.Parse()
	config := determineConfig()

	trackingData := read("./tracking.yml")

	switch config.Action {
	case "toggle-last":
		trackingData.ToggleLast()
	case "create-new":
		trackingData.CloseLastEvent()
		trackingData.AddEvent(config.Data)
	}

	persist(&trackingData, "./tracking.yml")
}
