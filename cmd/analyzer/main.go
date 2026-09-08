package main

import (
	"analyzer/internal/worker"
	"analyzer/pkg/models"
	"fmt"
	"sync"
	"time"
)

func main() {
	// Исходные данные
	urls := []string{
		"https://example1.com",
		"https://example2.com",
		"https://example3.com",
		"https://example4.com",
		"https://example5.com",
		"https://example6.com",
		"https://example7.com",
		"https://example8.com",
		"https://example9.com",
		"https://example10.com",
	}

	// Подготовка заданий
	var jobsSlice []models.Job
	for i, url := range urls {
		jobsSlice = append(jobsSlice, models.Job{ID: i, URL: url})
	}

	// Определяем размер буфера каналов
	bufferSize := len(jobsSlice)

	// Определяем размер пула воркеров
	workersNum := 5

	// Создаем каналы заданий и результатов
	jobs := make(chan models.Job, bufferSize)
	results := make(chan models.Result, bufferSize)

	var wg sync.WaitGroup

	// Запускаем пул воркеров
	for i := 1; i <= workersNum; i++ {
		wg.Add(1)
		go worker.Worker(jobs, results, &wg)
	}

	// Отправляем задания в канал задач
	for _, j := range jobsSlice {
		jobs <- j
	}
	close(jobs)

	// Дожидаемся окончания работы всех воркеров
	go func() {
		wg.Wait()
		close(results)
	}()

	// Читаем результаты из канала results
	var totalTime time.Duration
	var countSuccess int
	var resultSlice []models.Result
	for res := range results {
		resultSlice = append(resultSlice, res)
		totalTime += res.Duration
		if res.Status >= 200 && res.Status < 300 {
			countSuccess++
		}
	}

	// Выводим статистику
	fmt.Println("\n=====СТАТИСТИКА=====")
	fmt.Printf("Обработано адресов: %d \n", len(urls))
	fmt.Printf("Количество успешных запросов: %d\n", countSuccess)
	if len(resultSlice) > 0 {
		avgTime := totalTime / time.Duration(len(resultSlice))
		fmt.Printf("Среднее время обработки: %.3f с\n", avgTime.Seconds())
	} else {
		fmt.Printf("Нет данных для расчета среднего времени выполнения запроса")
	}
}
