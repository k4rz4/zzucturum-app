package models

type Counter struct {
	Domain    string
	NumberOfRequests  int64
	Timestamp int64
}

type Stats struct {
	Domain           string
	NumberOfRequests int64
}

