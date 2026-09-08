package worker

import (
	"analyzer/pkg/models"
	"math/rand"
	"sync"
	"time"
)

// MockHTTPRequest имитирует запрос к URL
func MockHTTPRequest(url string) (time.Duration, int) {
	delay := time.Duration(rand.Intn(500)+100) * time.Millisecond
	time.Sleep(delay)
	// Генерируем случайный статус
	statuses := []int{200, 401, 403, 500}
	status := statuses[rand.Intn(len(statuses))]

	return delay, status

}

// Worker обрабатывает задания из канала
func Worker(jobs <-chan models.Job, results chan<- models.Result, wg *sync.WaitGroup) {
	defer wg.Done()

	// Читаем канал jobs
	for job := range jobs {
		duration, status := MockHTTPRequest(job.URL)
		results <- models.Result{JobID: job.ID,
			URL:      job.URL,
			Status:   status,
			Duration: duration,
		}
	}
}
