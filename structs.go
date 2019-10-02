package main

type PDResponse struct {
	Incidents []Incident `json:"incidents"`
}

type Incident struct {
	Id           string `json:"id"`
	IncidentTime string `json:"created_at"`
	Title        string `json:"title"`
	IncidentUrl  string `json:"html_url"`
}

type pdInfo struct {
	Title         string
	IncidentTimes []string
	InsidentLinks []string
	count         int
}
