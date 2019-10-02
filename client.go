package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var client http.Client

func initClient() {
	client = http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
}

func getPDInfo(authKey string, since string, until string, team string) PDResponse {
	request, err := http.NewRequest("GET", "https://api.pagerduty.com/incidents", nil)
	if err != nil {
		log.Fatalln(err)
	}

	auth := "Token token="
	auth += authKey

	log.Println(auth)

	request.Header.Set("Authorization", auth)
	request.Header.Set("Accept", "application/vnd.pagerduty+json;version=2")

	query := request.URL.Query()
	query.Add("since", since)
	query.Add("until", until)
	query.Add("team_ids[]", team)
	query.Add("urgencies[]", "high")
	query.Add("time_zone", "Asia/Calcutta")

	request.URL.RawQuery = query.Encode()

	log.Println(request.Header.Get("Authorization"))
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp.Status)
	
	var response PDResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalln(err)
	}

	return response
}
