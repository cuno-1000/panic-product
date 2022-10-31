package handler

import (
	"fmt"
)

type Job interface {
	Do()
}

type ApplyRequest struct {
	UserId    uint64
	EventUuid string
	Link      string
	RequestId string
}

func (s *ApplyRequest) Do() {
	MqClientProcess(s.UserId, s.EventUuid, s.Link, s.RequestId)
}

type JobQueue chan Job

type Worker struct {
	JobChan JobQueue
}

func NewWorker() Worker {
	return Worker{JobChan: make(chan Job)}
}

func (w Worker) Run(wq chan JobQueue) {
	go func() {
		for {
			wq <- w.JobChan
			select {
			case job := <-w.JobChan:
				job.Do()
			}
		}
	}()
}

type WorkerPool struct {
	Workerlen   int
	JobQueue    JobQueue
	WorkerQueue chan JobQueue
}

var ApplyPool *WorkerPool

func NewWorkerPool(workerlen int) *WorkerPool {
	return &WorkerPool{
		Workerlen:   workerlen,
		JobQueue:    make(JobQueue),
		WorkerQueue: make(chan JobQueue, workerlen),
	}
}
func (wp *WorkerPool) Run() {
	fmt.Println("初始化worker")
	for i := 0; i < wp.Workerlen; i++ {
		worker := NewWorker()
		worker.Run(wp.WorkerQueue)
	}
	go func() {
		for {
			select {
			case job := <-wp.JobQueue:
				worker := <-wp.WorkerQueue
				worker <- job
			}
		}
	}()
}
