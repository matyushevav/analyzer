package worker

import (
	"analyzer/pkg/models"
	"math/rand"
	"sync"
	"time"
)

var statuses = []int{200, 401, 403, 500}

// MockHTTPRequest имитирует запрос к URL
func MockHTTPRequest(url string) int {
	delay := time.Duration(rand.Intn(500)+100) * time.Millisecond
	time.Sleep(delay)

	return statuses[rand.Intn(len(statuses))]
}

// Worker обрабатывает задания из канала
func Worker(jobs <-chan models.Job, results chan<- models.Result, wg *sync.WaitGroup) {
	defer wg.Done()

	// Читаем канал jobs
	for job := range jobs {
		start := time.Now()
		status := MockHTTPRequest(job.URL)
		duration := time.Since(start)

		results <- models.Result{
			JobID:    job.ID,
			URL:      job.URL,
			Status:   status,
			Duration: duration,
		}
	}
}
