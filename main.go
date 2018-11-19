package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/rjhorniii/ics-golang"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, stamp()+next10Events(""))
	})
	r.Run(":6342")
}

func stamp() string {
	return strconv.Itoa(time.Now().Day()) + ". " + time.Now().Month().String() + " " + strconv.Itoa(time.Now().Year()) + ";"
}

func next10Events(url string) string {
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
			for _, event := range calendar.GetUpcomingEvents(5) {
				if now.Before(event.GetEnd()) {
					allEvents = append(allEvents, event)
				}
			}
		}

		//Sortiere die Termine
		sort.Slice(allEvents, func(i, j int) bool {
			return allEvents[i].GetStart().Before(allEvents[j].GetStart())
		})

		returnString := ""
		for i, event := range allEvents {
			if i >= 5 {
				break
			}

			if len(event.GetDTZID()) == 0 {
				returnString += fmt.Sprintf(" %s; am %d.%d von %d:%02d bis %d:%02d Uhr; Ort: %s;", event.GetSummary(), event.GetStart().Local().Day(), event.GetStart().Local().Month(), event.GetStart().Local().Hour(), event.GetStart().Local().Minute(), event.GetEnd().Local().Hour(), event.GetEnd().Local().Minute(), event.GetLocation())
			}else {
				returnString += fmt.Sprintf(" %s; am %d.%d von %d:%02d bis %d:%02d Uhr; Ort: %s;", event.GetSummary(), event.GetStart().Day(), event.GetStart().Month(), event.GetStart().Hour(), event.GetStart().Minute(), event.GetEnd().Hour(), event.GetEnd().Minute(), event.GetLocation())
			}

		}
		return returnString
	}
	return ""
}
