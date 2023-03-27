package models

import "time"

type DocumentRequest struct {
	EventName  string `json:"event_name"`
	EventDescription string `json:"event_description"`
	EventAddress string `json:"event_address"`
}

type DocumentResponse struct {
	EventName     string    `json:"event_name"`
	CreatedAt time.Time `json:"created_at"`
	EventDescription   string    `json:"event_description"`
	EventAddress 		string `json:"event_address"`
}

type SearchResponse struct {
	Time string  `json:"time"`
	Hits string	`json:"hits"`
	ResultDocuments []DocumentResponse `json:"result_documents"`
}

type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}
type ESEvent struct {
	Event
	Location Location `json:"location"`

}
