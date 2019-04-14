package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"time"

	"teahub.io/momar/config"

	ics "github.com/rjhorniii/ics-golang"
)

func main() {
	ics.DeleteTempFiles = true
	c := config.Open("/var/E-Ink/config.yaml")
	fmt.Println("Update Data....")
	everything := []ics.Event{}
	for _, e := range c.Get("Server").StringList() {
		everything = append(everything, next5Events(e)...)
	}

	//Sortiere die Termine
	sort.Slice(everything, func(i, j int) bool {
		return everything[i].GetStart().Before(everything[j].GetStart())
	})

	returnString := ""
	for i, event := range everything {
		if i >= 5 {
			break
		}

		if len(event.GetDTZID()) == 0 {
			returnString += fmt.Sprintf(" %s; am %d.%d von %d:%02d bis %d:%02d Uhr; Ort: %s;", event.GetSummary(), event.GetStart().Local().Day(), event.GetStart().Local().Month(), event.GetStart().Local().Hour(), event.GetStart().Local().Minute(), event.GetEnd().Local().Hour(), event.GetEnd().Local().Minute(), event.GetLocation())
		} else {
			returnString += fmt.Sprintf(" %s; am %d.%d von %d:%02d bis %d:%02d Uhr; Ort: %s;", event.GetSummary(), event.GetStart().Day(), event.GetStart().Month(), event.GetStart().Hour(), event.GetStart().Minute(), event.GetEnd().Hour(), event.GetEnd().Minute(), event.GetLocation())
		}

	}

	err := ioutil.WriteFile("/var/E-Ink/cache.txt", []byte(stamp()+returnString), 0644)
	if err != nil {
		fmt.Println("Error writing File")
	}
	fmt.Println("Update done.")
}

func stamp() string {
	return strconv.Itoa(time.Now().Day()) + ". " + time.Now().Month().String() + " " + strconv.Itoa(time.Now().Year()) + ";"
}

func next5Events(url string) []ics.Event {
	fmt.Printf("parsing next 7 Events from %s\n", url)
	parser := ics.New()
	ics.RepeatRuleApply = true
	ics.MaxRepeats = 200
	inputChan := parser.GetInputChan()
	inputChan <- url
	parser.Wait()
	cal, err := parser.GetCalendars()
	if err == nil {
		now := time.Now()
		var allEvents []ics.Event

		//Füge alle zukünftigen und aktuellen Termine in ein Array zusammen
		for _, calendar := range cal {
			//  Alle Events, die noch nicht zuende sind werden gesammelt
			for _, event := range calendar.GetUpcomingEvents(7) {
				if now.Before(event.GetEnd()) {
					allEvents = append(allEvents, event)
				}
			}
		}
		fmt.Printf("next 7 events are: \n %v\n", allEvents)
		return allEvents
	}
	return nil
}
