package models

import "time"

type DocumentRequest struct {
	EventName  string `json:"event_name"`
	EventDescription string `json:"event_description"`
}

type DocumentResponse struct {
	EventName     string    `json:"event_name"`
	CreatedAt time.Time `json:"created_at"`
	EventDescription   string    `json:"event_description"`
}

type SearchResponse struct {
	Time string  `json:"time"`
	Hits string	`json:"hits"`
	ResultDocuments []DocumentResponse `json:"result_documents"`
}
