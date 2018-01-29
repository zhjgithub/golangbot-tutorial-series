package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func GoroutineExample() {
	bufferedChannel()
	workPool()
}

func g(ch chan string) {
	fmt.Println("send naveen")
	ch <- "naveen"
	fmt.Println("send paul")
	ch <- "paul"
	fmt.Println("send jim")
	ch <- "jim"
	fmt.Println("send end")
}

func bufferedChannel() {
	ch := make(chan string, 1) // change different capacity value to see different print out, e.g. 0, 1, 2, 3
	go g(ch)
	fmt.Println("receive naveen")
	fmt.Println(<-ch)
	fmt.Println("receive paul")
	fmt.Println(<-ch)
	fmt.Println("receive jim")
	fmt.Println(<-ch)
	fmt.Println("receive end")
}

type Job struct {
	id       int
	randomno int
}
type Result struct {
	job         Job
	sumofdigits int
	worker      int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

func worker(wg *sync.WaitGroup, w int) {
	for job := range jobs {
		output := Result{job, digits(job.randomno), w}
		results <- output
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg, i)
	}
	wg.Wait()
	close(results)
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d by worker %d\n", result.job.id, result.job.randomno, result.sumofdigits, result.worker)
	}
	done <- true
}

func workPool() {
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 25
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
