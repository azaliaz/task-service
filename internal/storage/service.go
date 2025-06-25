package storage

import (
	"sync"
	"time"
)

type TaskStatus string

const (
	StatusPending   TaskStatus = "pending"
	StatusRunning   TaskStatus = "running"
	StatusCompleted TaskStatus = "completed"
	StatusDeleted   TaskStatus = "deleted"
)

type Task struct {
	ID          string
	CreatedAt   time.Time
	StartedAt   time.Time
	CompletedAt time.Time
	Status      TaskStatus
	Result      string
}

type Service interface {
	CreateTask() (*Task, error)
	DeleteTask(id string) error
	GetTaskStatus(id string) (*Task, error)
	UpdateTaskStatus(id string, status TaskStatus) error
	UpdateTaskStartedAt(id string, startedAt time.Time) error
	UpdateTaskCompletedAt(id string, completedAt time.Time) error
	UpdateTaskResult(id string, result string) error
}

func NewService() *service {
	return &service{
		tasks: make(map[string]*Task),
	}
}

type service struct {
	mu    sync.RWMutex
	tasks map[string]*Task
}

func (s *service) runLongTask(task *Task) {
	s.mu.Lock()
	task.Status = StatusRunning
	task.StartedAt = time.Now()
	s.mu.Unlock()

	time.Sleep(4 * time.Minute)

	s.mu.Lock()
	task.Status = StatusCompleted
	task.CompletedAt = time.Now()
	task.Result = "Task completed successfully"
	s.mu.Unlock()
}
