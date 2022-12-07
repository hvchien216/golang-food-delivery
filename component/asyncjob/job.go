package asyncjob

import (
	"context"
	"time"
)

type Job interface {
	Execute(ctx context.Context) error
	Retry(ctx context.Context) error
	State() JobState
	SetRetryDurations(times []time.Duration)
}

const (
	defaultMaxTimeout    = time.Second * 10
	defaultMaxRetryCount = 3
)

var (
	defaultRetryTime = []time.Duration{time.Second, time.Second * 5, time.Second * 10}
)

type JobState int

type JobHandler func(ctx context.Context) error

const (
	StateInit JobState = iota
	StateRunning
	StateFailed
	StateTimeout
	StateCompleted
	StateRetryFailed
)

type jobConfig struct {
	MaxTimeout time.Duration
	Retries    []time.Duration
}

func (js JobState) String() string {
	return []string{"Init", "Running", "Failed", "Timeout", "Completed", "RetryFailed"}[js]
}

type job struct {
	config     jobConfig
	handler    JobHandler
	state      JobState
	retryIndex int
	stopChan   chan bool
}

func NewJob(handler JobHandler) *job {
	return &job{
		config: jobConfig{
			MaxTimeout: defaultMaxTimeout,
			Retries:    defaultRetryTime,
		},
		handler:    handler,
		state:      StateInit,
		retryIndex: -1,
		stopChan:   make(chan bool),
	}
}

func (j *job) Execute(ctx context.Context) error {
	j.state = StateRunning

	var err error
	err = j.handler(ctx)

	if err != nil {
		j.state = StateFailed
		return err
	}

	j.state = StateCompleted
	return nil

	//ch := make(chan error)
	//ctxJob, doneFunc := context.WithCancel(ctx)
	//
	//go func() {
	//	j.state = StateRunning
	//	var err error
	//
	//	err = j.handler(ctxJob)
	//
	//	if err != nil {
	//		j.state = StateFailed
	//		ch <- err
	//		return
	//	}
	//
	//	j.state = StateCompleted
	//	ch <- err
	//}()
	//
	//for {
	//	select {
	//	case <-j.stopChan:
	//		break
	//	default:
	//		fmt.Println("Hello world")
	//	}
	//}
	//
	//select {
	//case err := <-ch:
	//	doneFunc()
	//	return err
	//case <-j.stopChan:
	//	doneFunc()
	//	return nil
	//}
	//return <-ch
}

func (j *job) Retry(ctx context.Context) error {
	//if j.retryIndex == len(j.config.Retries)-1 {
	//	return nil
	//}

	j.retryIndex += 1
	time.Sleep(j.config.Retries[j.retryIndex])

	// j.state=StateRunning
	err := j.Execute(ctx)

	if err == nil {
		j.state = StateCompleted
		return nil
	}

	if j.retryIndex == len(j.config.Retries)-1 {
		j.state = StateRetryFailed
		return err
	}

	j.state = StateFailed
	return err
}

func (j *job) State() JobState { return j.state }
func (j *job) RetryIndex() int { return j.retryIndex }

func (j *job) SetRetryDurations(times []time.Duration) {
	if len(times) == 0 {
		return
	}

	j.config.Retries = times
}
