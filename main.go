package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/PuloV/ics-golang"
)

func main() {
	parser := ics.New()
	ics.RepeatRuleApply = true
	ics.MaxRepeats = 200
	inputChan := parser.GetInputChan()
	inputChan <- ""
	parser.Wait()
	cal, err := parser.GetCalendars()
	if err == nil {
		now := time.Now()
		var allEvents []ics.Event

		//Füge alle zukünftigen und aktuellen Termine in ein Array zusammen
		for _, calendar := range cal {
			//  Alle Events, die noch nicht zuende sind werden gesammelt
			for _, event := range calendar.GetEvents() {
				if now.Before(event.GetEnd()) {
					allEvents = append(allEvents, event)
				}
			}
		}

		//Sortiere die Termine
		sort.Slice(allEvents, func(i, j int) bool {
			return allEvents[i].GetStart().Before(allEvents[j].GetStart())
		})

		for i, event := range allEvents {
			if i >= 10 {
				break
			}
			fmt.Printf("%s am %d.%s.%d\n", event.GetSummary(), event.GetStart().Day(), event.GetStart().Month().String(), event.GetStart().Year())
		}

		fmt.Println("Das sind die nächsten 10 Termine")

	}
}
