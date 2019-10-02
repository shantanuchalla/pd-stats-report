package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {

	today := time.Now()
	lastWeek := today.AddDate(0, 0, -7)

	sinceDefault := getDateString(lastWeek)
	untilDefault := getDateString(today)

	apiKey := flag.String("apiKey", "empty", "pager duty user api key")
	since := flag.String("since", sinceDefault, "start date of incident summary")
	until := flag.String("until", untilDefault, "end date of incident summary")
	team := flag.String("teamId", "checkout", "team id to pull report for")

	flag.Parse()
	initClient()

	period := *since + " to " + *until

	pdResponse := getPDInfo(*apiKey, *since, *until, *team)

	var data map[string]*pdInfo = make(map[string]*pdInfo)
	for _, incident := range pdResponse.Incidents {
		info, ok := data[incident.Title]
		if ok {
			info.count++
			info.IncidentTimes = append(info.IncidentTimes, incident.IncidentTime)
			info.InsidentLinks = append(info.InsidentLinks, incident.IncidentUrl)
		} else {
			data[incident.Title] = &pdInfo{
				Title:         incident.Title,
				IncidentTimes: []string{incident.IncidentTime},
				InsidentLinks: []string{incident.IncidentUrl},
				count:         1,
			}
		}
	}

	var csvData [][]string
	csvData = append(csvData, []string{"Period", "Count", "Title", "Comments", "Next Steps", "triggered_at", "links"})
	for _, info := range data {
		csvData = append(csvData, []string{period, strconv.Itoa(info.count), info.Title, "", "", strings.Join(info.IncidentTimes, "\n"), strings.Join(info.InsidentLinks, "\n")})
	}
	log.Println(csvData)
	writeToCsv(csvData)
}

func getDateString(time time.Time) string {
	var year string = strconv.Itoa(time.Year())
	var month string = fmt.Sprintf("%02d", int(time.Month()))
	var dayOfMonth string = fmt.Sprintf("%02d", time.Day())

	return year + "-" + month + "-" + dayOfMonth + "T07:30:00"
}
