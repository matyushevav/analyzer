package models

import "time"

type Job struct {
	ID  int
	URL string
}

type Result struct {
	JobID    int
	URL      string
	Status   int
	Duration time.Duration
}
