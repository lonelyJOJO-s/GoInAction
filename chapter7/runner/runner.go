package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

type runner struct {
	// 来自系统的信号
	interrupt chan os.Signal
	// 工作是否完成
	complete chan error
	// 是否超时
	timeout <-chan time.Time
	// 任务队列
	tasks []func(int)
}

var ErrTimeout = errors.New("recieved timeout")
var ErrInterrupt = errors.New("recieved interrupt")

// 工厂函数
// 创建runner类
func New(d time.Duration) *runner {
	return &runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// 添加任务
func (r *runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// 执行注册的任务
func (r *runner) Run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

// gotInterruption的实现
func (r *runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}

func (r *runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt) //Notify函数让signal包将输入信号转发到c。如果没有列出要传递的信号，会将所有输入信号传递到c；否则只传递列出的输入信号。

	// 执行任务函数
	go func() {
		r.complete <- r.Run()
	}()

	//
	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}
